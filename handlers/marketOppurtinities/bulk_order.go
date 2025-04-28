package Marketoppurtinities

import (
	"context"
	"log"
	"net/http"
	"time"
	"github.com/gofiber/fiber/v2"
	"github.com/PragaL15/go_newBackend/go_backend/db"
)

type OrderDetail struct {
	OrderID          int       `json:"order_id"`
	DateOfOrder      time.Time `json:"date_of_order"`
	TotalOrderAmount float64   `json:"total_order_amount"`
	FinalAmount      float64   `json:"final_amount"`
	ProductName      string    `json:"product_name"`
	WholesellerName  string    `json:"wholeseller_name"`
	RetailerName     string    `json:"retailer_name"`
}

func GetAllBulkOrderDetailsHandler(c *fiber.Ctx) error {
	query := "SELECT * FROM business_schema.get_all_order_details();"
	rows, err := db.Pool.Query(context.Background(), query)
	if err != nil {
		log.Println("Error executing order detail query:", err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error":  "Database query failed",
			"detail": err.Error(),
		})
	}
	defer rows.Close()
	var orderDetails []OrderDetail
	for rows.Next() {
		var od OrderDetail
		if err := rows.Scan(
			&od.OrderID,
			&od.DateOfOrder,
			&od.TotalOrderAmount,
			&od.FinalAmount,
			&od.ProductName,
			&od.WholesellerName,
			&od.RetailerName,
		); err != nil {
			log.Println("Error scanning order row:", err)
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"error":  "Failed to scan row",
				"detail": err.Error(),
			})
		}
		orderDetails = append(orderDetails, od)
	}
	return c.Status(http.StatusOK).JSON(orderDetails)
}
