package Masterhandlers

import (
	"context"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/PragaL15/go_newBackend/go_backend/db"
)

type CategoryRegionalLanguage struct {
	ID           int64  `json:"id"`
	LanguageName string `json:"language_name"` 
	RegionalName string `json:"regional_name"`
}

func GetProductCategoryRegional(c *fiber.Ctx) error {
	rows, err := db.Pool.Query(context.Background(), "SELECT * FROM admin_schema.get_all_category_regional_names()")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to fetch category regional names",
			"details": err.Error(),
		})
	}
	defer rows.Close()

	var categories []CategoryRegionalLanguage
	for rows.Next() {
		var category CategoryRegionalLanguage
		if err := rows.Scan(&category.ID, &category.RegionalName, &category.LanguageName); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":   "Error scanning category regional data",
				"details": err.Error(),
			})
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

	var category CategoryRegionalLanguage
	err = db.Pool.QueryRow(context.Background(), "SELECT * FROM admin_schema.get_category_regional_name_by_id($1)", id).
		Scan(&category.ID, &category.RegionalName, &category.LanguageName)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   "Category regional name not found",
			"details": err.Error(),
		})
	}
	return c.JSON(category)
}

func InsertProductCategoryRegional(c *fiber.Ctx) error {
	type insertRequest struct {
		LanguageID   int64  `json:"language_id"`
		RegionalName string `json:"regional_name"`
	}

	var req insertRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request format"})
	}

	_, err := db.Pool.Exec(context.Background(),
		"SELECT admin_schema.insert_category_regional_name($1, $2)", req.LanguageID, req.RegionalName)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to insert category regional name",
			"details": err.Error(),
		})
	}
	return c.JSON(fiber.Map{"message": "Category regional name added successfully"})
}

func UpdateProductCategoryRegional(c *fiber.Ctx) error {
	type updateRequest struct {
		ID           int64  `json:"id"`
		LanguageID   int64  `json:"language_id"`
		RegionalName string `json:"regional_name"`
	}

	var req updateRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request format"})
	}

	_, err := db.Pool.Exec(context.Background(),
		"SELECT admin_schema.update_category_regional_name($1, $2, $3)",
		req.ID, req.LanguageID, req.RegionalName)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to update category regional name",
			"details": err.Error(),
		})
	}
	return c.JSON(fiber.Map{"message": "Category regional name updated successfully"})
}
