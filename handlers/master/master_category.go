package Masterhandlers

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/PragaL15/go_newBackend/go_backend/db"
	"github.com/go-playground/validator/v10"
	"log"
"strconv"
	"github.com/guregu/null/v5"
)

func InsertCategory(c *fiber.Ctx) error {
	type Request struct {
		CategoryName       string    `json:"category_name" validate:"required,max=255"`
		SuperCatID         null.Int  `json:"super_cat_id"`
		ImgPath            *string   `json:"img_path"`
		ActiveStatus       *int      `json:"active_status"`
		CategoryRegionalID *int      `json:"category_regional_id"`
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
		SELECT admin_schema.insert_product_category($1, $2, $3, $4, $5);
	`, req.CategoryName, req.SuperCatID, req.ImgPath, req.ActiveStatus, req.CategoryRegionalID)

	if err != nil {
		log.Printf("Failed to insert category: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to insert category"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Category added successfully"})
}

func UpdateCategory(c *fiber.Ctx) error {
	type Request struct {
		CategoryID         int      `json:"category_id" validate:"required,min=1"`
		CategoryName       string   `json:"category_name" validate:"required,max=255"`
		SuperCatID         *int     `json:"super_cat_id"`
		ImgPath            *string  `json:"img_path"`
		ActiveStatus       *int     `json:"active_status"`
		CategoryRegionalID *int     `json:"category_regional_id"`
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
		SELECT admin_schema.update_product_category($1, $2, $3, $4, $5, $6);
	`, req.CategoryID, req.CategoryName, req.SuperCatID, req.ImgPath, req.ActiveStatus, req.CategoryRegionalID)

	if err != nil {
		log.Printf("Failed to update category: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update category"})
	}

	return c.JSON(fiber.Map{"message": "Category updated successfully"})
}

func GetCategoryByID(c *fiber.Ctx) error {
	categoryID := c.Params("category_id")
	if categoryID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Category ID is required"})
	}

	row := db.Pool.QueryRow(context.Background(), "SELECT admin_schema.get_category_with_subcategories($1);", categoryID)

	type Category struct {
		CategoryID         int     `json:"category_id"`
		CategoryName       string  `json:"category_name"`
		SuperCatID         *int    `json:"super_cat_id"`
		ImgPath            *string `json:"img_path"`
		ActiveStatus       *int    `json:"active_status"`
		CategoryRegionalID *int    `json:"category_regional_id"`
	}

	var category Category
	err := row.Scan(
		&category.CategoryID,
		&category.CategoryName,
		&category.SuperCatID,
		&category.ImgPath,
		&category.ActiveStatus,
		&category.CategoryRegionalID,
	)

	if err != nil {
		log.Printf("Failed to fetch category: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch category"})
	}

	return c.JSON(category)
}

func GetCategoriesBySuperCatID(c *fiber.Ctx) error {
	superCatIDParam := c.Params("super_cat_id")
	if superCatIDParam == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "super_cat_id is required"})
	}

	id, err := strconv.Atoi(superCatIDParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid super_cat_id"})
	}

	rows, err := db.Pool.Query(context.Background(), "SELECT * FROM admin_schema.get_categories_by_super_cat_id($1);", id)
	if err != nil {
		log.Printf("DB query failed: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch categories"})
	}
	defer rows.Close()

	type Category struct {
		CategoryID   int    `json:"category_id"`
		CategoryName string `json:"category_name"`
	}

	var categories []Category

	for rows.Next() {
		var (
			categoryID   int
			categoryName string
		)

		err := rows.Scan(&categoryID, &categoryName)
		if err != nil {
			log.Printf("Row scan error: %v", err)
			continue
		}

		categories = append(categories, Category{
			CategoryID:   categoryID,
			CategoryName: categoryName,
		})
	}

	return c.JSON(categories)
}

func GetSuperCategories(c *fiber.Ctx) error {
	rows, err := db.Pool.Query(context.Background(), "SELECT * FROM admin_schema.get_supercategories();")
	if err != nil {
		log.Printf("DB query failed: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch supercategories"})
	}
	defer rows.Close()

	type Category struct {
		CategoryID         int     `json:"category_id"`
		CategoryName       string  `json:"category_name"`
		SuperCatID         *int    `json:"super_cat_id"`
		ImgPath            *string `json:"img_path"`
		ActiveStatus       *int    `json:"active_status"`
		CategoryRegionalID *int    `json:"category_regional_id"`
	}

	var categories []Category

	for rows.Next() {
		var cat Category
		err := rows.Scan(
			&cat.CategoryID,
			&cat.CategoryName,
			&cat.SuperCatID,
			&cat.ImgPath,
			&cat.ActiveStatus,
			&cat.CategoryRegionalID,
		)
		if err != nil {
			log.Printf("Row scan error: %v", err)
			continue
		}
		categories = append(categories, cat)
	}

	return c.JSON(categories)
}
