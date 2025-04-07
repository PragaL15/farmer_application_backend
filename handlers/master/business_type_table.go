package Masterhandlers

import (
	"context"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/PragaL15/go_newBackend/go_backend/db"
)

type BusinessType struct {
	BTypeID   int    `json:"b_typeid"`
	BTypeName string `json:"b_typename"`
	Remarks   string `json:"remarks"`
}

// Get all business types
func GetBusinessTypes(c *fiber.Ctx) error {
	rows, err := db.Pool.Query(context.Background(), "SELECT * FROM admin_schema.get_business_types()")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	defer rows.Close()

	var businessTypes []BusinessType
	for rows.Next() {
		var bt BusinessType
		if err := rows.Scan(&bt.BTypeID, &bt.BTypeName, &bt.Remarks); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		businessTypes = append(businessTypes, bt)
	}
	return c.JSON(businessTypes)
}

// Get business type by ID
func GetBusinessTypeByID(c *fiber.Ctx) error {
	id := c.Params("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID format"})
	}

	var bTypeName string
	var remarks string

	err = db.Pool.QueryRow(context.Background(), "SELECT * FROM admin_schema.get_business_type_by_id($1)", idInt).Scan(&bTypeName, &remarks)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{
		"b_typeid":   idInt,
		"b_typename": bTypeName,
		"remarks":    remarks,
	})
}

// Insert new business type
func InsertBusinessType(c *fiber.Ctx) error {
	var req BusinessType
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	var inserted BusinessType
	err := db.Pool.QueryRow(
		context.Background(),
		"SELECT * FROM admin_schema.insert_business_type($1, $2)",
		req.BTypeName, req.Remarks,
	).Scan(&inserted.BTypeID, &inserted.BTypeName, &inserted.Remarks)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to insert", "details": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(inserted)
}


func UpdateBusinessType(c *fiber.Ctx) error {
	var req BusinessType
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	_, err := db.Pool.Exec(context.Background(), "CALL admin_schema.update_business_type($1, $2, $3)", req.BTypeID, req.BTypeName, req.Remarks)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update", "details": err.Error()})
	}

	return c.JSON(fiber.Map{
		"message": "Business type updated successfully",
	})
}
