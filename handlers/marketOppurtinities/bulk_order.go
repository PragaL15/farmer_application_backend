package Marketoppurtinities

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/PragaL15/go_newBackend/go_backend/db"
)

type OrderItem struct {
	ProductID      int     `json:"product_id"`
	ProductName    string  `json:"product_name"`
	Quantity       float64 `json:"quantity"`
	PriceOfProduct float64 `json:"price_of_product"`
}

type OrderDetail struct {
	OrderID          int         `json:"order_id"`
	DateOfOrder      time.Time   `json:"date_of_order"`
	TotalOrderAmount float64     `json:"total_order_amount"`
	RetailerName     string      `json:"retailer_name"`
	WholesellerName  string      `json:"wholeseller_name"`
	Items            []OrderItem `json:"items"`
}

// GetAllBulkOrderDetailsHandler godoc
// @Summary      Get all bulk order details
// @Description  Retrieves all bulk order details with nested item information per order
// @Tags         Orders
// @Produce      json
// @Success      200 {array} Marketoppurtinities.OrderDetail
// @Failure      500 {object} map[string]string
// @Router       /api/orders/bulk [get]
func GetAllBulkOrderDetailsHandler(c *fiber.Ctx) error {
	orderQuery := `
		SELECT order_id, date_of_order, total_order_amount, retailer_name, wholeseller_name 
		FROM business_schema.get_all_order_details();
	`

	rows, err := db.Pool.Query(context.Background(), orderQuery)
	if err != nil {
		log.Println("Error executing order detail query:", err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error":  "Database query failed",
			"detail": err.Error(),
		})
	}
	defer rows.Close()

	orders := make(map[int]*OrderDetail)

	for rows.Next() {
		var od OrderDetail
		if err := rows.Scan(
			&od.OrderID,
			&od.DateOfOrder,
			&od.TotalOrderAmount,
			&od.RetailerName,
			&od.WholesellerName,
		); err != nil {
			log.Println("Error scanning order row:", err)
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"error":  "Failed to scan order row",
				"detail": err.Error(),
			})
		}
		od.Items = []OrderItem{}
		orders[od.OrderID] = &od
	}

	itemQuery := `
		SELECT order_id, product_id, product_name, quantity, price_of_product 
		FROM business_schema.get_bulk_order_items_details();
	`

	itemRows, err := db.Pool.Query(context.Background(), itemQuery)
	if err != nil {
		log.Println("Error executing item detail query:", err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error":  "Item query failed",
			"detail": err.Error(),
		})
	}
	defer itemRows.Close()

	for itemRows.Next() {
		var item OrderItem
		var orderID int

		if err := itemRows.Scan(
			&orderID,
			&item.ProductID,
			&item.ProductName,
			&item.Quantity,
			&item.PriceOfProduct,
		); err != nil {
			log.Println("Error scanning item row:", err)
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"error":  "Failed to scan item row",
				"detail": err.Error(),
			})
		}

		if order, exists := orders[orderID]; exists {
			order.Items = append(order.Items, item)
		}
	}

	var result []OrderDetail
	for _, order := range orders {
		result = append(result, *order)
	}

	return c.Status(http.StatusOK).JSON(result)
}
