package Masterhandlers

import (
	"context"
	"strconv"

	"github.com/PragaL15/go_newBackend/go_backend/db"
	"github.com/gofiber/fiber/v2"
	"github.com/go-playground/validator/v10"
)

var validate2 = validator.New()

type Mandi struct {
	MandiID          int64  `json:"mandi_id,omitempty"`
	MandiLocation    string `json:"mandi_location" validate:"required"`
	MandiIncharge    string `json:"mandi_incharge" validate:"required"`
	MandiInchargeNum string `json:"mandi_incharge_num" validate:"required"`
	MandiPincode     string `json:"mandi_pincode" validate:"required"`
	MandiAddress     string `json:"mandi_address"`
	MandiState       int64  `json:"mandi_state"`
	MandiName        string `json:"mandi_name"`
	MandiShortnames  string `json:"mandi_shortnames"`
}

func InsertMandiDetails(c *fiber.Ctx) error {
	var req Mandi
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request format"})
	}

	if err := validate.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Missing required fields"})
	}

	_, err := db.Pool.Exec(context.Background(), "SELECT admin_schema.insert_mandi($1, $2, $3, $4, $5, $6, $7, $8)",
		req.MandiLocation, req.MandiIncharge, req.MandiInchargeNum, req.MandiPincode,
		req.MandiAddress, req.MandiState, req.MandiName, req.MandiShortnames)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to insert mandi", "details": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Mandi inserted successfully"})
}

func GetAllMandiDetails(c *fiber.Ctx) error {
	rows, err := db.Pool.Query(context.Background(), "SELECT * FROM admin_schema.get_all_mandi()")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch mandi records", "details": err.Error()})
	}
	defer rows.Close()

	var mandis []Mandi
	for rows.Next() {
		var m Mandi
		if err := rows.Scan(&m.MandiID, &m.MandiLocation, &m.MandiIncharge, &m.MandiInchargeNum, &m.MandiPincode, &m.MandiAddress, &m.MandiState, &m.MandiName, &m.MandiShortnames); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error parsing mandi data", "details": err.Error()})
		}
		mandis = append(mandis, m)
	}

	return c.JSON(mandis)
}

func GetMandiDetailsByID(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID format"})
	}

	var m Mandi
	err = db.Pool.QueryRow(context.Background(), "SELECT * FROM admin_schema.get_mandi_by_id($1)", id).
		Scan(&m.MandiID, &m.MandiLocation, &m.MandiIncharge, &m.MandiInchargeNum, &m.MandiPincode, &m.MandiAddress, &m.MandiState, &m.MandiName, &m.MandiShortnames)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Mandi not found"})
	}

	return c.JSON(m)
}

func UpdateMandiDetails(c *fiber.Ctx) error {
	var req Mandi
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request format"})
	}

	if err := validate.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Missing required fields"})
	}

	_, err := db.Pool.Exec(context.Background(), "SELECT admin_schema.update_mandi($1, $2, $3, $4, $5, $6, $7, $8, $9)",
		req.MandiID, req.MandiLocation, req.MandiIncharge, req.MandiInchargeNum, req.MandiPincode,
		req.MandiAddress, req.MandiState, req.MandiName, req.MandiShortnames)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update mandi", "details": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Mandi updated successfully"})
}
