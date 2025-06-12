package Masterhandlers

import (
	"context"
	"strconv"

	"farmerapp/go_backend/db"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

type Language struct {
	ID           int64  `json:"id,omitempty"`
	LanguageName string `json:"language_name" validate:"required"`
}

func InsertLanguage(c *fiber.Ctx) error {
	var req Language
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request format"})
	}

	if err := validate.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Language name is required"})
	}

	_, err := db.Pool.Exec(context.Background(), "SELECT admin_schema.insert_language($1)", req.LanguageName)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to insert language", "details": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Language inserted successfully"})
}

func GetAllLanguages(c *fiber.Ctx) error {
	rows, err := db.Pool.Query(context.Background(), "SELECT * FROM admin_schema.get_all_languages()")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch languages", "details": err.Error()})
	}
	defer rows.Close()

	var languages []Language
	for rows.Next() {
		var lang Language
		if err := rows.Scan(&lang.ID, &lang.LanguageName); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error parsing language data", "details": err.Error()})
		}
		languages = append(languages, lang)
	}

	return c.JSON(languages)
}

func GetLanguageByID(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID format"})
	}

	var lang Language
	err = db.Pool.QueryRow(context.Background(), "SELECT * FROM admin_schema.get_language_by_id($1)", id).
		Scan(&lang.ID, &lang.LanguageName)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Language not found"})
	}

	return c.JSON(lang)
}

func UpdateLanguage(c *fiber.Ctx) error {
	var req Language
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request format"})
	}

	if err := validate.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Language name is required"})
	}

	_, err := db.Pool.Exec(context.Background(), "SELECT admin_schema.update_language($1, $2)", req.ID, req.LanguageName)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update language", "details": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Language updated successfully"})
}
