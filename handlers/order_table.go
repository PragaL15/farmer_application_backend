package handlers

import (
	"context"
	"database/sql"


	"github.com/gofiber/fiber/v2"
	"github.com/PragaL15/go_newBackend/go_backend/db"
)

type Order struct {
	OrderID              int     `json:"order_id"`
	DateOfOrder          string  `json:"date_of_order"`
	OrderStatus          int     `json:"order_status"`
	OrderStatusName      string  `json:"order_status_name"`
	ExpectedDeliveryDate string  `json:"expected_delivery_date"`
	ActualDeliveryDate   string  `json:"actual_delivery_date"`
	RetailerID           int     `json:"retailer_id"`
	RetailerName         string  `json:"retailer_name"`
	WholesellerID        int     `json:"wholeseller_id"`
	WholesellerName      string  `json:"wholeseller_name"`
	LocationID           int     `json:"location_id"`
	LocationName         string  `json:"location_name"`
	StateID              int     `json:"state_id"`
	StateName            string  `json:"state_name"`
	Pincode              string  `json:"pincode"`
	Address              string  `json:"address"`
}

// Format time safely, return an empty string for NULL values
func formatTimestamp(t sql.NullTime) string {
	if t.Valid {
		return t.Time.Format("2006-01-02 15:04:05") // Format: YYYY-MM-DD HH:MM:SS
	}
	return "" // Return empty string if NULL
}

// Get all orders with formatted timestamps and NULL-safe handling

func GetOrders(c *fiber.Ctx) error {
	rows, err := db.Pool.Query(context.Background(), "SELECT * FROM sp_get_orders()")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	defer rows.Close()

	var orders []Order
	for rows.Next() {
		var order Order
		var dateOfOrder, expectedDeliveryDate, actualDeliveryDate sql.NullTime

		err := rows.Scan(
			&order.OrderID, &dateOfOrder, &order.OrderStatus, &order.OrderStatusName,
			&expectedDeliveryDate, &actualDeliveryDate, &order.RetailerID, &order.RetailerName,
			&order.WholesellerID, &order.WholesellerName, &order.LocationID, &order.LocationName,
			&order.StateID, &order.StateName, &order.Pincode, &order.Address,
		)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		// Format timestamps safely
		order.DateOfOrder = formatTimestamp(dateOfOrder)
		order.ExpectedDeliveryDate = formatTimestamp(expectedDeliveryDate)
		order.ActualDeliveryDate = formatTimestamp(actualDeliveryDate)

		orders = append(orders, order)
	}

	return c.JSON(orders)
}



// Insert a new order
func InsertOrder(c *fiber.Ctx) error {
	var req Order
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	_, err := db.Pool.Exec(context.Background(),
		"SELECT sp_insert_order($1, $2, $3, $4, $5, $6, $7, $8, $9)",
		req.OrderStatus, req.ExpectedDeliveryDate, req.ActualDeliveryDate,
		req.RetailerID, req.WholesellerID, req.LocationID,
		req.StateID, req.Pincode, req.Address,
	)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "Order inserted successfully"})
}

// Update an existing order
func UpdateOrder(c *fiber.Ctx) error {
	var req struct {
		OrderID            int    `json:"order_id"`
		OrderStatus        int    `json:"order_status"`
		ActualDeliveryDate string `json:"actual_delivery_date"`
	}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	_, err := db.Pool.Exec(context.Background(),
		"SELECT sp_update_order($1, $2, $3)",
		req.OrderID, req.OrderStatus, req.ActualDeliveryDate,
	)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "Order updated successfully"})
}
