package TrendHandlers

import (
	"context"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/PragaL15/go_newBackend/go_backend/db"
)
type CurrentStockData struct {
	ProductID     int     `json:"product_id"`
	ProductName   string  `json:"product_name"`
	MandiID       int     `json:"mandi_id"`
	MandiName     string  `json:"mandi_name"`
	CurrentStock  float64 `json:"current_stock"`
}

type LeastStockedData struct {
	ProductID     int     `json:"product_id"`
	ProductName   string  `json:"product_name"`
	MandiID       int     `json:"mandi_id"`
	MandiName     string  `json:"mandi_name"`
	StockLeft     float64 `json:"stock_left"`
}

type StockAvailabilityData struct {
	StockID                    int     `json:"stock_id"`
	ProductID                  int     `json:"product_id"`
	ProductName                string  `json:"product_name"`
	MandiID                    int     `json:"mandi_id"`
	MandiName                  string  `json:"mandi_name"`
	StockLeft                  float64 `json:"stock_left"`
	MaximumStockLevel         float64 `json:"maximum_stock_level"`
	StockAvailabilityPercentage float64 `json:"stock_availability_percentage"`
}

type LowStockItemData struct {
	ProductID       int     `json:"product_id"`
	ProductName     string  `json:"product_name"`
	MandiID         int     `json:"mandi_id"`
	MandiName       string  `json:"mandi_name"`
	StockLeft       float64 `json:"stock_left"`
	MinimumStockLevel float64 `json:"minimum_stock_level"`
}
func GetCurrentStockByMandiHandler(c *fiber.Ctx) error {
	mandiID := c.Params("mandi_id")
	query := `SELECT * FROM business_schema.get_current_stock_by_mandi($1);`
	rows, err := db.Pool.Query(context.Background(), query, mandiID)
	if err != nil {
		log.Println("Error fetching current stock data:", err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error":  "Database query failed",
			"detail": err.Error(),
		})
	}
	defer rows.Close()

	var results []CurrentStockData
	for rows.Next() {
		var data CurrentStockData
		if err := rows.Scan(&data.ProductID, &data.ProductName, &data.MandiID, &data.MandiName, &data.CurrentStock); err != nil {
			log.Println("Error scanning row:", err)
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"error":  "Failed to scan row",
				"detail": err.Error(),
			})
		}
		results = append(results, data)
	}

	return c.JSON(results)
}
func GetLeastStockedProductsHandler(c *fiber.Ctx) error {
	query := `SELECT * FROM business_schema.get_least_stocked_products();`
	rows, err := db.Pool.Query(context.Background(), query)
	if err != nil {
		log.Println("Error fetching least stocked products:", err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error":  "Database query failed",
			"detail": err.Error(),
		})
	}
	defer rows.Close()

	var results []LeastStockedData
	for rows.Next() {
		var data LeastStockedData
		if err := rows.Scan(&data.ProductID, &data.ProductName, &data.MandiID, &data.MandiName, &data.StockLeft); err != nil {
			log.Println("Error scanning row:", err)
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"error":  "Failed to scan row",
				"detail": err.Error(),
			})
		}
		results = append(results, data)
	}

	return c.JSON(results)
}
func GetStockAvailabilityPercentageHandler(c *fiber.Ctx) error {
	query := `SELECT * FROM business_schema.get_stock_availability_percentage();`
	rows, err := db.Pool.Query(context.Background(), query)
	if err != nil {
		log.Println("Error fetching stock availability percentage:", err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error":  "Database query failed",
			"detail": err.Error(),
		})
	}
	defer rows.Close()

	var results []StockAvailabilityData
	for rows.Next() {
		var data StockAvailabilityData
		if err := rows.Scan(&data.StockID, &data.ProductID, &data.ProductName, &data.MandiID, &data.MandiName, &data.StockLeft, &data.MaximumStockLevel, &data.StockAvailabilityPercentage); err != nil {
			log.Println("Error scanning row:", err)
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"error":  "Failed to scan row",
				"detail": err.Error(),
			})
		}
		results = append(results, data)
	}

	return c.JSON(results)
}
func GetLowStockItemsHandler(c *fiber.Ctx) error {
	query := `SELECT * FROM business_schema.get_low_stock_items();`
	rows, err := db.Pool.Query(context.Background(), query)
	if err != nil {
		log.Println("Error fetching low stock items:", err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error":  "Database query failed",
			"detail": err.Error(),
		})
	}
	defer rows.Close()

	var results []LowStockItemData
	for rows.Next() {
		var data LowStockItemData
		if err := rows.Scan(&data.ProductID, &data.ProductName, &data.MandiID, &data.MandiName, &data.StockLeft, &data.MinimumStockLevel); err != nil {
			log.Println("Error scanning row:", err)
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"error":  "Failed to scan row",
				"detail": err.Error(),
			})
		}
		results = append(results, data)
	}

	return c.JSON(results)
}
