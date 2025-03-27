package Masterhandlers

import (
	"context"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/PragaL15/go_newBackend/go_backend/db"
)

type CashPayment struct {
	ID          int64  `json:"id"`
	PaymentType string `json:"payment_type"`
}

func GetAllCashPaymentsType(c *fiber.Ctx) error {
	rows, err := db.Pool.Query(context.Background(), "SELECT * FROM admin_schema.get_all_cash_payments()")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch cash payments", "details": err.Error()})
	}
	defer rows.Close()

	var payments []CashPayment
	for rows.Next() {
		var payment CashPayment
		if err := rows.Scan(&payment.ID, &payment.PaymentType); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error scanning cash payment data", "details": err.Error()})
		}
		payments = append(payments, payment)
	}
	return c.JSON(payments)
}

// Get cash payment by ID
func GetCashPaymentByID(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID format"})
	}

	var payment CashPayment
	err = db.Pool.QueryRow(context.Background(), "SELECT * FROM admin_schema.get_cash_payment_by_id($1)", id).
		Scan(&payment.ID, &payment.PaymentType)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Cash payment not found"})
	}
	return c.JSON(payment)
}

func InsertPaymentType(c *fiber.Ctx) error {
	type Request struct {
		PaymentType string `json:"payment_type"`
	}

	var req Request
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request format"})
	}

	_, err := db.Pool.Exec(context.Background(), "SELECT admin_schema.insert_cash_payment($1)", req.PaymentType)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to insert cash payment", "details": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "Cash payment added successfully"})
}

// Update a cash payment
func UpdatePaymentType(c *fiber.Ctx) error {
	type Request struct {
		ID          int64  `json:"id"`
		PaymentType string `json:"payment_type"`
	}

	var req Request
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request format"})
	}

	_, err := db.Pool.Exec(context.Background(), "SELECT admin_schema.update_cash_payment($1, $2)", req.ID, req.PaymentType)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update cash payment", "details": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "Cash payment updated successfully"})
}
