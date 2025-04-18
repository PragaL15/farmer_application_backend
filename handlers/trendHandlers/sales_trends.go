package TrendHandlers

import (
	"context"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/PragaL15/go_newBackend/go_backend/db"
)

type SalesData struct {
	MonthYear    string  `json:"month_year"`
	TotalOrders  int64   `json:"total_orders"`
	TotalRevenue float64 `json:"total_revenue"`
}

func GetSalesMonthlyHandler(c *fiber.Ctx) error {
	query := "SELECT * FROM get_sales_by_duration('monthly');"
	rows, err := db.Pool.Query(context.Background(), query)
	if err != nil {
		log.Println("Error fetching monthly sales:", err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error":  "Database query failed",
			"detail": err.Error(),
		})
	}
	defer rows.Close()
	var results []SalesData
	for rows.Next() {
		var data SalesData
		if err := rows.Scan(&data.MonthYear, &data.TotalOrders, &data.TotalRevenue); err != nil {
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
func GetSalesWeeklyHandler(c *fiber.Ctx) error {
	query := "select * from get_weekday_sales();"
	rows, err := db.Pool.Query(context.Background(), query)
	if err != nil {
		log.Println("Error fetching weekly sales:", err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error":  "Database query failed",
			"detail": err.Error(),
		})
	}
	defer rows.Close()

	var results []SalesData
	for rows.Next() {
		var data SalesData
		if err := rows.Scan(&data.MonthYear, &data.TotalOrders, &data.TotalRevenue); err != nil {
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

func GetSalesYearlyHandler(c *fiber.Ctx) error {
	query := "SELECT * FROM get_sales_by_duration('yearly');"
	rows, err := db.Pool.Query(context.Background(), query)
	if err != nil {
		log.Println("Error fetching yearly sales:", err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error":  "Database query failed",
			"detail": err.Error(),
		})
	}
	defer rows.Close()

	var results []SalesData
	for rows.Next() {
		var data SalesData
		if err := rows.Scan(&data.MonthYear, &data.TotalOrders, &data.TotalRevenue); err != nil {
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
