package Masterhandlers

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/PragaL15/go_newBackend/go_backend/db"
)

type CashPayment struct {
	ID          int64  `json:"id"`
	PaymentType string `json:"payment_type"`
	IsActive    bool   `json:"is_active"`
}


func GetAllCashPaymentsType(c *fiber.Ctx) error {
	rows, err := db.Pool.Query(context.Background(), "SELECT * FROM admin_schema.get_all_payment_types()")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to fetch payment types",
			"details": err.Error(),
		})
	}
	defer rows.Close()

	var payments []CashPayment
	for rows.Next() {
		var payment CashPayment
		if err := rows.Scan(&payment.ID, &payment.PaymentType, &payment.IsActive); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":   "Error scanning payment type data",
				"details": err.Error(),
			})
		}
		payments = append(payments, payment)
	}
	return c.JSON(payments)
}


func InsertCashPaymentType(c *fiber.Ctx) error {
	type Request struct {
		PaymentType string `json:"payment_type"`
	}

	var req Request
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request format"})
	}

	_, err := db.Pool.Exec(context.Background(), "SELECT admin_schema.insert_payment_type($1)", req.PaymentType)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to insert payment type",
			"details": err.Error(),
		})
	}
	return c.JSON(fiber.Map{"message": "Payment type added successfully"})
}


func UpdateCashPaymentType(c *fiber.Ctx) error {
	type Request struct {
		ID          int64   `json:"id"`
		PaymentType *string `json:"payment_type"` 
		IsActive    *bool   `json:"is_active"`    
	}

	var req Request
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request format"})
	}

	_, err := db.Pool.Exec(context.Background(), "SELECT admin_schema.update_payment_type($1, $2, $3)",
		req.ID, req.PaymentType, req.IsActive)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to update payment type",
			"details": err.Error(),
		})
	}
	return c.JSON(fiber.Map{"message": "Payment type updated successfully"})
}
