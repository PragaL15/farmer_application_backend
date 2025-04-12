package Masterhandlers

import (
	"context"
	"strconv"

	"github.com/PragaL15/go_newBackend/go_backend/db"
	"github.com/gofiber/fiber/v2"
	"github.com/guregu/null/v5"
)

type Productget struct {
	ProductID         int64  `json:"product_id"`
	CategoryID        int64  `json:"category_id"`
	CategoryName      string `json:"category_name"`
	ProductName       string `json:"product_name"`
	ImagePath         null.String `json:"image_path"`
	ActiveStatus      int    `json:"active_status"`
	ProductRegionalID null.Int64  `json:"product_regional_id"`
	ProductRegionalName null.String `json:"product_regional_name"`
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
		if err := rows.Scan(&p.ProductID, &p.CategoryID, &p.CategoryName, &p.ProductName, &p.ImagePath, &p.ActiveStatus, &p.ProductRegionalID, &p.ProductRegionalName); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error scanning product data", "details": err.Error()})
		}
		products = append(products, p)
	}
	return c.JSON(products)
}

type ProductInsertRequest struct {
	CategoryID        int    `json:"category_id"`
	ProductName       string `json:"product_name"`
	ImagePath         string `json:"image_path"`
	ActiveStatus      int    `json:"active_status"`
	ProductRegionalID int    `json:"product_regional_id"`
}

func InsertProduct(c *fiber.Ctx) error {
	var req ProductInsertRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request format"})
	}

	if req.ActiveStatus == 0 {
		req.ActiveStatus = 1
	}

	_, err := db.Pool.Exec(
		context.Background(),
		"SELECT admin_schema.insert_product($1, $2, $3, $4, $5)",
		req.CategoryID, req.ProductName, req.ImagePath, req.ActiveStatus, req.ProductRegionalID,
	)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to insert product",
			"details": err.Error(),
		})
	}

	return c.JSON(fiber.Map{"message": "Product added successfully"})
}

func UpdateProductHandler(c *fiber.Ctx) error {
	type Request struct {
		ProductID         int64  `json:"product_id"`
		CategoryID        int64  `json:"category_id"`
		ProductName       string `json:"product_name"`
		ImagePath         string `json:"image_path"`
		ActiveStatus      int    `json:"active_status"`
		ProductRegionalID int64  `json:"product_regional_id"`
	}

	var req Request
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request format",
		})
	}

	_, err := db.Pool.Exec(context.Background(),
		"SELECT admin_schema.update_product($1, $2, $3, $4, $5, $6)",
		req.ProductID, req.CategoryID, req.ProductName, req.ImagePath, req.ActiveStatus, req.ProductRegionalID,
	)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to update product",
			"details": err.Error(),
		})
	}

	return c.JSON(fiber.Map{"message": "Product updated successfully"})
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
		var product Productget
		if err := rows.Scan(&product.ProductID, &product.CategoryID, &product.CategoryName, &product.ProductName, &product.ImagePath, &product.ActiveStatus, &product.ProductRegionalID, &product.ProductRegionalName); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error scanning product data", "details": err.Error()})
		}
		products = append(products, product)
	}
	return c.JSON(products)
}

func GetProductByID(c *fiber.Ctx) error {
	productID, err := strconv.ParseInt(c.Params("product_id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid product ID format"})
	}

	var product Productget
	query := `SELECT * FROM admin_schema.get_product_by_id($1)`
	err = db.Pool.QueryRow(context.Background(), query, productID).
		Scan(
			&product.ProductID,
			&product.CategoryID,
			&product.CategoryName,
			&product.ProductName,
			&product.ImagePath,
			&product.ActiveStatus,
			&product.ProductRegionalID,
			&product.ProductRegionalName,
		)

	if err != nil {
		if err.Error() == "no rows in result set" {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Product not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Database error", "details": err.Error()})
	}

	return c.JSON(product)
}
