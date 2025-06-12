package TrendHandlers

import (
	"context"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgtype"

	"farmerapp/go_backend/db"
)

// Struct to scan DB values (supports NULL)
type SalesData struct {
	MonthYear    pgtype.Text
	TotalOrders  pgtype.Int8
	TotalRevenue pgtype.Float8
}

// Struct to send JSON response
type SalesDataJSON struct {
	MonthYear    string  `json:"month_year"`
	TotalOrders  int64   `json:"total_orders"`
	TotalRevenue float64 `json:"total_revenue"`
}

// Convert scanned SalesData into clean JSON format
func convertToJSON(data SalesData) SalesDataJSON {
	return SalesDataJSON{
		MonthYear:    data.MonthYear.String,
		TotalOrders:  data.TotalOrders.Int64,
		TotalRevenue: data.TotalRevenue.Float64,
	}
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

	var results []SalesDataJSON
	for rows.Next() {
		var data SalesData
		if err := rows.Scan(&data.MonthYear, &data.TotalOrders, &data.TotalRevenue); err != nil {
			log.Println("Error scanning monthly row:", err)
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"error":  "Failed to scan row",
				"detail": err.Error(),
			})
		}
		results = append(results, convertToJSON(data))
	}
	return c.JSON(results)
}

func GetSalesWeeklyHandler(c *fiber.Ctx) error {
	query := "SELECT * FROM get_weekday_sales();"
	rows, err := db.Pool.Query(context.Background(), query)
	if err != nil {
		log.Println("Error fetching weekly sales:", err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error":  "Database query failed",
			"detail": err.Error(),
		})
	}
	defer rows.Close()

	var results []SalesDataJSON
	for rows.Next() {
		var data SalesData
		if err := rows.Scan(&data.MonthYear, &data.TotalOrders, &data.TotalRevenue); err != nil {
			log.Println("Error scanning weekly row:", err)
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"error":  "Failed to scan row",
				"detail": err.Error(),
			})
		}
		results = append(results, convertToJSON(data))
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

	var results []SalesDataJSON
	for rows.Next() {
		var data SalesData
		if err := rows.Scan(&data.MonthYear, &data.TotalOrders, &data.TotalRevenue); err != nil {
			log.Println("Error scanning yearly row:", err)
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"error":  "Failed to scan row",
				"detail": err.Error(),
			})
		}
		results = append(results, convertToJSON(data))
	}
	return c.JSON(results)
}
