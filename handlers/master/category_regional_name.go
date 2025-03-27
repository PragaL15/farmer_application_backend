package Masterhandlers

import (
	"context"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/PragaL15/go_newBackend/go_backend/db"
)

type CategoryRegionallanguage struct {
	ID                int64  `json:"category_regional_id"`
	LanguageID        int64  `json:"language_id"`
	CategoryID        int64  `json:"category_id"`
	RegionalName      string `json:"category_regional_name"`
}

func GetProductCategoryRegional(c *fiber.Ctx) error {
	rows, err := db.Pool.Query(context.Background(), "SELECT * FROM admin_schema.get_all_category_regional_names()")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch category regional names", "details": err.Error()})
	}
	defer rows.Close()

	var categories []CategoryRegional
	for rows.Next() {
		var category CategoryRegional
		if err := rows.Scan(&category.ID, &category.LanguageID, &category.CategoryID, &category.RegionalName); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error scanning category regional data", "details": err.Error()})
		}
		categories = append(categories, category)
	}
	return c.JSON(categories)
}

func GetProductCategoryRegionalByID(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID format"})
	}

	var category CategoryRegional
	err = db.Pool.QueryRow(context.Background(), "SELECT * FROM admin_schema.get_category_regional_name_by_id($1)", id).
		Scan(&category.ID, &category.LanguageID, &category.CategoryID, &category.RegionalName)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Category regional name not found"})
	}
	return c.JSON(category)
}

// Create a new category regional name
func InsertProductCategoryRegional(c *fiber.Ctx) error {
	type Request struct {
		LanguageID   int64  `json:"language_id"`
		CategoryID   int64  `json:"category_id"`
		RegionalName string `json:"category_regional_name"`
	}

	var req Request
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request format"})
	}

	_, err := db.Pool.Exec(context.Background(), "SELECT admin_schema.insert_category_regional_name($1, $2, $3)", req.LanguageID, req.CategoryID, req.RegionalName)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to insert category regional name", "details": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "Category regional name added successfully"})
}

func UpdateProductCategoryRegional(c *fiber.Ctx) error {
	type Request struct {
		ID          int64  `json:"category_regional_id"`
		LanguageID  int64  `json:"language_id"`
		CategoryID  int64  `json:"category_id"`
		RegionalName string `json:"category_regional_name"`
	}

	var req Request
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request format"})
	}

	_, err := db.Pool.Exec(context.Background(), "SELECT admin_schema.update_category_regional_name($1, $2, $3, $4)", req.ID, req.LanguageID, req.CategoryID, req.RegionalName)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update category regional name", "details": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "Category regional name updated successfully"})
}
