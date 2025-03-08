package Masterhandlers

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/PragaL15/go_newBackend/go_backend/db"
)

type OrderStatus struct {
	OrderID     int       `json:"order_id"`
	OrderStatus string    `json:"order_status"`
}

func GetOrderStatuses(c *fiber.Ctx) error {
	rows, err := db.Pool.Query(context.Background(), "SELECT * FROM sp_get_order_status()")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	defer rows.Close()
	var statuses []map[string]interface{}
	for rows.Next() {
		var orderID int
		var orderStatus string
		err := rows.Scan(&orderID, &orderStatus)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		statuses = append(statuses, map[string]interface{}{
			"order_id":     orderID,
			"order_status": orderStatus,		
		})
	}
	return c.JSON(statuses)
}

func InsertOrderStatus(c *fiber.Ctx) error {
	type Request struct {
		OrderStatus string `json:"order_status"`
	}
	var req Request
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}
	_, err := db.Pool.Exec(context.Background(), "SELECT insert_order_status($1)", req.OrderStatus)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "Order status inserted successfully"})
}

func UpdateOrderStatus(c *fiber.Ctx) error {
	type Request struct {
		OrderID     int    `json:"order_id"`
		OrderStatus string `json:"order_status"`
	}
	var req Request
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}
	_, err := db.Pool.Exec(context.Background(), "SELECT update_order_status($1, $2)", req.OrderID, req.OrderStatus)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "Order status updated successfully"})
}
