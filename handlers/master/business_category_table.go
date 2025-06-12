package Masterhandlers

import (
	"context"
	"strconv"

	"farmerapp/go_backend/db"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v4"
)

// BusinessCategory represents a business category
type BusinessCategory struct {
	BCategoryID   int    `json:"b_category_id"`
	BCategoryName string `json:"b_category_name"`
}

// InsertBusinessCategoryRequest is the request body for inserting a category
type InsertBusinessCategoryRequest struct {
	BCategoryName string `json:"b_category_name"`
}

// UpdateBusinessCategoryRequest is the request body for updating a category
type UpdateBusinessCategoryRequest struct {
	BCategoryID   int    `json:"b_category_id"`
	BCategoryName string `json:"b_category_name"`
}

// GetBusinessCategories godoc
// @Summary Get all or specific business categories
// @Description Get all business categories or a specific one if ID is passed
// @Tags BusinessCategory
// @Accept json
// @Produce json
// @Param id query int false "Business Category ID"
// @Success 200 {array} BusinessCategory
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /business-category [get]
func GetBusinessCategories(c *fiber.Ctx) error {
	id := c.Query("id")

	var rows pgx.Rows
	var err error

	if id == "" {
		rows, err = db.Pool.Query(context.Background(), "SELECT * FROM admin_schema.get_business_category()")
	} else {
		idInt, convErr := strconv.Atoi(id)
		if convErr != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID format"})
		}
		rows, err = db.Pool.Query(context.Background(), "SELECT * FROM admin_schema.get_business_category($1)", idInt)
	}

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	defer rows.Close()

	var categories []BusinessCategory
	for rows.Next() {
		var bc BusinessCategory
		if err := rows.Scan(&bc.BCategoryID, &bc.BCategoryName); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		categories = append(categories, bc)
	}

	return c.JSON(categories)
}

// InsertBusinessCategory godoc
// @Summary Insert a new business category
// @Description Add a new business category name
// @Tags BusinessCategory
// @Accept json
// @Produce json
// @Param data body InsertBusinessCategoryRequest true "Business Category Name"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /business-category [post]
func InsertBusinessCategory(c *fiber.Ctx) error {
	var req InsertBusinessCategoryRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	var newID int
	err := db.Pool.QueryRow(context.Background(), "SELECT admin_schema.insert_business_category($1)", req.BCategoryName).Scan(&newID)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Business category inserted successfully", "id": newID})
}

// UpdateBusinessCategory godoc
// @Summary Update a business category
// @Description Update the name of a business category using its ID
// @Tags BusinessCategory
// @Accept json
// @Produce json
// @Param data body UpdateBusinessCategoryRequest true "Updated Info"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /business-category [put]
func UpdateBusinessCategory(c *fiber.Ctx) error {
	var req UpdateBusinessCategoryRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	var updated bool
	err := db.Pool.QueryRow(context.Background(), "SELECT admin_schema.update_business_category($1, $2)", req.BCategoryID, req.BCategoryName).Scan(&updated)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	if !updated {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "No rows updated"})
	}

	return c.JSON(fiber.Map{"message": "Business category updated successfully"})
}
