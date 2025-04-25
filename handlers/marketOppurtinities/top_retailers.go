package Marketoppurtinities

import (
	"context"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/PragaL15/go_newBackend/go_backend/db"
)

type OrderSummary struct {
	OrderID       int     `json:"order_id"`
	TotalQuantity float64 `json:"total_quantity"`
	TotalPrice    float64 `json:"total_price"`
	RetailerID    int     `json:"retailer_id"`
	RetailerName  string  `json:"retailer_name"`
}

func GetTopRetailerHandler(c *fiber.Ctx) error {
	query := "SELECT * FROM business_schema.get_order_summary();"
	rows, err := db.Pool.Query(context.Background(), query)
	if err != nil {
		log.Println("Error executing order summary query:", err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error":  "Database query failed",
			"detail": err.Error(),
		})
	}
	defer rows.Close()

	var summaries []OrderSummary
	for rows.Next() {
		var summary OrderSummary
		err := rows.Scan(
			&summary.OrderID,
			&summary.TotalQuantity,
			&summary.TotalPrice,
			&summary.RetailerID,
			&summary.RetailerName,
		)
		if err != nil {
			log.Println("Error scanning order summary row:", err)
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"error":  "Failed to scan row",
				"detail": err.Error(),
			})
		}
		summaries = append(summaries, summary)
	}

	return c.Status(http.StatusOK).JSON(summaries)
}
