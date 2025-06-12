package handlers

import (
	"context"
	"encoding/json"
	"log"

	"farmerapp/go_backend/db"

	"github.com/gofiber/fiber/v2"
)

type InvoiceRequest struct {
	OrderID              int     `json:"order_id"`
	WholesellerID        int     `json:"wholeseller_id"`
	TotalAmount          float64 `json:"total_amount"`
	DiscountAmount       float64 `json:"discount_amount,omitempty"`
	TaxAmount            float64 `json:"tax_amount,omitempty"`
	ProposedDeliveryDate string  `json:"proposed_delivery_date,omitempty"`
	Notes                string  `json:"notes,omitempty"`
	DecisionNotes        string  `json:"decision_notes,omitempty"`
	InvoiceNumber        string  `json:"invoice_number,omitempty"`
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

type InvoiceDetail struct {
	InvoiceID      int64       `json:"invoice_id"`
	OrderID        int64       `json:"order_id"`
	WholesellerID  int64       `json:"wholeseller_id"`
	TotalAmount    float64     `json:"total_amount"`
	DiscountAmount float64     `json:"discount_amount"`
	TaxAmount      float64     `json:"tax_amount"`
	FinalAmount    float64     `json:"final_amount"`
	InvoiceDate    string      `json:"invoice_date"`
	DueDate        string      `json:"due_date"`
	OrderItems     []OrderItem `json:"order_items"`
}

type OrderItem struct {
	OrderItemID int64   `json:"order_item_id"`
	ProductID   int64   `json:"product_id"`
	ProductName string  `json:"product_name"`
	UnitID      int     `json:"unit_id"`
	UnitName    string  `json:"unit_name"`
	Quantity    float64 `json:"quantity"`
}

func GetInvoiceDetails(c *fiber.Ctx) error {
	invoiceID := c.Params("invoice_id")

	query := `SELECT
			i.id AS invoice_id,
			i.order_id,
			i.wholeseller_id,
			i.total_amount,
			i.discount_amount,
			i.tax_amount,
			i.final_amount,
			i.invoice_date,
			i.due_date,
			oi.order_item_id,
			oi.product_id,
			mp.product_name,
			oi.unit_id,
			mu.unit_name,
			oi.quantity
		FROM business_schema.invoice_table AS i
		JOIN business_schema.invoice_details_table AS idt ON i.id = idt.invoice_id
		JOIN business_schema.order_item_table AS oi ON idt.order_item_id = oi.order_item_id
		JOIN admin_schema.master_product AS mp ON oi.product_id = mp.product_id
		JOIN admin_schema.units_table AS mu ON oi.unit_id = mu.id
		WHERE i.id = $1;`

	rows, err := db.Pool.Query(context.Background(), query, invoiceID)
	if err != nil {
		log.Printf("Error fetching invoice details: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error fetching invoice details"})
	}
	defer rows.Close()

	var invoice InvoiceDetail
	var orderItems []OrderItem

	for rows.Next() {
		var item OrderItem
		err := rows.Scan(
			&invoice.InvoiceID, &invoice.OrderID, &invoice.WholesellerID,
			&invoice.TotalAmount, &invoice.DiscountAmount, &invoice.TaxAmount,
			&invoice.FinalAmount, &invoice.InvoiceDate, &invoice.DueDate,
			&item.OrderItemID, &item.ProductID, &item.ProductName,
			&item.UnitID, &item.UnitName, &item.Quantity,
		)
		if err != nil {
			log.Printf("Error scanning invoice details: %v", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error processing invoice details"})
		}
		orderItems = append(orderItems, item)
	}

	if len(orderItems) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Invoice not found"})
	}

	invoice.OrderItems = orderItems
	return c.JSON(invoice)
}
