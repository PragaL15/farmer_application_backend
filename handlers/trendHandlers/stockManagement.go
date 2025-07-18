package TrendHandlers

import (
	"context"
	"encoding/json"
	"farmerapp/go_backend/db"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type CurrentStockData struct {
	ProductID    int     `json:"product_id"`
	ProductName  string  `json:"product_name"`
	MandiID      int     `json:"mandi_id"`
	MandiName    string  `json:"mandi_name"`
	CurrentStock float64 `json:"current_stock"`
}
type LeastStockedData struct {
	ProductID   int     `json:"product_id"`
	ProductName string  `json:"product_name"`
	MandiID     int     `json:"mandi_id"`
	MandiName   string  `json:"mandi_name"`
	StockLeft   float64 `json:"stock_left"`
}
type StockAvailabilityData struct {
	StockID                     int     `json:"stock_id"`
	ProductID                   int     `json:"product_id"`
	ProductName                 string  `json:"product_name"`
	MandiID                     int     `json:"mandi_id"`
	MandiName                   string  `json:"mandi_name"`
	StockLeft                   float64 `json:"stock_left"`
	MaximumStockLevel           float64 `json:"maximum_stock_level"`
	StockAvailabilityPercentage float64 `json:"stock_availability_percentage"`
}
type LowStockProduct struct {
	ProductID    int         `json:"product_id"`
	ProductName  string      `json:"product_name"`
	CurrentStock float64     `json:"current_stock"`
	Mandis       []MandiData `json:"mandis"`
}
type MandiData struct {
	MandiID    int     `json:"mandi_id"`
	MandiName  string  `json:"mandi_name"`
	MandiStock float64 `json:"mandi_stock"`
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

func GetCurrentStockByProductHandler(c *fiber.Ctx) error {
	productID := c.Params("product_id")
	query := `SELECT * FROM business_schema.get_current_stock_by_product($1);`

	rows, err := db.Pool.Query(context.Background(), query, productID)
	if err != nil {
		log.Println("Error fetching stock by product:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
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
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
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
	var results []LowStockProduct
	for rows.Next() {
		var item LowStockProduct
		var mandisRaw []byte
		err := rows.Scan(&item.ProductID, &item.ProductName, &item.CurrentStock, &mandisRaw)
		if err != nil {
			log.Println("Error scanning row:", err)
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"error":  "Failed to scan row",
				"detail": err.Error(),
			})
		}
		if err := json.Unmarshal(mandisRaw, &item.Mandis); err != nil {
			log.Println("Error unmarshaling mandis JSON:", err)
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"error":  "Failed to parse mandis JSON",
				"detail": err.Error(),
			})
		}
		results = append(results, item)
	}
	return c.JSON(results)
}
