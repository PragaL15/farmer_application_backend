
package handlers

import (
	"context"
	"database/sql"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/PragaL15/go_newBackend/go_backend/db"
)

type DailyPriceUpdate struct {
	ProductID      int     `json:"product_id"`
	ProductName    string  `json:"product_name"`
	Price          float64 `json:"price"`
	Status         int     `json:"status"`
	UnitID         int     `json:"unit_id"`
	UnitName       string  `json:"unit_name"`
	WholesellerID  int     `json:"wholeseller_id"`
	WholesellerName string `json:"wholeseller_name"`
	BMandiID       int     `json:"b_mandi_id"`
	MandiName      string  `json:"mandi_name"`
}

func GetDailyPriceUpdates(c *fiber.Ctx) error {
	rows, err := db.Pool.Query(context.Background(), "SELECT * FROM get_daily_price_update()")
	if err != nil {
		log.Printf("Failed to fetch daily price updates: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch data"})
	}
	defer rows.Close()

	var updates []DailyPriceUpdate

	for rows.Next() {
		var update DailyPriceUpdate
		var price sql.NullFloat64
		var productName, unitName, wholesellerName, mandiName sql.NullString

		err := rows.Scan(
			&update.ProductID, &productName, &price, &update.Status,
			&update.UnitID, &unitName, &update.WholesellerID, &wholesellerName,
			&update.BMandiID, &mandiName,
		)
		if err != nil {
			log.Printf("Error scanning row: %v", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error processing data"})
		}

		update.ProductName = formatNullString(productName)
		update.Price = formatNullFloat64(price)
		update.UnitName = formatNullString(unitName)
		update.WholesellerName = formatNullString(wholesellerName)
		update.MandiName = formatNullString(mandiName)

		updates = append(updates, update)
	}

	return c.JSON(updates)
}

type UpdatePriceRequest struct {
	ProductID     int     `json:"product_id"`
	Price         float64 `json:"price"`
	Status        int     `json:"status"`
	UnitID        int     `json:"unit_id"`
	WholesellerID int     `json:"wholeseller_id"`
}

func UpdateDailyPrice(c *fiber.Ctx) error {
	var req UpdatePriceRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	_, err := db.Pool.Exec(
		context.Background(),
		"UPDATE daily_price_update SET price = $1, status = $2, unit_id = $3, wholeseller_id = $4 WHERE product_id = $5",
		req.Price, req.Status, req.UnitID, req.WholesellerID, req.ProductID,
	)
	if err != nil {
		log.Printf("Error updating data: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error updating data"})
	}

	return c.JSON(fiber.Map{"message": "Update successful"})
}

func formatNullStrings(ns sql.NullString) string {
	if ns.Valid {
		return ns.String
	}
	return ""
}

func formatNullFloat64s(nf sql.NullFloat64) float64 {
	if nf.Valid {
		return nf.Float64
	}
	return 0.0
}
