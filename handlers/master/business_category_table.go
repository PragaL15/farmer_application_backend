package Masterhandlers

import (
	"context"
	"strconv"

	"github.com/PragaL15/go_newBackend/go_backend/db"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v4"
)

type BusinessCategory struct {
	BCategoryID   int    `json:"b_category_id"`
	BCategoryName string `json:"b_category_name"`
}

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

func InsertBusinessCategory(c *fiber.Ctx) error {
	type Request struct {
		BCategoryName string `json:"b_category_name"`
	}

	var req Request
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

func UpdateBusinessCategory(c *fiber.Ctx) error {
	type Request struct {
		BCategoryID   int    `json:"b_category_id"`
		BCategoryName string `json:"b_category_name"`
	}

	var req Request
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
