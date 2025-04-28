package handlers

import (
	"context"
	"log"
	"net/http"
	"github.com/gofiber/fiber/v2"
	"github.com/PragaL15/go_newBackend/go_backend/db"
)

type OrderItems struct {
	OrderItemID   int     `json:"order_item_id"`
	ProductID     int     `json:"product_id"`
	ProductName   string  `json:"product_name"`
	Quantity      float64 `json:"quantity"`
	UnitID        int     `json:"unit_id"`
	UnitName      string  `json:"unit_name"`
	MaxItemPrice  float64 `json:"max_item_price"`
}

type OrderDetail struct {
	OrderID          int          `json:"order_id"`
	TotalOrderAmount float64      `json:"total_order_amount"`
	OrderItems       []OrderItems `json:"order_items"`
}

func GetAllOrderItemDetailsHandler(c *fiber.Ctx) error {
	query := `
		SELECT 
			order_id,
			total_order_amount,
			order_item_id,
			product_id,
			product_name,
			quantity,
			unit_id,
			max_item_price,
			unit_name
		FROM 
			business_schema.get_order_details();
	`

	rows, err := db.Pool.Query(context.Background(), query)
	if err != nil {
		log.Println("Error executing order details query:", err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error":  "Database query failed",
			"detail": err.Error(),
		})
	}
	defer rows.Close()

	orderDetailsMap := make(map[int]*OrderDetail)
	for rows.Next() {
		var od OrderDetail
		var oi OrderItems
		if err := rows.Scan(
			&od.OrderID,
			&od.TotalOrderAmount,
			&oi.OrderItemID,
			&oi.ProductID,
			&oi.ProductName,
			&oi.Quantity,
			&oi.UnitID,
			&oi.MaxItemPrice,
			&oi.UnitName,
		); err != nil {
			log.Println("Error scanning order row:", err)
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"error":  "Failed to scan row",
				"detail": err.Error(),
			})
		}

		if existingOrder, exists := orderDetailsMap[od.OrderID]; exists {
			existingOrder.OrderItems = append(existingOrder.OrderItems, oi)
		} else {
			od.OrderItems = append(od.OrderItems, oi)
			orderDetailsMap[od.OrderID] = &od
		}
	}

	var orderDetails []OrderDetail
	for _, order := range orderDetailsMap {
		orderDetails = append(orderDetails, *order)
	}

	return c.Status(http.StatusOK).JSON(orderDetails)
}
