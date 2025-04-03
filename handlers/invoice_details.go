package handlers

import (
	"context"
	"encoding/json"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/PragaL15/go_newBackend/go_backend/db"
)

// InvoiceRequest represents the request payload
type InvoiceRequest struct {
	OrderID              int     `json:"order_id"`
	WholesellerID        int     `json:"wholeseller_id"`
	TotalAmount         float64 `json:"total_amount"`
	DiscountAmount      float64 `json:"discount_amount,omitempty"`
	TaxAmount           float64 `json:"tax_amount,omitempty"`
	ProposedDeliveryDate string  `json:"proposed_delivery_date,omitempty"`
	Notes               string  `json:"notes,omitempty"`
	DecisionNotes       string  `json:"decision_notes,omitempty"`
	InvoiceNumber       string  `json:"invoice_number,omitempty"`
}

func InsertInvoiceDetails(c *fiber.Ctx) error {
	var req InvoiceRequest
	if err := c.BodyParser(&req); err != nil {
		log.Printf("Invalid request payload: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	var jsonResponse string
	err := db.Pool.QueryRow(context.Background(),
		`SELECT business_schema.create_and_send_invoice($1, $2, $3, $4, $5, $6, $7, $8, $9)`,
		req.OrderID, req.WholesellerID, req.TotalAmount, req.DiscountAmount, req.TaxAmount,
		req.ProposedDeliveryDate, req.Notes, req.DecisionNotes, req.InvoiceNumber,
	).Scan(&jsonResponse)

	if err != nil {
		log.Printf("Database error: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create or update invoice"})
	}

	var response map[string]interface{}
	if err := json.Unmarshal([]byte(jsonResponse), &response); err != nil {
		log.Printf("JSON parse error: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error processing response"})
	}

	return c.JSON(response)
}
