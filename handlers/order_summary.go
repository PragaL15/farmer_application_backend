package handlers

import (
	"context"
	"farmerapp/go_backend/db"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type WholesellerStockDetail struct {
	WholesellerID int     `json:"wholeseller_id"`
	MandiID       int     `json:"mandi_id"`
	ProductID     int     `json:"product_id"`
	ProductName   string  `json:"product_name"`
	StockLeft     float64 `json:"stock_left"`
	StockIn       float64 `json:"stock_in"`
}

func GetAllWholesellerStockDetailsHandler(c *fiber.Ctx) error {
	userId := c.Params("user_Id")
	query := "SELECT * FROM business_schema.get_wholeseller_stock_details(" + userId + ")"

	rows, err := db.Pool.Query(context.Background(), query)
	if err != nil {
		log.Println("Error executing wholeseller stock detail query:", err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error":  "Database query failed",
			"detail": err.Error(),
		})
	}
	defer rows.Close()

	stockDetails := []WholesellerStockDetail{}
	for rows.Next() {
		var sd WholesellerStockDetail
		if err := rows.Scan(
			&sd.WholesellerID,
			&sd.MandiID,
			&sd.ProductID,
			&sd.ProductName,
			&sd.StockLeft,
			&sd.StockIn,
		); err != nil {
			log.Println("Error scanning wholeseller stock row:", err)
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"error":  "Failed to scan row",
				"detail": err.Error(),
			})
		}
		stockDetails = append(stockDetails, sd)
	}

	return c.Status(http.StatusOK).JSON(stockDetails)
}
