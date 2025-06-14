package Masterhandlers

import (
	"context"
	"database/sql"
	"farmerapp/go_backend/db"
	"log"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func InsertDriver(c *fiber.Ctx) error {
	type Request struct {
		DriverName       string `json:"driver_name" validate:"required,max=255"`
		DriverAge        int    `json:"driver_age"`
		DriverLicense    string `json:"driver_license" validate:"required,max=50"`
		DriverNumber     string `json:"driver_number" validate:"required,max=15"`
		DriverAddress    string `json:"driver_address"`
		DriverStatus     string `json:"driver_status"`
		DateOfJoining    string `json:"date_of_joining"`
		ExperienceYears  int    `json:"experience_years"`
		VehicleID        int    `json:"vehicle_id"`
		LicenseExpiry    string `json:"license_expiry_date" validate:"required"`
		EmergencyContact string `json:"emergency_contact"`
		AssignedRouteID  int    `json:"assigned_route_id"`
		Col1             string `json:"col1"`
		Col2             string `json:"col2"`
		DOB              string `json:"d_o_b"`
		Violation        int    `json:"violation"`
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
		DriverID         int    `json:"driver_id" validate:"required,min=1"`
		DriverName       string `json:"driver_name"`
		DriverAge        int    `json:"driver_age"`
		DriverLicense    string `json:"driver_license"`
		DriverNumber     string `json:"driver_number"`
		DriverAddress    string `json:"driver_address"`
		DriverStatus     string `json:"driver_status"`
		DateOfJoining    string `json:"date_of_joining"`
		ExperienceYears  int    `json:"experience_years"`
		VehicleID        int    `json:"vehicle_id"`
		LicenseExpiry    string `json:"license_expiry_date"`
		EmergencyContact string `json:"emergency_contact"`
		AssignedRouteID  int    `json:"assigned_route_id"`
		Col1             string `json:"col1"`
		Col2             string `json:"col2"`
		DOB              string `json:"d_o_b"`
		Violation        int    `json:"violation"`
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
	type Driver struct {
		DriverID         int           `json:"driver_id"`
		DriverName       string        `json:"driver_name"`
		DriverLicense    string        `json:"driver_license"`
		DriverNumber     string        `json:"driver_number"`
		DriverAddress    string        `json:"driver_address"`
		DateOfJoining    *string       `json:"date_of_joining"`
		EmergencyContact string        `json:"emergency_contact"`
		CreatedAt        *string       `json:"created_at"`
		DOB              *string       `json:"d_o_b"`
		LicenseTypeID    sql.NullInt32 `json:"driver_license_type"`
	}

	rows, err := db.Pool.Query(context.Background(), "SELECT * FROM admin_schema.get_all_drivers()")
	if err != nil {
		log.Printf("Failed to fetch drivers: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch drivers"})
	}
	defer rows.Close()

	var drivers []Driver

	for rows.Next() {
		var driver Driver

		var (
			doj           sql.NullTime
			created       sql.NullTime
			dob           sql.NullTime
			driverAddress sql.NullString
			emergency     sql.NullString
		)

		err := rows.Scan(
			&driver.DriverID, &driver.DriverName, &driver.DriverLicense,
			&driver.DriverNumber, &driverAddress, &doj,
			&emergency, &created, &dob, &driver.LicenseTypeID,
		)
		if err != nil {
			log.Printf("Error scanning row: %v", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error processing data"})
		}

		driver.DriverAddress = nullStringToString(driverAddress)
		driver.EmergencyContact = nullStringToString(emergency)
		driver.DateOfJoining = formatNullableTime(doj)
		driver.CreatedAt = formatNullableTime(created)
		driver.DOB = formatNullableTime(dob)

		drivers = append(drivers, driver)
	}

	return c.JSON(drivers)
}

// formatNullableTime formats a sql.NullTime to *string
func formatNullableTime(t sql.NullTime) *string {
	if t.Valid {
		formatted := t.Time.Format("2006-01-02")
		return &formatted
	}
	return nil
}

func nullStringToString(s sql.NullString) string {
	if s.Valid {
		return s.String
	}
	return ""
}

func GetDriverByID(c *fiber.Ctx) error {
	id := c.Params("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID format"})
	}

	var driver struct {
		DriverID         int
		DriverName       string
		DriverAge        *int
		DriverLicense    string
		DriverNumber     string
		DriverAddress    *string
		DriverStatus     *string
		DateOfJoining    *time.Time
		ExperienceYears  *int
		LicenseExpiry    time.Time
		EmergencyContact *string
		AssignedRouteID  *int
		CreatedAt        time.Time
		UpdatedAt        time.Time
		DOB              *time.Time
		Col1             *string
		Col2             *string
	}

	err = db.Pool.QueryRow(context.Background(), "SELECT * FROM get_driver_by_id($1)", idInt).Scan(
		&driver.DriverID, &driver.DriverName, &driver.DriverAge, &driver.DriverLicense, &driver.DriverNumber,
		&driver.DriverAddress, &driver.DriverStatus, &driver.DateOfJoining, &driver.ExperienceYears,
		&driver.LicenseExpiry, &driver.EmergencyContact, &driver.AssignedRouteID, &driver.CreatedAt,
		&driver.UpdatedAt, &driver.DOB, &driver.Col1, &driver.Col2,
	)
	if err != nil {

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(driver)
}
