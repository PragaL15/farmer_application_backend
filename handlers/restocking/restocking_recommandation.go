// Package RestockingStock handles restocking-related APIs
package RestockingStock

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"farmerapp/go_backend/db"

	"github.com/gofiber/fiber/v2"
)

// Mandi represents mandi-wise stock details
type Mandi struct {
	MandiID    int     `json:"mandi_id"`
	MandiName  string  `json:"mandi_name"`
	MandiStock float64 `json:"mandi_stock"`
}

// RestockingProduct represents a product that needs restocking
type RestockingProduct struct {
	ProductID         int     `json:"product_id"`
	ProductName       string  `json:"product_name"`
	StockToSalesRatio float64 `json:"stock_to_sales_ratio"`
	StockStatus       string  `json:"stock_status"`
	Mandi             []Mandi `json:"mandi"`
}

// GetRestockingProductsHandler godoc
// @Summary      Get restocking product recommendations
// @Description  Returns products with stock-to-sales ratio < 4, sorted by urgency, with mandi-wise breakdown
// @Tags         Restocking
// @Produce      json
// @Success      200 {array} RestockingProduct
// @Failure      500 {object} map[string]interface{}
// @Router       /restocking/recommendations [get]
func GetRestockingProductsHandler(c *fiber.Ctx) error {
	query := "SELECT * FROM business_schema.get_re_stock_products();"
	rows, err := db.Pool.Query(context.Background(), query)
	if err != nil {
		log.Println("Error executing restocking query:", err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error":  "Database query failed",
			"detail": err.Error(),
		})
	}
	defer rows.Close()

	var restockingProducts []RestockingProduct

	for rows.Next() {
		var product RestockingProduct
		var mandiJSON []byte

		err := rows.Scan(
			&product.ProductID,
			&product.ProductName,
			&product.StockToSalesRatio,
			&product.StockStatus,
			&mandiJSON,
		)
		if err != nil {
			log.Println("Error scanning row:", err)
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"error":  "Failed to scan row",
				"detail": err.Error(),
			})
		}

		// Unmarshal JSONB into Go struct
		if err := json.Unmarshal(mandiJSON, &product.Mandi); err != nil {
			log.Println("Error unmarshaling mandi JSON:", err)
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"error":  "Failed to parse mandi data",
				"detail": err.Error(),
			})
		}

		restockingProducts = append(restockingProducts, product)
	}

	return c.Status(http.StatusOK).JSON(restockingProducts)
}
