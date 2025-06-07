package Masterhandlers

import (
	"context"
	"strconv"

	"farmerapp/go_backend/db"

	"github.com/gofiber/fiber/v2"
)

// ProductRegional defines the structure for regional product names
type ProductRegional struct {
	ID                  int64  `json:"product_regional_id"`
	LanguageID          int64  `json:"language_id"`
	ProductID           int64  `json:"product_id"`
	ProductName         string `json:"product_name"`
	ProductRegionalName string `json:"product_regional_name"`
}

// @Summary      Get all product regional names
// @Description  Retrieves all product regional names from the database
// @Tags         Product Regional Names
// @Produce      json
// @Success      200 {array} ProductRegional
// @Router       /product-regional [get]
func GetAllProductRegionalNames(c *fiber.Ctx) error {
	rows, err := db.Pool.Query(context.Background(), "SELECT * FROM admin_schema.get_all_product_regional_names()")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch product regional names", "details": err.Error()})
	}
	defer rows.Close()

	var products []ProductRegional
	for rows.Next() {
		var product ProductRegional
		if err := rows.Scan(&product.ID, &product.LanguageID, &product.ProductID, &product.ProductName, &product.ProductRegionalName); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error scanning product regional data", "details": err.Error()})
		}
		products = append(products, product)
	}
	return c.JSON(products)
}

// @Summary      Get product regional name by ID
// @Description  Retrieves a product regional name by its ID
// @Tags         Product Regional Names
// @Produce      json
// @Param        id   path      int  true  "Product Regional ID"
// @Success      200  {object}  ProductRegional
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Router       /product-regional/{id} [get]
func GetProductRegionalNameByID(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID format"})
	}

	var product ProductRegional
	err = db.Pool.QueryRow(context.Background(), "SELECT * FROM admin_schema.get_product_regional_name_by_id($1)", id).
		Scan(&product.ID, &product.LanguageID, &product.ProductID, &product.ProductName, &product.ProductRegionalName)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Product regional name not found"})
	}
	return c.JSON(product)
}

// @Summary      Insert a new product regional name
// @Description  Adds a new product regional name to the database
// @Tags         Product Regional Names
// @Accept       json
// @Produce      json
// @Param        request  body  ProductRegional  true  "Product Regional Name Input"
// @Success      200  {object}  map[string]string
// @Failure      400  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /product-regional [post]
func InsertProductRegionalName(c *fiber.Ctx) error {
	type Request struct {
		LanguageID          int64  `json:"language_id"`
		ProductID           int64  `json:"product_id"`
		ProductRegionalName string `json:"product_regional_name"`
	}

	var req Request
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request format"})
	}

	_, err := db.Pool.Exec(context.Background(), "SELECT admin_schema.insert_product_regional_name($1, $2, $3)",
		req.LanguageID, req.ProductID, req.ProductRegionalName)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to insert product regional name", "details": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "Product regional name added successfully"})
}

// @Summary      Update product regional name
// @Description  Updates a product regional name by ID
// @Tags         Product Regional Names
// @Accept       json
// @Produce      json
// @Param        request  body  ProductRegional  true  "Product Regional Name Input"
// @Success      200  {object}  map[string]string
// @Failure      400  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /product-regional [put]
func UpdateProductRegionalName(c *fiber.Ctx) error {
	type Request struct {
		ID                  int64  `json:"product_regional_id"`
		LanguageID          int64  `json:"language_id"`
		ProductID           int64  `json:"product_id"`
		ProductRegionalName string `json:"product_regional_name"`
	}

	var req Request
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request format"})
	}

	_, err := db.Pool.Exec(context.Background(), "SELECT admin_schema.update_product_regional_name($1, $2, $3, $4)",
		req.ID, req.LanguageID, req.ProductID, req.ProductRegionalName)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update product regional name", "details": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "Product regional name updated successfully"})
}
