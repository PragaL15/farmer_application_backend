package Marketoppurtinities

import (
	"context"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/PragaL15/go_newBackend/go_backend/db"
)

type TopRetailerSummary struct {
	RetailerID     int64   `json:"retailer_id"`
	RetailerName   string  `json:"retailer_name"`
	TotalQuantity  float64 `json:"total_quantity"`
	TotalOrderValue float64 `json:"total_order_value"`
}

func GetTopRetailersHandler(c *fiber.Ctx) error {
	query := "SELECT * FROM business_schema.get_top_5_bulk_ordering_retailers();"
	rows, err := db.Pool.Query(context.Background(), query)
	if err != nil {
		log.Println("Error executing top 5 retailers query:", err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error":  "Database query failed",
			"detail": err.Error(),
		})
	}
	defer rows.Close()

	var topRetailers []TopRetailerSummary
	for rows.Next() {
		var retailer TopRetailerSummary
		err := rows.Scan(
			&retailer.RetailerID,
			&retailer.RetailerName,
			&retailer.TotalQuantity,
			&retailer.TotalOrderValue,
		)
		if err != nil {
			log.Println("Error scanning retailer row:", err)
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"error":  "Failed to scan row",
				"detail": err.Error(),
			})
		}
		topRetailers = append(topRetailers, retailer)
	}

	return c.Status(http.StatusOK).JSON(topRetailers)
}
