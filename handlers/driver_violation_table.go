package handlers

import (
	"farmerapp/go_backend/db"

	"github.com/gofiber/fiber/v2"

	"context"
	"time"
)

type DriverViolation struct {
	ID            int    `json:"id"`
	DriverID      int    `json:"driver_id"`
	DriverName    string `json:"driver_name"`
	ViolationID   int    `json:"violation_id"`
	ViolationName string `json:"violation_name"`
	EntryDate     string `json:"entry_date"`
}

func GetDriverViolations(c *fiber.Ctx) error {
	rows, err := db.Pool.Query(context.Background(), "SELECT * FROM get_driver_violations()")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	defer rows.Close()

	var violations []map[string]interface{}

	for rows.Next() {
		var id, driverID, violationID int
		var entryDate time.Time
		var driverName, violationName string

		err := rows.Scan(&id, &driverID, &driverName, &violationID, &violationName, &entryDate)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		violations = append(violations, map[string]interface{}{
			"id":             id,
			"driver_id":      driverID,
			"driver_name":    driverName,
			"violation_id":   violationID,
			"violation_name": violationName,
			"entry_date":     entryDate,
		})
	}

	return c.JSON(violations)
}
