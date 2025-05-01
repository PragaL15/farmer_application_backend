package Marketoppurtinities

import (
	"context"
	"log"
	"net/http"
	"github.com/gofiber/fiber/v2"
	"github.com/PragaL15/go_newBackend/go_backend/db"
)
// TopRetailerSummary represents a retailer and their aggregated order data.
type TopRetailerSummary struct {
	RetailerID      int64   `json:"retailer_id" example:"101"`
	RetailerName    string  `json:"retailer_name" example:"Sri Lakshmi Stores"`
	TotalQuantity   float64 `json:"total_quantity" example:"350.5"`
	TotalOrderValue float64 `json:"total_order_value" example:"78000.75"`
}
// GetTopRetailersHandler godoc
//
// @Summary      Get top 5 bulk ordering retailers
// @Description  Returns the top 5 retailers who place the largest quantity of bulk orders
// @Tags         Market Opportunities
// @Accept       json
// @Produce      json
// @Success      200  {array}   TopRetailerSummary
// @Failure      500  {object}  map[string]string
// @Router       /api/top-retailers [get]
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
