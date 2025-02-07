package handlers

import (
	"context"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/PragaL15/go_newBackend/go_backend/db"
	"github.com/go-playground/validator/v10"
)

func InsertMasterViolation(c *fiber.Ctx) error {
	type Request struct {
		ViolationName   string `json:"violation_name" validate:"required,max=255"`
		LevelOfSerious  string `json:"level_of_serious" validate:"required,max=255"`
		Status          int    `json:"status" validate:"required"`
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
		CALL insert_master_violation($1, $2, $3);
	`, req.ViolationName, req.LevelOfSerious, req.Status)

	if err != nil {
		log.Printf("Failed to insert violation: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to insert violation"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Violation added successfully"})
}

func UpdateMasterViolation(c *fiber.Ctx) error {
	type Request struct {
		ID              int    `json:"id" validate:"required,min=1"`
		ViolationName   string `json:"violation_name" validate:"required,max=255"`
		LevelOfSerious  string `json:"level_of_serious" validate:"required,max=255"`
		Status          int    `json:"status" validate:"required"`
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
		CALL update_master_violation($1, $2, $3, $4);
	`, req.ID, req.ViolationName, req.LevelOfSerious, req.Status)

	if err != nil {
		log.Printf("Failed to update violation: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update violation"})
	}

	return c.JSON(fiber.Map{"message": "Violation updated successfully"})
}
func DeleteMasterViolation(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ID is required"})
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID format"})
	}

	_, err = db.Pool.Exec(context.Background(), `
		CALL delete_master_violation($1);
	`, idInt)

	if err != nil {
		log.Printf("Failed to delete violation: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete violation"})
	}

	return c.JSON(fiber.Map{"message": "Violation deleted successfully"})
}
