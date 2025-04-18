package TrendHandlers

import (
	"context"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/PragaL15/go_newBackend/go_backend/db"
)

type TopSellingData struct {
	MonthYear    string  `json:"month_year"`
	TotalOrders  int64   `json:"total_orders"`
	TotalRevenue float64 `json:"total_revenue"`
}

func GetTopSellingDailyHandler(c *fiber.Ctx) error {
	query := "SELECT * FROM get_sales_by_duration('daily');"
	rows, err := db.Pool.Query(context.Background(), query)
	if err != nil {
		log.Println("Error fetching daily sales:", err)
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

		if data.TotalOrders == 0 {
			data.TotalOrders = 0 
		}
		if data.TotalRevenue == 0 {
			data.TotalRevenue = 0.0 
		}

		results = append(results, data)
	}
	return c.JSON(results)
}

func GetTopSellingWeeklyHandler(c *fiber.Ctx) error {
	query := "SELECT * FROM get_sales_weekly();"
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

		if data.TotalOrders == 0 {
			data.TotalOrders = 0 
		}
		if data.TotalRevenue == 0 {
			data.TotalRevenue = 0.0 
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

		if data.TotalOrders == 0 {
			data.TotalOrders = 0 
		}
		if data.TotalRevenue == 0 {
			data.TotalRevenue = 0.0 
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

		if data.TotalOrders == 0 {
			data.TotalOrders = 0 
		}
		if data.TotalRevenue == 0 {
			data.TotalRevenue = 0.0 
		}

		results = append(results, data)
	}

	return c.JSON(results)
}
