package Masterhandlers

import (
	"context"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/PragaL15/go_newBackend/go_backend/db"
)

type OrderStatus struct {
	OrderStatusID int64  `json:"order_status_id"`
	OrderStatus   string `json:"order_status"`
}

func GetAllOrderStatuses(c *fiber.Ctx) error {
	rows, err := db.Pool.Query(context.Background(), "SELECT * FROM admin_schema.get_all_order_statuses()")
	if err != nil {
		log.Println("Error fetching order statuses:", err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch order statuses"})
	}
	defer rows.Close()

	var orderStatuses []OrderStatus
	for rows.Next() {
		var os OrderStatus
		if err := rows.Scan(&os.OrderStatusID, &os.OrderStatus); err != nil {
			log.Println("Error scanning row:", err)
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to parse data"})
		}
		orderStatuses = append(orderStatuses, os)
	}
	return c.JSON(orderStatuses)
}

func GetOrderStatusByID(c *fiber.Ctx) error {
	id := c.Params("id")
	var os OrderStatus

	err := db.Pool.QueryRow(context.Background(), "SELECT * FROM admin_schema.get_order_status_by_id($1)", id).
		Scan(&os.OrderStatusID, &os.OrderStatus)

	if err != nil {
		log.Println("Error fetching order status:", err)
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Order status not found"})
	}
	return c.JSON(os)
}

func InsertOrderStatus(c *fiber.Ctx) error {
	var os OrderStatus
	if err := c.BodyParser(&os); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	_, err := db.Pool.Exec(context.Background(), "SELECT admin_schema.insert_order_status($1)", os.OrderStatus)
	if err != nil {
		log.Println("Error inserting order status:", err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to insert order status"})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{"message": "Order status created successfully"})
}

// UpdateOrderStatus updates an existing order status
func UpdateOrderStatus(c *fiber.Ctx) error {
	id := c.Params("id")
	var os OrderStatus
	if err := c.BodyParser(&os); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	_, err := db.Pool.Exec(context.Background(), "SELECT admin_schema.update_order_status($1, $2)", id, os.OrderStatus)
	if err != nil {
		log.Println("Error updating order status:", err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update order status"})
	}

	return c.JSON(fiber.Map{"message": "Order status updated successfully"})
}
