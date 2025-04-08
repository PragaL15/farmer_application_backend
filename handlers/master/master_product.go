package Masterhandlers

import (
	"context"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/PragaL15/go_newBackend/go_backend/db"
)

type Productget struct {
	ProductID           int64  `json:"product_id"`
	CategoryID          int64  `json:"category_id"`
	CategoryName        string `json:"category_name"`
	ProductName         string `json:"product_name"`
	ImagePath           string `json:"image_path"`
	ActiveStatus        int    `json:"active_status"`
	ProductRegionalID   int64  `json:"product_regional_id"`
	ProductRegionalName string `json:"product_regional_name"`
}

type ProductInsertRequest struct {
	CategoryID        int64  `json:"category_id"`
	ProductName       string `json:"product_name"`
	ImagePath         string `json:"image_path"`
	ActiveStatus      int    `json:"active_status"`
	ProductRegionalID int64  `json:"product_regional_id"`
}

type ProductUpdateRequest struct {
	ProductID         int    `json:"productID"`
	CategoryID        int    `json:"categoryID"`
	ProductName       string `json:"productName"`
	ImagePath         string `json:"imagePath"`
	ActiveStatus      int    `json:"activeStatus"`
	ProductRegionalID int    `json:"productRegionalID"`
}


func GetAllProducts(c *fiber.Ctx) error {
	rows, err := db.Pool.Query(context.Background(), "SELECT * FROM admin_schema.get_all_products()")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch products", "details": err.Error()})
	}
	defer rows.Close()

	var products []Productget
	for rows.Next() {
		var p Productget
		err := rows.Scan(&p.ProductID, &p.CategoryID, &p.CategoryName, &p.ProductName, &p.ImagePath, &p.ActiveStatus, &p.ProductRegionalID, &p.ProductRegionalName)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error scanning product data", "details": err.Error()})
		}
		products = append(products, p)
	}
	return c.JSON(products)
}

func GetProductByID(c *fiber.Ctx) error {
	productID, err := strconv.ParseInt(c.Params("product_id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid product ID format"})
	}

	var p Productget
	err = db.Pool.QueryRow(context.Background(), "SELECT * FROM admin_schema.get_product_by_id($1)", productID).
		Scan(&p.ProductID, &p.CategoryID, &p.CategoryName, &p.ProductName, &p.ImagePath, &p.ActiveStatus, &p.ProductRegionalID, &p.ProductRegionalName)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Product not found", "details": err.Error()})
	}
	return c.JSON(p)
}

func GetProductsByCategoryID(c *fiber.Ctx) error {
	categoryID, err := strconv.ParseInt(c.Params("category_id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid category ID format"})
	}

	rows, err := db.Pool.Query(context.Background(), "SELECT * FROM admin_schema.get_products_by_category($1)", categoryID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch products", "details": err.Error()})
	}
	defer rows.Close()

	var products []Productget
	for rows.Next() {
		var p Productget
		err := rows.Scan(&p.ProductID, &p.CategoryID, &p.CategoryName, &p.ProductName, &p.ImagePath, &p.ActiveStatus, &p.ProductRegionalID, &p.ProductRegionalName)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error scanning product data", "details": err.Error()})
		}
		products = append(products, p)
	}
	return c.JSON(products)
}

func InsertProduct(c *fiber.Ctx) error {
	var req ProductInsertRequest

	// Parse JSON body into the struct
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request format",
		})
	}

	// Default active status to 1 if 0 or missing
	if req.ActiveStatus == 0 {
		req.ActiveStatus = 1
	}

	// Execute stored procedure with explicit type casts
	query := `
		SELECT admin_schema.insert_product(
			$1::integer, $2::text, $3::text, $4::integer, $5::integer
		)
	`

	_, err := db.Pool.Exec(context.Background(), query,
		req.CategoryID,
		req.ProductName,
		req.ImagePath,
		req.ActiveStatus,
		req.ProductRegionalID,
	)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to insert product",
			"details": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Product added successfully",
	})
}


func UpdateProductHandler(c *fiber.Ctx) error {
	var req ProductUpdateRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request format",
			"details": err.Error(),
		})
	}

	// Debug: print parsed request
	// fmt.Printf("Parsed update request: %+v\n", req)

	// Ensure ProductID is valid
	if req.ProductID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "ProductID is required and must be greater than 0",
		})
	}

	var existing ProductInsertRequest
	query := `
		SELECT category_id, product_name, image_path, active_status, COALESCE(product_regional_id, 0)
		FROM admin_schema.master_product WHERE product_id = $1
	`
	err := db.Pool.QueryRow(context.Background(), query, req.ProductID).Scan(
		&existing.CategoryID,
		&existing.ProductName,
		&existing.ImagePath,
		&existing.ActiveStatus,
		&existing.ProductRegionalID,
	)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   "Product not found",
			"details": err.Error(),
		})
	}

	if req.CategoryID == 0 {
		req.CategoryID = int(existing.CategoryID)
	}
	if req.ProductName == "" {
		req.ProductName = existing.ProductName
	}
	if req.ImagePath == "" {
		req.ImagePath = existing.ImagePath
	}
	if req.ActiveStatus == 0 {
		req.ActiveStatus = existing.ActiveStatus
	}
	if req.ProductRegionalID == 0 {
		req.ProductRegionalID = int(existing.ProductRegionalID)
	}

	updateQuery := `SELECT admin_schema.update_product($1, $2, $3, $4, $5, $6)`
	_, err = db.Pool.Exec(context.Background(), updateQuery,
		req.ProductID,
		req.CategoryID,
		req.ProductName,
		req.ImagePath,
		req.ActiveStatus,
		req.ProductRegionalID,
	)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to update product",
			"details": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Product updated successfully",
	})
}
