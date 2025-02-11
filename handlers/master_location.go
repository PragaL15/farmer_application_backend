package handlers

import (
	"context"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/PragaL15/go_newBackend/go_backend/db"
	"github.com/go-playground/validator/v10"
)

func InsertLocation(c *fiber.Ctx) error {
	type Request struct {
		Location string `json:"location" validate:"required,max=50"`
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
		CALL insert_location($1);
	`, req.Location)

	if err != nil {
		log.Printf("Failed to insert location: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to insert location"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Location added successfully"})
}

func UpdateLocation(c *fiber.Ctx) error {
	type Request struct {
		ID       int    `json:"id" validate:"required,min=1"`
		Location string `json:"location" validate:"required,max=50"`
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
		CALL update_location($1, $2);
	`, req.ID, req.Location)

	if err != nil {
		log.Printf("Failed to update location: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update location"})
	}

	return c.JSON(fiber.Map{"message": "Location updated successfully"})
}

func DeleteLocation(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ID is required"})
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID format"})
	}

	_, err = db.Pool.Exec(context.Background(), `
		CALL delete_location($1);
	`, idInt)

	if err != nil {
		log.Printf("Failed to delete location: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete location"})
	}

	return c.JSON(fiber.Map{"message": "Location deleted successfully"})
}


func GetLocations(c *fiber.Ctx) error {
	rows, err := db.Pool.Query(context.Background(), "SELECT * FROM get_master_locations()")
	if err != nil {
		log.Printf("Failed to fetch locations: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch locations"})
	}
	defer rows.Close()

	var locations []map[string]interface{}

	for rows.Next() {
		var id int
		var location *string

		if err := rows.Scan(&id, &location); err != nil {
			log.Printf("Error scanning row: %v", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error processing data"})
		}

		locations = append(locations, map[string]interface{}{
			"id":       id,
			"location": location,
		})
	}

	return c.JSON(locations)
}