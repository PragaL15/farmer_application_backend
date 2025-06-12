package Marketoppurtinities

import (
	"context"
	"farmerapp/go_backend/db"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type TopRetailerProduct struct {
	RetailerID   int64   `json:"retailer_id"`
	RetailerName string  `json:"retailer_name"`
	ProductID    int     `json:"product_id"`
	ProductName  string  `json:"product_name"`
	UnitID       int     `json:"unit_id"`
	Quantity     float64 `json:"quantity"`
	OrderValue   float64 `json:"order_value"`
}

// GetTopRetailersHandler godoc
// @Summary      Get top 5 bulk ordering retailers (product-wise)
// @Description  Returns top 5 bulk orders made by retailers per product
// @Tags         Market Opportunities
// @Accept       json
// @Produce      json
// @Success      200  {array}   TopRetailerProduct
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
	var topRetailers []TopRetailerProduct
	for rows.Next() {
		var retailer TopRetailerProduct
		err := rows.Scan(
			&retailer.RetailerID,
			&retailer.RetailerName,
			&retailer.ProductID,
			&retailer.ProductName,
			&retailer.UnitID,
			&retailer.Quantity,
			&retailer.OrderValue,
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
