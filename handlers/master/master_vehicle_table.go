package Masterhandlers

import (
	"context"
	"database/sql"
	"farmerapp/go_backend/db"
	"log"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

type Vehicle struct {
	VehicleID             int64         `json:"vehicle_id"`
	VehicleName           string        `json:"vehicle_name"`
	VehicleRegistrationNo string        `json:"vehicle_registration_no"`
	VehicleYear           string        `json:"vehicle_manufacture_year"`
	VehicleWarranty       string        `json:"vehicle_warranty"`
	VehicleMake           sql.NullInt64 `json:"vehicle_make"`
	VehicleModel          sql.NullInt64 `json:"vehicle_model"`
	VehicleEngineType     sql.NullInt64 `json:"vehicle_engine_type"`
	VehiclePurchaseDate   time.Time     `json:"vehicle_purchase_date"`
	VehicleColor          string        `json:"vehicle_color"`
	CreatedAt             time.Time     `json:"created_at"`
	UpdatedAt             time.Time     `json:"updated_at"`
}

func GetAllVehicles(c *fiber.Ctx) error {
	rows, err := db.Pool.Query(context.Background(), "SELECT * FROM admin_schema.get_all_vehicles()")
	if err != nil {
		log.Println("Error fetching vehicles:", err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch vehicles"})
	}
	defer rows.Close()

	var vehicles []Vehicle
	for rows.Next() {
		var v Vehicle
		if err := rows.Scan(
			&v.VehicleID, &v.VehicleName, &v.VehicleRegistrationNo, &v.VehicleYear,
			&v.VehicleWarranty, &v.VehicleMake, &v.VehicleModel, &v.VehicleEngineType,
			&v.VehiclePurchaseDate, &v.VehicleColor, &v.CreatedAt, &v.UpdatedAt,
		); err != nil {
			log.Println("Error scanning row:", err)
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to parse data"})
		}
		vehicles = append(vehicles, v)
	}
	return c.JSON(vehicles)
}

func GetVehicleByID(c *fiber.Ctx) error {
	vehicleID := c.Params("id")
	var v Vehicle

	err := db.Pool.QueryRow(context.Background(), "SELECT * FROM admin_schema.get_vehicle_by_id($1)", vehicleID).
		Scan(
			&v.VehicleID, &v.VehicleName, &v.VehicleRegistrationNo, &v.VehicleYear,
			&v.VehicleWarranty, &v.VehicleMake, &v.VehicleModel, &v.VehicleEngineType,
			&v.VehiclePurchaseDate, &v.VehicleColor, &v.CreatedAt, &v.UpdatedAt,
		)
	if err != nil {
		log.Println("Error fetching vehicle:", err)
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Vehicle not found"})
	}
	return c.JSON(v)
}

func InsertVehicle(c *fiber.Ctx) error {
	var v Vehicle
	if err := c.BodyParser(&v); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	_, err := db.Pool.Exec(
		context.Background(),
		"CALL admin_schema.insert_vehicle($1, $2, $3, $4, $5, $6, $7, $8, $9)",
		v.VehicleName, v.VehicleRegistrationNo, v.VehicleYear, v.VehicleWarranty,
		v.VehicleMake, v.VehicleModel, v.VehicleEngineType, v.VehiclePurchaseDate, v.VehicleColor,
	)
	if err != nil {
		log.Println("Error inserting vehicle:", err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to insert vehicle"})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{"message": "Vehicle created successfully"})
}

func UpdateVehicle(c *fiber.Ctx) error {
	vehicleID := c.Params("id")
	var v Vehicle
	if err := c.BodyParser(&v); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	_, err := db.Pool.Exec(
		context.Background(),
		"CALL admin_schema.update_vehicle($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)",
		vehicleID, v.VehicleName, v.VehicleRegistrationNo, v.VehicleYear, v.VehicleWarranty,
		v.VehicleMake, v.VehicleModel, v.VehicleEngineType, v.VehiclePurchaseDate, v.VehicleColor,
	)
	if err != nil {
		log.Println("Error updating vehicle:", err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update vehicle"})
	}

	return c.JSON(fiber.Map{"message": "Vehicle updated successfully"})
}
