package handlers

import (
	"context"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/PragaL15/go_newBackend/go_backend/db"
	"github.com/go-playground/validator/v10"
)

func InsertMasterState(c *fiber.Ctx) error {
	type Request struct {
		State string `json:"state" validate:"required,max=50"`
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
		CALL insert_master_state($1);
	`, req.State)

	if err != nil {
		log.Printf("Failed to insert state: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to insert state"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "State added successfully"})
}
func UpdateMasterState(c *fiber.Ctx) error {
	type Request struct {
		ID    int    `json:"id" validate:"required,min=1"`
		State string `json:"state" validate:"required,max=50"`
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
		CALL update_master_state($1, $2);
	`, req.ID, req.State)

	if err != nil {
		log.Printf("Failed to update state: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update state"})
	}

	return c.JSON(fiber.Map{"message": "State updated successfully"})
}
func DeleteMasterState(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ID is required"})
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID format"})
	}

	_, err = db.Pool.Exec(context.Background(), `
		CALL delete_master_state($1);
	`, idInt)

	if err != nil {
		log.Printf("Failed to delete state: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete state"})
	}

	return c.JSON(fiber.Map{"message": "State deleted successfully"})
}

func GetStates(c *fiber.Ctx) error {
	rows, err := db.Pool.Query(context.Background(), "SELECT * FROM get_master_states()")
	if err != nil {
		log.Printf("Failed to fetch state records: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch states"})
	}
	defer rows.Close()

	var states []map[string]interface{}

	for rows.Next() {
		var id *int
		var stateName *string

		if err := rows.Scan(&id, &stateName); err != nil {
			log.Printf("Error scanning row: %v", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error processing data"})
		}

		states = append(states, map[string]interface{}{
			"id":    id,
			"state": stateName,
		})
	}

	return c.JSON(states)
}