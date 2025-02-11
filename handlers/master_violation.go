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
func GetViolations(c *fiber.Ctx) error {
	rows, err := db.Pool.Query(context.Background(), "SELECT * FROM get_master_violations()")
	if err != nil {
		log.Printf("Failed to fetch violation records: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch violations"})
	}
	defer rows.Close()

	var violations []map[string]interface{}

	for rows.Next() {
		var id, status *int
		var violationName, levelOfSerious *string

		if err := rows.Scan(&id, &violationName, &levelOfSerious, &status); err != nil {
			log.Printf("Error scanning row: %v", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error processing data"})
		}

		violations = append(violations, map[string]interface{}{
			"id":               id,
			"violation_name":   violationName,
			"level_of_serious": levelOfSerious,
			"status":           status,
		})
	}

	return c.JSON(violations)
}