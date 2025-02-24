package handlers

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/PragaL15/go_newBackend/go_backend/db"
	"github.com/go-playground/validator/v10"
)

func InsertListPaymentMethod(c *fiber.Ctx) error {
	type Request struct {
		PaymentType string `json:"payment_type" validate:"required,max=50"`
	}

	var req Request
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	_, err := db.Pool.Exec(context.Background(), `CALL add_cash_payment($1);`, req.PaymentType)

	if err != nil {
		log.Printf("Failed to insert payment method: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to insert payment method"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Payment method added successfully"})
}

func UpdateListPaymentMethod(c *fiber.Ctx) error {
	type Request struct {
		ID          int    `json:"id" validate:"required,min=1"`
		PaymentType string `json:"payment_type" validate:"required,max=50"`
	}

	var req Request
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	_, err := db.Pool.Exec(context.Background(), `CALL update_cash_payment($1, $2);`, req.ID, req.PaymentType)

	if err != nil {
		log.Printf("Failed to update payment method: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update payment method"})
	}

	return c.JSON(fiber.Map{"message": "Payment method updated successfully"})
}

func GetListPaymentMethods(c *fiber.Ctx) error {
	rows, err := db.Pool.Query(context.Background(), "SELECT * FROM get_cash_payment_list()")
	if err != nil {
		log.Printf("Failed to fetch payment methods: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch payment methods"})
	}
	defer rows.Close()

	var payments []map[string]interface{}

	for rows.Next() {
		var id int
		var paymentType *string

		if err := rows.Scan(&id, &paymentType); err != nil {
			log.Printf("Error scanning row: %v", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error processing data"})
		}

		payments = append(payments, map[string]interface{}{
			"id":           id,
			"payment_type": paymentType,
		})
	}

	return c.JSON(payments)
}
