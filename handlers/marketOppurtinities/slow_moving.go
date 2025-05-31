package Marketoppurtinities

import (
	"context"
	"log"
	"net/http"

	"farmerapp/go_backend/db"

	"github.com/gofiber/fiber/v2"
)

type SlowMovingProduct struct {
	ProductName string  `json:"product_name"`
	MandiName   string  `json:"mandi_name"`
	StockLeft   float64 `json:"stock_left"`
	WeeklySales float64 `json:"weekly_sales"`
	DaysInStock int     `json:"days_in_stock"`
}

// GetSlowMovingProductsHandler godoc
// @Summary      Get slow moving products
// @Description  Returns products that have been in stock for more than 3 days and are moving slowly
// @Tags         Market Opportunities
// @Accept       json
// @Produce      json
// @Success      200  {array}   SlowMovingProduct
// @Failure      500  {object}  map[string]string
// @Router       /api/slow-moving-products [get]
func GetSlowMovingProductsHandler(c *fiber.Ctx) error {
	query := "SELECT * FROM business_schema.get_slow_moving_products();"
	rows, err := db.Pool.Query(context.Background(), query)
	if err != nil {
		log.Println("Error executing slow moving products query:", err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error":  "Database query failed",
			"detail": err.Error(),
		})
	}
	defer rows.Close()
	var slowProducts []SlowMovingProduct
	for rows.Next() {
		var product SlowMovingProduct
		err := rows.Scan(
			&product.ProductName,
			&product.MandiName,
			&product.StockLeft,
			&product.WeeklySales,
			&product.DaysInStock,
		)
		if err != nil {
			log.Println("Error scanning slow product row:", err)
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"error":  "Failed to scan row",
				"detail": err.Error(),
			})
		}
		slowProducts = append(slowProducts, product)
	}
	return c.Status(http.StatusOK).JSON(slowProducts)
}

// Product is in stock more than 7 days
// Its weekly sales < 10% of current stock
