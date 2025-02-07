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
