package TrendHandlers

import (
	"context"
	"log"
	"net/http"
	"strings"
	"github.com/gofiber/fiber/v2"
	"github.com/PragaL15/go_newBackend/go_backend/db"
)

type PriceComparison struct {
	ProductName    string  `json:"product_name"`
	WholesellerID  int     `json:"wholeseller_id"`
	PricePerKg     float64 `json:"price_per_kg"`
}

func GetWholesellerPriceComparisonHandler(c *fiber.Ctx) error {
	productIDsParam := c.Query("product_ids")
	if productIDsParam == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Query param 'product_ids' is required (e.g. ?product_ids=2,3,4)",
		})
	}

	formatted := "{" + strings.ReplaceAll(productIDsParam, ",", ",") + "}"

	query := `SELECT * FROM business_schema.wholeseller_price_comparison($1::int[]);`

	rows, err := db.Pool.Query(context.Background(), query, formatted)
	if err != nil {
		log.Println("Error executing query:", err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error":  "Database query failed",
			"detail": err.Error(),
		})
	}
	defer rows.Close()

	var results []PriceComparison

	for rows.Next() {
		var data PriceComparison
		if err := rows.Scan(&data.ProductName, &data.WholesellerID, &data.PricePerKg); err != nil {
			log.Println("Row scan error:", err)
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"error":  "Failed to scan result row",
				"detail": err.Error(),
			})
		}
		results = append(results, data)
	}

	return c.JSON(results)
}
