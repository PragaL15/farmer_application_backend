package handlers

import (
	"context"
	"log"
	"strconv"
	"github.com/gofiber/fiber/v2"
	"github.com/PragaL15/go_newBackend/go_backend/db"
	"github.com/go-playground/validator/v10"
)
func InsertMasterMandi(c *fiber.Ctx) error {
	type Request struct {
		MandiLocation     string `json:"mandi_location" validate:"required,max=255"`
		MandiNumber       string `json:"mandi_number" validate:"required,max=50"`
		MandiIncharge     string `json:"mandi_incharge" validate:"required,max=255"`
		MandiInchargeNum  string `json:"mandi_incharge_num" validate:"required,max=15"`
		MandiPincode      string `json:"mandi_pincode" validate:"required,max=6"`
		MandiAddress      string `json:"mandi_address" validate:"max=1000"`
		Remarks           string `json:"remarks" validate:"max=1000"`
		MandiCity         int    `json:"mandi_city" validate:"required"`
		MandiState        int    `json:"mandi_state" validate:"required"`
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
		CALL insert_master_mandi($1, $2, $3, $4, $5, $6, $7, $8, $9);
	`, req.MandiLocation, req.MandiNumber, req.MandiIncharge, req.MandiInchargeNum,
		req.MandiPincode, req.MandiAddress, req.Remarks, req.MandiCity, req.MandiState)

	if err != nil {
		log.Printf("Failed to insert mandi: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to insert mandi"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Mandi added successfully"})
}

func UpdateMasterMandi(c *fiber.Ctx) error {
	type Request struct {
		ID                int    `json:"id" validate:"required,min=1"`
		MandiLocation     string `json:"mandi_location" validate:"required,max=255"`
		MandiNumber       string `json:"mandi_number" validate:"required,max=50"`
		MandiIncharge     string `json:"mandi_incharge" validate:"required,max=255"`
		MandiInchargeNum  string `json:"mandi_incharge_num" validate:"required,max=15"`
		MandiPincode      string `json:"mandi_pincode" validate:"required,max=6"`
		MandiAddress      string `json:"mandi_address" validate:"max=1000"`
		Remarks           string `json:"remarks" validate:"max=1000"`
		MandiCity         int    `json:"mandi_city" validate:"required"`
		MandiState        int    `json:"mandi_state" validate:"required"`
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
		CALL update_master_mandi($1, $2, $3, $4, $5, $6, $7, $8, $9, $10);
	`, req.ID, req.MandiLocation, req.MandiNumber, req.MandiIncharge, req.MandiInchargeNum,
		req.MandiPincode, req.MandiAddress, req.Remarks, req.MandiCity, req.MandiState)

	if err != nil {
		log.Printf("Failed to update mandi: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update mandi"})
	}

	return c.JSON(fiber.Map{"message": "Mandi updated successfully"})
}
func DeleteMasterMandi(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ID is required"})
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID format"})
	}

	_, err = db.Pool.Exec(context.Background(), `
		CALL delete_master_mandi($1);
	`, idInt)

	if err != nil {
		log.Printf("Failed to delete mandi: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete mandi"})
	}

	return c.JSON(fiber.Map{"message": "Mandi deleted successfully"})
}
func GetMandi(c *fiber.Ctx) error {
	rows, err := db.Pool.Query(context.Background(), "SELECT * FROM get_master_mandis()")
	if err != nil {
		log.Printf("Failed to fetch mandi records: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch mandi records"})
	}
	defer rows.Close()

	var mandis []map[string]interface{}

	for rows.Next() {
		var mandiID, mandiCity, mandiState *int
		var mandiLocation, mandiNumber, mandiIncharge, mandiInchargeNum, mandiPincode, mandiAddress, remarks *string
		

		if err := rows.Scan(
			&mandiID, &mandiLocation, &mandiNumber, &mandiIncharge, &mandiInchargeNum,
			&mandiPincode, &mandiAddress, &remarks, &mandiCity, &mandiState,
		); err != nil {
			log.Printf("Error scanning row: %v", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error processing data"})
		}

		mandi := map[string]interface{}{
			"mandi_id":           mandiID,
			"mandi_location":     mandiLocation,
			"mandi_number":       mandiNumber,
			"mandi_incharge":     mandiIncharge,
			"mandi_incharge_num": mandiInchargeNum,
			"mandi_pincode":      mandiPincode,
			"mandi_address":      mandiAddress,
			"remarks":            remarks,
			"mandi_city":         mandiCity,
			"mandi_state":        mandiState,
		}

		mandis = append(mandis, mandi)
	}

	return c.JSON(mandis)
}