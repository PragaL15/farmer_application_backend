package TrendHandlers

import (
	"context"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/PragaL15/go_newBackend/go_backend/db"
)

type WholesellerPrice struct {
	WholesellerID int     `json:"wholeseller_id"`
	PricePerKg    float64 `json:"price_per_kg"`
}

type GroupedPriceComparison struct {
	ProductName string              `json:"product_name"`
	Prices      []WholesellerPrice  `json:"prices"`
}

func GetWholesellerPriceComparisonHandler(c *fiber.Ctx) error {
	productIDsParam := c.Query("product_ids")
	if productIDsParam == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Query param 'product_ids' is required (e.g. ?product_ids=2,3,4)",
		})
	}

	formatted := "{" + productIDsParam + "}"

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

	groupedData := make(map[string][]WholesellerPrice)

	for rows.Next() {
		var productName string
		var wholesellerID int
		var pricePerKg float64

		if err := rows.Scan(&productName, &wholesellerID, &pricePerKg); err != nil {
			log.Println("Row scan error:", err)
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"error":  "Failed to scan result row",
				"detail": err.Error(),
			})
		}

		groupedData[productName] = append(groupedData[productName], WholesellerPrice{
			WholesellerID: wholesellerID,
			PricePerKg:    pricePerKg,
		})
	}

	// If there's only one product, return it as an object
	if len(groupedData) == 1 {
		for productName, prices := range groupedData {
			single := GroupedPriceComparison{
				ProductName: productName,
				Prices:      prices,
			}
			return c.JSON(single)
		}
	}

	// Else return as array
	var response []GroupedPriceComparison
	for productName, prices := range groupedData {
		response = append(response, GroupedPriceComparison{
			ProductName: productName,
			Prices:      prices,
		})
	}

	return c.JSON(response)
}
