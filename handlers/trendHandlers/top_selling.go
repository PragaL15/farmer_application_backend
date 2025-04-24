package TrendHandlers

import (
	"context"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/PragaL15/go_newBackend/go_backend/db"
	"gopkg.in/guregu/null.v4"
)

type TopSalesData struct {
	ProductID           int         `json:"product_id"`
	ProductName         null.String `json:"product_name"`
	MandiID             int         `json:"mandi_id"`
	MandiName           null.String `json:"mandi_name"`
	UnitID              int         `json:"unit_id"`
	Quantity            int         `json:"quantity"`
	Price               null.Float  `json:"price"`
	TotalQuantityKg     null.Float  `json:"total_quantity_kg"`
	ActualDeliveryDate  null.String `json:"actual_delivery_date"`
	TotalPrice          null.Float  `json:"total_price"`
}

func GetTopSellingWeeklyHandler(c *fiber.Ctx) error {
	query := "select * from get_sales_weekly();"
	rows, err := db.Pool.Query(context.Background(), query)
	if err != nil {
		log.Println("Error fetching weekly sales:", err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error":  "Database query failed",
			"detail": err.Error(),
		})
	}
	defer rows.Close()

	var results []TopSalesData
	for rows.Next() {
		var data TopSalesData
		if err := rows.Scan(
			&data.ProductID,
			&data.ProductName,
			&data.MandiID,
			&data.MandiName,
			&data.UnitID,
			&data.Quantity,
			&data.Price,
			&data.TotalQuantityKg,
			&data.ActualDeliveryDate,
			&data.TotalPrice,
		); err != nil {
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

func GetTopSellingMonthlyHandler(c *fiber.Ctx) error {
	query := "SELECT * FROM get_sales_monthly();"
	rows, err := db.Pool.Query(context.Background(), query)
	if err != nil {
		log.Println("Error fetching monthly sales:", err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error":  "Database query failed",
			"detail": err.Error(),
		})
	}
	defer rows.Close()

	var results []TopSalesData
	for rows.Next() {
		var data TopSalesData
		if err := rows.Scan(
			&data.ProductID,
			&data.ProductName,
			&data.MandiID,
			&data.MandiName,
			&data.UnitID,
			&data.Quantity,
			&data.Price,
			&data.TotalQuantityKg,
			&data.ActualDeliveryDate,
			&data.TotalPrice,
		); err != nil {
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

func GetTopSellingYearlyHandler(c *fiber.Ctx) error {
	query := "SELECT * FROM get_sales_yearly();"
	rows, err := db.Pool.Query(context.Background(), query)
	if err != nil {
		log.Println("Error fetching yearly sales:", err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error":  "Database query failed",
			"detail": err.Error(),
		})
	}
	defer rows.Close()

	var results []TopSalesData
	for rows.Next() {
		var data TopSalesData
		if err := rows.Scan(
			&data.ProductID,
			&data.ProductName,
			&data.MandiID,
			&data.MandiName,
			&data.UnitID,
			&data.Quantity,
			&data.Price,
			&data.TotalQuantityKg,
			&data.ActualDeliveryDate,
			&data.TotalPrice,
		); err != nil {
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
