package Masterhandlers

import (
	"context"
	"log"
	"strconv"

	"farmerapp/go_backend/db"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func InsertLocation(c *fiber.Ctx) error {
	type Request struct {
		Location       string `json:"location" validate:"required,max=50"`
		CityShortnames int    `json:"city_shortnames" validate:"required"`
		State          int    `json:"state" validate:"required"`
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
		SELECT insert_location($1, $2, $3);
	`, req.CityShortnames, req.State, req.Location)

	if err != nil {
		log.Printf("Failed to insert location: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to insert location"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Location added successfully"})
}

func UpdateLocation(c *fiber.Ctx) error {
	type Request struct {
		ID             int    `json:"id" validate:"required"`
		Location       string `json:"location" validate:"required,max=50"`
		CityShortnames int    `json:"city_shortnames" validate:"required"`
		State          int    `json:"state" validate:"required"`
	}

	var req Request
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	_, err := db.Pool.Exec(context.Background(),
		`SELECT update_location($1, $2, $3, $4)`,
		req.ID, req.CityShortnames, req.State, req.Location,
	)

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
		var (
			id        int
			location  *string
			cityID    int
			cityName  *string
			stateID   int
			stateName *string
		)

		if err := rows.Scan(&id, &location, &cityID, &cityName, &stateID, &stateName); err != nil {
			log.Printf("Error scanning row: %v", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error processing data"})
		}

		locations = append(locations, map[string]interface{}{
			"id":         id,
			"location":   location,
			"city_id":    cityID,
			"city_name":  cityName,
			"state_id":   stateID,
			"state_name": stateName,
		})
	}

	return c.JSON(locations)
}
