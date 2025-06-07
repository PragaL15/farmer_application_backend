package Masterhandlers

import (
	"context"
	"log"
	"strconv"

	"farmerapp/go_backend/db"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v4"
)

func InsertPaymentMode(c *fiber.Ctx) error {
	type Request struct {
		PaymentMode string `json:"payment_mode"`
	}
	var req Request
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}
	_, err := db.Pool.Exec(context.Background(),
		`INSERT INTO admin_schema.mode_of_payments_list (payment_mode, is_active) VALUES ($1, 1)`,
		req.PaymentMode)
	if err != nil {
		log.Println("Error inserting:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to insert"})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Payment mode added"})
}

func UpdatePaymentMode(c *fiber.Ctx) error {
	id := c.Params("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID format"})
	}
	type Request struct {
		PaymentMode string `json:"payment_mode"`
	}
	var req Request
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}
	_, err = db.Pool.Exec(context.Background(),
		`UPDATE admin_schema.mode_of_payments_list SET payment_mode = $1 WHERE id = $2 AND is_active = 1`,
		req.PaymentMode, idInt)
	if err != nil {
		log.Println("Error updating:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update"})
	}
	return c.JSON(fiber.Map{"message": "Payment mode updated"})
}

func DeletePaymentMode(c *fiber.Ctx) error {
	id := c.Params("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID format"})
	}
	_, err = db.Pool.Exec(context.Background(),
		`UPDATE admin_schema.mode_of_payments_list SET is_active = 0 WHERE id = $1`,
		idInt)
	if err != nil {
		log.Println("Error soft deleting:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete"})
	}
	return c.JSON(fiber.Map{"message": "Payment mode soft deleted"})
}

func GetPaymentModes(c *fiber.Ctx) error {
	rows, err := db.Pool.Query(context.Background(),
		`SELECT id, payment_mode FROM admin_schema.mode_of_payments_list WHERE is_active = 1`)
	if err != nil {
		log.Println("Error fetching:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch data"})
	}
	defer rows.Close()

	var paymentModes []map[string]interface{}
	for rows.Next() {
		var id int
		var paymentMode string
		if err := rows.Scan(&id, &paymentMode); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error scanning data"})
		}
		paymentModes = append(paymentModes, map[string]interface{}{
			"id":           id,
			"payment_mode": paymentMode,
		})
	}
	return c.JSON(paymentModes)
}

func GetPaymentModeByID(c *fiber.Ctx) error {
	id := c.Params("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID format"})
	}
	var paymentMode string
	err = db.Pool.QueryRow(context.Background(),
		`SELECT payment_mode FROM admin_schema.mode_of_payments_list WHERE id = $1 AND is_active = 1`,
		idInt).Scan(&paymentMode)
	if err != nil {
		if err == pgx.ErrNoRows {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Payment mode not found"})
		}
		log.Println("Error fetching by ID:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch data"})
	}
	return c.JSON(fiber.Map{"id": idInt, "payment_mode": paymentMode})
}
