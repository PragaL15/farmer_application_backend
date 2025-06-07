package Masterhandlers

import (
	"context"
	"farmerapp/go_backend/db"
	"log"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// Insert City
func InsertCity(c *fiber.Ctx) error {
	type Request struct {
		CityShortNames string `json:"city_shortnames" validate:"required,max=5"`
		CityName       string `json:"city_name" validate:"required,max=50"`
	}

	var req Request
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	_, err := db.Pool.Exec(context.Background(), "CALL admin_schema.insert_city($1, $2);", req.CityShortNames, req.CityName)
	if err != nil {
		log.Printf("Failed to insert city: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to insert city"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "City added successfully"})
}

// Get All Cities
func GetCities(c *fiber.Ctx) error {
	rows, err := db.Pool.Query(context.Background(), "SELECT * FROM admin_schema.get_all_cities();")
	if err != nil {
		log.Printf("Failed to fetch cities: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch cities"})
	}
	defer rows.Close()

	type City struct {
		ID             int    `json:"id"`
		CityShortNames string `json:"city_shortnames"`
		CityName       string `json:"city_name"`
	}

	cities := make([]City, 0)
	for rows.Next() {
		var city City
		if err := rows.Scan(&city.ID, &city.CityShortNames, &city.CityName); err != nil {
			log.Printf("Failed to scan city: %v", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to process cities"})
		}
		cities = append(cities, city)
	}

	return c.JSON(cities)
}

// Get City by ID
func GetCityByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ID is required"})
	}
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID format"})
	}

	row := db.Pool.QueryRow(context.Background(), "SELECT * FROM admin_schema.get_city_by_id($1);", idInt)
	type City struct {
		ID             int    `json:"id"`
		CityShortNames string `json:"city_shortnames"`
		CityName       string `json:"city_name"`
	}

	var city City
	if err := row.Scan(&city.ID, &city.CityShortNames, &city.CityName); err != nil {
		log.Printf("Failed to fetch city: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch city"})
	}

	return c.JSON(city)
}

// Update City
func UpdateCity(c *fiber.Ctx) error {
	type Request struct {
		ID             int    `json:"id" validate:"required,min=1"`
		CityShortNames string `json:"city_shortnames" validate:"required,max=5"`
		CityName       string `json:"city_name" validate:"required,max=50"`
	}

	var req Request
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	_, err := db.Pool.Exec(context.Background(), "CALL admin_schema.update_city($1, $2, $3);", req.ID, req.CityShortNames, req.CityName)
	if err != nil {
		log.Printf("Failed to update city: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update city"})
	}

	return c.JSON(fiber.Map{"message": "City updated successfully"})
}
