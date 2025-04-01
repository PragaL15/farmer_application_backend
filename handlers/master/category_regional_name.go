package Masterhandlers

import (
	"context"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/PragaL15/go_newBackend/go_backend/db"
)

type CategoryRegionallanguage struct {
	ID           int64  `json:"category_regional_id"`
	LanguageID   int64  `json:"language_id"`
	CategoryID   int64  `json:"category_id"`
	RegionalName string `json:"category_regional_name"`
}

// Get all category regional names
func GetProductCategoryRegional(c *fiber.Ctx) error {
	rows, err := db.Pool.Query(context.Background(), "SELECT id, language_id, category_id, regional_name FROM admin_schema.get_all_category_regional_names()")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch category regional names", "details": err.Error()})
	}
	defer rows.Close()

	var categories []CategoryRegionallanguage
	for rows.Next() {
		var category CategoryRegionallanguage
		if err := rows.Scan(&category.ID, &category.LanguageID, &category.CategoryID, &category.RegionalName); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error scanning category regional data", "details": err.Error()})
		}
		categories = append(categories, category)
	}
	return c.JSON(categories)
}

// Get category regional name by ID
func GetProductCategoryRegionalByID(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID format"})
	}

	var category CategoryRegionallanguage
	err = db.Pool.QueryRow(context.Background(), "SELECT id, language_id, category_id, regional_name FROM admin_schema.get_category_regional_name_by_id($1)", id).
		Scan(&category.ID, &category.LanguageID, &category.CategoryID, &category.RegionalName)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Category regional name not found", "details": err.Error()})
	}
	return c.JSON(category)
}

// Insert a new category regional name
func InsertProductCategoryRegional(c *fiber.Ctx) error {
	var req CategoryRegionallanguage
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request format"})
	}

	_, err := db.Pool.Exec(context.Background(), "INSERT INTO admin_schema.category_regional_names (language_id, category_id, regional_name) VALUES ($1, $2, $3)", req.LanguageID, req.CategoryID, req.RegionalName)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to insert category regional name", "details": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "Category regional name added successfully"})
}

// Update an existing category regional name
func UpdateProductCategoryRegional(c *fiber.Ctx) error {
	var req CategoryRegionallanguage
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request format"})
	}

	_, err := db.Pool.Exec(context.Background(), "UPDATE admin_schema.category_regional_names SET language_id=$1, category_id=$2, regional_name=$3 WHERE id=$4", req.LanguageID, req.CategoryID, req.RegionalName, req.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update category regional name", "details": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "Category regional name updated successfully"})
}
