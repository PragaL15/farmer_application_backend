package handlers

import (
	"context"
	"log"
	"strconv"
"time"
	"github.com/gofiber/fiber/v2"
	"github.com/PragaL15/go_newBackend/go_backend/db"
	"github.com/go-playground/validator/v10"
)

func InsertDriver(c *fiber.Ctx) error {
	type Request struct {
		DriverName       string  `json:"driver_name" validate:"required,max=255"`
		DriverAge        *int    `json:"driver_age"`
		DriverLicense    string  `json:"driver_license" validate:"required,max=50"`
		DriverNumber     string  `json:"driver_number" validate:"required,max=15"`
		DriverAddress    *string `json:"driver_address"`
		DriverStatus     *string `json:"driver_status"`
		DateOfJoining    *string `json:"date_of_joining"`
		ExperienceYears  *int    `json:"experience_years"`
		VehicleID        *int    `json:"vehicle_id"`
		LicenseExpiry    string  `json:"license_expiry_date" validate:"required"`
		EmergencyContact *string `json:"emergency_contact"`
		AssignedRouteID  *int    `json:"assigned_route_id"`
		Col1             *string `json:"col1"`
		Col2             *string `json:"col2"`
		DOB              *string `json:"d_o_b"`
		Violation        *int    `json:"violation"`
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
		CALL insert_driver($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16);
	`, req.DriverName, req.DriverAge, req.DriverLicense, req.DriverNumber, req.DriverAddress, req.DriverStatus,
		req.DateOfJoining, req.ExperienceYears, req.VehicleID, req.LicenseExpiry, req.EmergencyContact,
		req.AssignedRouteID, req.Col1, req.Col2, req.DOB, req.Violation)

	if err != nil {
		log.Printf("Failed to insert driver: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to insert driver"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Driver added successfully"})
}

func UpdateDriver(c *fiber.Ctx) error {
	type Request struct {
		DriverID         int     `json:"driver_id" validate:"required,min=1"`
		DriverName       *string `json:"driver_name"`
		DriverAge        *int    `json:"driver_age"`
		DriverLicense    *string `json:"driver_license"`
		DriverNumber     *string `json:"driver_number"`
		DriverAddress    *string `json:"driver_address"`
		DriverStatus     *string `json:"driver_status"`
		DateOfJoining    *string `json:"date_of_joining"`
		ExperienceYears  *int    `json:"experience_years"`
		VehicleID        *int    `json:"vehicle_id"`
		LicenseExpiry    *string `json:"license_expiry_date"`
		EmergencyContact *string `json:"emergency_contact"`
		AssignedRouteID  *int    `json:"assigned_route_id"`
		Col1             *string `json:"col1"`
		Col2             *string `json:"col2"`
		DOB              *string `json:"d_o_b"`
		Violation        *int    `json:"violation"`
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
		CALL update_driver($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17);
	`, req.DriverID, req.DriverName, req.DriverAge, req.DriverLicense, req.DriverNumber, req.DriverAddress,
		req.DriverStatus, req.DateOfJoining, req.ExperienceYears, req.VehicleID, req.LicenseExpiry, req.EmergencyContact,
		req.AssignedRouteID, req.Col1, req.Col2, req.DOB, req.Violation)

	if err != nil {
		log.Printf("Failed to update driver: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update driver"})
	}

	return c.JSON(fiber.Map{"message": "Driver updated successfully"})
}

func DeleteDriver(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Driver ID is required"})
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID format"})
	}

	_, err = db.Pool.Exec(context.Background(), `
		CALL delete_driver($1);
	`, idInt)

	if err != nil {
		log.Printf("Failed to delete driver: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete driver"})
	}

	return c.JSON(fiber.Map{"message": "Driver deleted successfully"})
}

func GetDrivers(c *fiber.Ctx) error {
	rows, err := db.Pool.Query(context.Background(), "SELECT * FROM get_master_drivers()")
	if err != nil {
		log.Printf("Failed to fetch drivers: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch drivers"})
	}
	defer rows.Close()

	var drivers []map[string]interface{}

	for rows.Next() {
		var driverID, driverAge, experienceYears, vehicleID, assignedRouteID, violation *int
		var driverName, driverLicense, driverNumber, driverAddress, driverStatus, emergencyContact, col1, col2 *string
		var dateOfJoining, licenseExpiryDate, d_o_b *time.Time
		var createdAt, updatedAt time.Time

		if err := rows.Scan(
			&driverID, &driverName, &driverAge, &driverLicense, &driverNumber,
			&driverAddress, &driverStatus, &dateOfJoining, &experienceYears, &vehicleID,
			&licenseExpiryDate, &emergencyContact, &assignedRouteID, &createdAt, &col1, &col2,
			&d_o_b, &violation, &updatedAt,
		); err != nil {
			log.Printf("Error scanning row: %v", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error processing data"})
		}

		driver := map[string]interface{}{
			"driver_id":           driverID,
			"driver_name":         driverName,
			"driver_age":          driverAge,
			"driver_license":      driverLicense,
			"driver_number":       driverNumber,
			"driver_address":      driverAddress,
			"driver_status":       driverStatus,
			"date_of_joining":     formatDate(dateOfJoining),
			"experience_years":    experienceYears,
			"vehicle_id":          vehicleID,
			"license_expiry_date": formatDate(licenseExpiryDate),
			"emergency_contact":   emergencyContact,
			"assigned_route_id":   assignedRouteID,
			"created_at":          createdAt.Format(time.RFC3339),
			"col1":                col1,
			"col2":                col2,
			"d_o_b":               formatDate(d_o_b),
			"violation":           violation,
			"updated_at":          updatedAt.Format(time.RFC3339),
		}

		drivers = append(drivers, driver)
	}

	return c.JSON(drivers)
}

// formatDate handles nil time values
func formatDate(t *time.Time) interface{} {
	if t == nil {
		return nil
	}
	return t.Format("2006-01-02")
}
