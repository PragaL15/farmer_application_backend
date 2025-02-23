package handlers

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/PragaL15/go_newBackend/go_backend/db"
)

// BusinessType represents the structure of the business_type_table
type BusinessType struct {
	BTypeID   int    `json:"b_typeid"`
	BTypeName string `json:"b_typename"`
	Remarks   string `json:"remarks"`
}

// GetBusinessTypes retrieves all business types
func GetBusinessTypes(c *fiber.Ctx) error {
	rows, err := db.Pool.Query(context.Background(), "SELECT * FROM get_business_types()")
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

// InsertBusinessType inserts a new business type
func InsertBusinessType(c *fiber.Ctx) error {
	type Request struct {
		BTypeName string `json:"b_typename"`
		Remarks   string `json:"remarks"`
	}
	var req Request
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}
	_, err := db.Pool.Exec(context.Background(), "SELECT insert_business_type($1, $2)", req.BTypeName, req.Remarks)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "Business type inserted successfully"})
}

// UpdateBusinessType updates an existing business type
func UpdateBusinessType(c *fiber.Ctx) error {
	type Request struct {
		BTypeID   int    `json:"b_typeid"`
		BTypeName string `json:"b_typename"`
		Remarks   string `json:"remarks"`
	}
	var req Request
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}
	_, err := db.Pool.Exec(context.Background(), "SELECT update_business_type($1, $2, $3)", req.BTypeID, req.BTypeName, req.Remarks)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "Business type updated successfully"})
}
