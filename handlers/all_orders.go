package handlers

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v4"
	"github.com/PragaL15/go_newBackend/go_backend/db"
)

type Order struct {
	OrderID              int64   `json:"order_id"`
	DateOfOrder          string  `json:"date_of_order"`
	OrderStatus          int     `json:"order_status"`
	ExpectedDeliveryDate string  `json:"expected_delivery_date"`
	ActualDeliveryDate   string  `json:"actual_delivery_date"`
	RetailerID           int     `json:"retailer_id"`
	WholesellerID        int     `json:"wholeseller_id"`
	LocationID           int     `json:"location_id"`
	StateID              int     `json:"state_id"`
	Address              string  `json:"address"`
	Pincode              string  `json:"pincode"`
	TotalOrderAmount     float64 `json:"total_order_amount"`
	DiscountAmount       float64 `json:"discount_amount"`
	TaxAmount            float64 `json:"tax_amount"`
	FinalAmount          float64 `json:"final_amount"`
}

// Insert Order Handler
func InsertOrderHandler(c *fiber.Ctx) error {
	var req Order
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	var orderID int64
	query := `SELECT business_schema.insert_order($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)`

	err := db.Pool.QueryRow(context.Background(), query,
		req.DateOfOrder, req.OrderStatus, req.ExpectedDeliveryDate, req.ActualDeliveryDate,
		req.RetailerID, req.WholesellerID, req.LocationID, req.StateID, req.Address, req.Pincode,
		req.TotalOrderAmount, req.DiscountAmount, req.TaxAmount, req.FinalAmount,
	).Scan(&orderID)

	if err != nil {
		log.Println("Database insert error:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to insert order"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Order created successfully", "order_id": orderID})
}

// Get Order Handler
func GetOrderHandler(c *fiber.Ctx) error {
	orderID := c.Query("order_id")

	if orderID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Missing order_id parameter"})
	}

	query := `SELECT * FROM business_schema.get_order($1)`
	row := db.Pool.QueryRow(context.Background(), query, orderID)

	var order Order
	err := row.Scan(
		&order.OrderID, &order.DateOfOrder, &order.OrderStatus, &order.ExpectedDeliveryDate, &order.ActualDeliveryDate,
		&order.RetailerID, &order.WholesellerID, &order.LocationID, &order.StateID, &order.Address, &order.Pincode,
		&order.TotalOrderAmount, &order.DiscountAmount, &order.TaxAmount, &order.FinalAmount,
	)

	if err == pgx.ErrNoRows {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Order not found"})
	} else if err != nil {
		log.Println("Database fetch error:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal server error"})
	}

	return c.JSON(order)
}

func UpdateOrderHandler(c *fiber.Ctx) error {
	var req Order
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	query := `SELECT business_schema.update_order($1, $2, $3, $4, $5, $6, $7)`

	_, err := db.Pool.Exec(context.Background(), query,
		req.OrderID, req.OrderStatus, req.ExpectedDeliveryDate, req.ActualDeliveryDate,
		req.DiscountAmount, req.TaxAmount, req.FinalAmount,
	)

	if err != nil {
		log.Println("Database update error:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update order"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Order updated successfully"})
}
