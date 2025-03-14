package Masterhandlers

import (
	"context"
	"log"
	"strconv"
"database/sql"
	"github.com/gofiber/fiber/v2"
	"github.com/PragaL15/go_newBackend/go_backend/db"
	"github.com/go-playground/validator/v10"
)

func InsertMasterProduct(c *fiber.Ctx) error {
	type Request struct {
		CategoryID    int    `json:"category_id" validate:"required,min=1"`
		ProductName   string `json:"product_name" validate:"required,max=100"`
		Status        int    `json:"status" validate:"required"`
		ImagePath     string `json:"image_path"`
		RegionalName1 string `json:"regional_name1"`
		RegionalName2 string `json:"regional_name2"`
		RegionalName3 string `json:"regional_name3"`
		RegionalName4 string `json:"regional_name4"`
	}

	var req Request
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	_, err := db.Pool.Exec(context.Background(),
		`CALL insert_master_product($1, $2, $3, $4, $5, $6, $7, $8)`,
		req.CategoryID, req.ProductName, req.Status, req.ImagePath,
		req.RegionalName1, req.RegionalName2, req.RegionalName3, req.RegionalName4)

	if err != nil {
		log.Printf("Failed to insert product: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to insert product"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Product added successfully"})
}

func UpdateMasterProduct(c *fiber.Ctx) error {
	type Request struct {
		ProductID     int    `json:"product_id" validate:"required,min=1"`
		CategoryID    int    `json:"category_id" validate:"required,min=1"`
		ProductName   string `json:"product_name" validate:"required,max=100"`
		Status        int    `json:"status" validate:"required"`
		ImagePath     string `json:"image_path"`
		RegionalName1 string `json:"regional_name1"`
		RegionalName2 string `json:"regional_name2"`
		RegionalName3 string `json:"regional_name3"`
		RegionalName4 string `json:"regional_name4"`
	}

	var req Request
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	_, err := db.Pool.Exec(context.Background(),
		`CALL update_master_product($1, $2, $3, $4, $5, $6, $7, $8, $9)`,
		req.ProductID, req.CategoryID, req.ProductName, req.Status, req.ImagePath,
		req.RegionalName1, req.RegionalName2, req.RegionalName3, req.RegionalName4)

	if err != nil {
		log.Printf("Failed to update product: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update product"})
	}

	return c.JSON(fiber.Map{"message": "Product updated successfully"})
}

func DeleteMasterProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ID is required"})
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID format"})
	}

	_, err = db.Pool.Exec(context.Background(), `CALL delete_master_product($1)`, idInt)
	if err != nil {
		log.Printf("Failed to delete product: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete product"})
	}

	return c.JSON(fiber.Map{"message": "Product deleted successfully"})
}

func GetProducts(c *fiber.Ctx) error {
	rows, err := db.Pool.Query(context.Background(), "SELECT * FROM get_master_products()")
	if err != nil {
		log.Printf("Failed to fetch product records: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch products"})
	}
	defer rows.Close()

	var products []map[string]interface{}

	for rows.Next() {
		var productID, categoryID, status *int
		var productName, categoryName, imagePath, regionalName1, regionalName2, regionalName3, regionalName4 *string

		if err := rows.Scan(&productID, &categoryID, &categoryName, &productName, &status, &imagePath, &regionalName1, &regionalName2, &regionalName3, &regionalName4); err != nil {
			log.Printf("Error scanning row: %v", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error processing data"})
		}

		products = append(products, map[string]interface{}{
			"product_id":     productID,
			"category_id":    categoryID,
			"category_name":  categoryName,
			"product_name":   productName,
			"status":         status,
			"image_path":     imagePath,
			"regional_name1": regionalName1,
			"regional_name2": regionalName2,
			"regional_name3": regionalName3,
			"regional_name4": regionalName4,
		})
	}

	return c.JSON(products)
}

func GetProductByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ID is required"})
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID format"})
	}

	var productID, categoryID, status *int
	var productName, categoryName, imagePath, regionalName1, regionalName2, regionalName3, regionalName4 *string

	err = db.Pool.QueryRow(context.Background(),
		`SELECT * FROM get_master_product_by_id($1)`, idInt).
		Scan(&productID, &categoryID, &categoryName, &productName, &status, &imagePath, &regionalName1, &regionalName2, &regionalName3, &regionalName4)

	if err != nil {
		log.Printf("Failed to fetch product: %v", err)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Product not found"})
	}

	product := map[string]interface{}{
		"product_id":     productID,
		"category_id":    categoryID,
		"category_name":  categoryName,
		"product_name":   productName,
		"status":         status,
		"image_path":     imagePath,
		"regional_name1": regionalName1,
		"regional_name2": regionalName2,
		"regional_name3": regionalName3,
		"regional_name4": regionalName4,
	}

	return c.JSON(product)
}
func GetProductsByCategoryID(c *fiber.Ctx) error {
	categoryID := c.Params("category_id")
	if categoryID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Category ID is required"})
	}

	categoryIDInt, err := strconv.Atoi(categoryID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Category ID format"})
	}

	// Query to fetch products based on category_id
	rows, err := db.Pool.Query(context.Background(), `
		SELECT 
			mp.product_id, 
			mp.category_id, 
			mc.category_name, 
			mp.product_name, 
			mp.status, 
			mp.image_path, 
			mp.regional_name1, 
			mp.regional_name2, 
			mp.regional_name3, 
			mp.regional_name4
		FROM master_product mp
		JOIN master_category_table mc ON mp.category_id = mc.category_id
		WHERE mp.category_id = $1`, categoryIDInt)

	if err != nil {
		log.Printf("Failed to fetch products: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch products"})
	}
	defer rows.Close()

	// Store the fetched products
	var products []map[string]interface{}

	for rows.Next() {
		var productID, categoryID sql.NullInt64
		var productName, categoryName, imagePath, regionalName1, regionalName2, regionalName3, regionalName4 sql.NullString
		var status sql.NullInt64

		err := rows.Scan(&productID, &categoryID, &categoryName, &productName, &status, &imagePath, &regionalName1, &regionalName2, &regionalName3, &regionalName4)
		if err != nil {
			log.Printf("Error scanning row: %v", err)
			continue
		}

		product := map[string]interface{}{
			"product_id":     nullIntToInterface(productID),
			"category_id":    nullIntToInterface(categoryID),
			"category_name":  nullStringToInterface(categoryName),
			"product_name":   nullStringToInterface(productName),
			"status":         nullIntToInterface(status),
			"image_path":     nullStringToInterface(imagePath),
			"regional_name1": nullStringToInterface(regionalName1),
			"regional_name2": nullStringToInterface(regionalName2),
			"regional_name3": nullStringToInterface(regionalName3),
			"regional_name4": nullStringToInterface(regionalName4),
		}

		products = append(products, product)
	}

	// If no products found, return a message
	if len(products) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "No products found for this category"})
	}

	// Return the product list
	return c.JSON(products)
}

// Utility functions for handling NULL values
func nullIntToInterface(value sql.NullInt64) interface{} {
	if value.Valid {
		return value.Int64
	}
	return nil
}

func nullStringToInterface(value sql.NullString) interface{} {
	if value.Valid {
		return value.String
	}
	return nil
}