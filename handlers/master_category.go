package handlers

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/PragaL15/go_newBackend/go_backend/db"
	"github.com/go-playground/validator/v10"
	"log"
	"strconv"
	"database/sql"
	"github.com/guregu/null/v5"

)

func InsertCategory(c *fiber.Ctx) error {
	type Request struct {
		CategoryName string `json:"category_name" validate:"required,max=255"`
		SuperCatID   null.Int   `json:"super_cat_id"`
		Col1         string `json:"col1"`
		Col2         string `json:"col2"`
		Remarks      string `json:"remarks"`
	}

	var req Request
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if req.SuperCatID.Int64 == -1 {
		req.SuperCatID = null.IntFromPtr(nil)
	}

	_, err := db.Pool.Exec(context.Background(), `
		CALL insert_category($1, $2);
	`, req.CategoryName, req.SuperCatID)

	if err != nil {
		log.Printf("Failed to insert category: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to insert category"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Category added successfully"})
}

func UpdateCategory(c *fiber.Ctx) error {
	type Request struct {
		CategoryID   int    `json:"category_id" validate:"required,min=1"`
		CategoryName string `json:"category_name" validate:"required,max=255"`
		SuperCatID   *int   `json:"super_cat_id"`
		Col1         string `json:"col1"`
		Col2         string `json:"col2"`
		Remarks      string `json:"remarks"`
	}

	var req Request
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	_, err := db.Pool.Exec(context.Background(), `
		CALL update_category($1, $2, $3);
	`, req.CategoryID, req.CategoryName, req.SuperCatID)

	if err != nil {
		log.Printf("Failed to update category: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update category"})
	}

	return c.JSON(fiber.Map{"message": "Category updated successfully"})
}

func DeleteCategory(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ID is required"})
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID format"})
	}

	_, err = db.Pool.Exec(context.Background(), `
		CALL delete_category($1);
	`, idInt)

	if err != nil {
		log.Printf("Failed to delete category: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete category"})
	}

	return c.JSON(fiber.Map{"message": "Category deleted successfully"})
}

func GetCategories(c *fiber.Ctx) error {
	rows, err := db.Pool.Query(context.Background(), "SELECT * FROM get_categories();")
	if err != nil {
		log.Printf("Failed to fetch categories: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch categories"})
	}
	defer rows.Close()

	type Category struct {
		CategoryID   int    `json:"category_id"`
		CategoryName string `json:"category_name"`
		SuperCatID   int    `json:"super_cat_id"`
		Remarks      string `json:"remarks"`
	}

	categories := make([]Category, 0)
	for rows.Next() {
		var category Category
		var superCatID sql.NullInt32
		var remarks sql.NullString

		if err := rows.Scan(&category.CategoryID, &category.CategoryName, &superCatID, &remarks); err != nil {
			log.Printf("Failed to scan category: %v", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to process categories"})
		}

		category.SuperCatID = int(superCatID.Int32) 
		if !remarks.Valid {
			category.Remarks = "No remarks" 
		} else {
			category.Remarks = remarks.String
		}

		categories = append(categories, category)
	}

	return c.JSON(categories)
}
