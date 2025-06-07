package Masterhandlers

import (
	"context"
	"farmerapp/go_backend/db"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type State struct {
	ID              int64  `json:"id"`
	State           string `json:"state"`
	StateShortNames string `json:"state_shortnames"`
}

func GetAllStates(c *fiber.Ctx) error {
	rows, err := db.Pool.Query(context.Background(), "SELECT * FROM admin_schema.get_all_states()")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch states", "details": err.Error()})
	}
	defer rows.Close()
	var states []State
	for rows.Next() {
		var state State
		if err := rows.Scan(&state.ID, &state.State, &state.StateShortNames); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error scanning state data", "details": err.Error()})
		}
		states = append(states, state)
	}
	return c.JSON(states)
}

func GetStateByID(c *fiber.Ctx) error {
	stateID, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid state ID format"})
	}
	var state State
	err = db.Pool.QueryRow(context.Background(), "SELECT * FROM admin_schema.get_state_by_id($1)", stateID).
		Scan(&state.ID, &state.State, &state.StateShortNames)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "State not found"})
	}
	return c.JSON(state)
}

func InsertState(c *fiber.Ctx) error {
	var state State
	if err := c.BodyParser(&state); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}
	_, err := db.Pool.Exec(context.Background(), "SELECT admin_schema.insert_state($1, $2)", state.State, state.StateShortNames)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to insert state", "details": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "State inserted successfully"})
}

func UpdateState(c *fiber.Ctx) error {
	stateID, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid state ID format"})
	}
	var state State
	if err := c.BodyParser(&state); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	_, err = db.Pool.Exec(context.Background(), "SELECT admin_schema.update_state($1, $2, $3)", stateID, state.State, state.StateShortNames)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update state", "details": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "State updated successfully"})
}
