package Masterhandlers

import (
	"context"
	"strconv"

	"github.com/PragaL15/go_newBackend/go_backend/db"
	"github.com/gofiber/fiber/v2"
	"github.com/go-playground/validator/v10"
)

var validate1 = validator.New()

// Struct for Mandi
type Mandi struct {
	MandiID          int    `json:"mandi_id"`
	MandiLocation    string `json:"mandi_location"`
	MandiIncharge    string `json:"mandi_incharge"`
	MandiInchargeNum string `json:"mandi_incharge_num"`
	MandiPincode     string `json:"mandi_pincode"`
	MandiAddress     string `json:"mandi_address"`
	MandiStateID     int    `json:"mandi_state_id"`
	StateName        string `json:"state_name"`
	StateShortnames  string `json:"state_shortnames"`
	MandiName        string `json:"mandi_name"`
	MandiShortnames  string `json:"mandi_shortnames"`
	MandiCityID      int    `json:"mandi_city_id"`
	CityName         string `json:"city_name"`
	CityShortnames   string `json:"city_shortnames"`
}

// Insert Mandi Details
func InsertMandiDetails(c *fiber.Ctx) error {
	var req Mandi
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request format"})
	}

	if err := validate1.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	_, err := db.Pool.Exec(context.Background(), "SELECT admin_schema.insert_mandi($1, $2, $3, $4, $5, $6, $7, $8, $9)",
req.MandiLocation,        
req.MandiIncharge,       
req.MandiInchargeNum,     
req.MandiPincode,        
req.MandiAddress,         
req.MandiStateID,        
req.MandiName,           
req.MandiShortnames,      
req.MandiCityID           )
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
		err := rows.Scan(&m.MandiID, &m.MandiLocation, &m.MandiIncharge, &m.MandiInchargeNum,
			&m.MandiPincode, &m.MandiAddress, &m.MandiStateID, &m.StateName, &m.StateShortnames,
			&m.MandiName, &m.MandiShortnames, &m.MandiCityID, &m.CityName, &m.CityShortnames)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error parsing mandi data", "details": err.Error()})
		}
		mandis = append(mandis, m)
	}

	return c.JSON(mandis)
}

// Get Mandi Details by ID
func GetMandiDetailsByID(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID format"})
	}

	var m Mandi
	err = db.Pool.QueryRow(context.Background(), "SELECT * FROM admin_schema.get_mandi_by_id($1)", id).
		Scan(&m.MandiID, &m.MandiLocation, &m.MandiIncharge, &m.MandiInchargeNum,
			&m.MandiPincode, &m.MandiAddress, &m.MandiStateID, &m.StateName, &m.StateShortnames,
			&m.MandiName, &m.MandiShortnames, &m.MandiCityID, &m.CityName, &m.CityShortnames)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Mandi not found", "details": err.Error()})
	}

	return c.JSON(m)
}

// Update Mandi Details
func UpdateMandiDetails(c *fiber.Ctx) error {
	var req Mandi
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request format"})
	}

	if err := validate1.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	var result string
	err := db.Pool.QueryRow(context.Background(),
		"SELECT admin_schema.update_mandi($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)",
		req.MandiID, req.MandiLocation, req.MandiIncharge, req.MandiInchargeNum, req.MandiPincode,
		req.MandiAddress, req.MandiStateID, req.MandiName, req.MandiShortnames, req.MandiCityID).Scan(&result)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update mandi", "details": err.Error()})
	}

	return c.JSON(fiber.Map{"message": result})
}
