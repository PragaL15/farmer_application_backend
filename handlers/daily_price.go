package handlers

import (
	"context"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v4"
	"github.com/PragaL15/go_newBackend/go_backend/db"
)

type PriceData struct {
	ProductID     int     `json:"product_id"`
	Price         float64 `json:"price"`
	UnitID        int     `json:"unit_id"`
	WholesellerID int     `json:"wholeseller_id"`
	Currency      string  `json:"currency"`
	CreatedAt     string  `json:"created_at"`
	UpdatedAt     string  `json:"updated_at"`
	Remarks       string  `json:"remarks"`
}

// Fetch price data based on product, unit, and wholesaler
func fetchPriceData(productID, unitID, wholesellerID string) (*PriceData, error) {
	query := `SELECT * FROM business_schema.fetch_price_data($1, $2, $3)`

	row := db.Pool.QueryRow(context.Background(), query, productID, unitID, wholesellerID)

	var priceData PriceData
	var createdAt, updatedAt time.Time

	err := row.Scan(
		&priceData.ProductID, &priceData.Price, &priceData.UnitID, &priceData.WholesellerID,
		&priceData.Currency, &createdAt, &updatedAt, &priceData.Remarks,
	)
	if err != nil {
		return nil, err
	}

	priceData.CreatedAt = createdAt.Format(time.RFC3339)
	priceData.UpdatedAt = updatedAt.Format(time.RFC3339)
	return &priceData, nil
}

// Insert price data with branch mapping
func insertPriceData(req PriceData) error {
	query := `SELECT business_schema.insert_daily_price($1, $2, $3, $4, $5, $6)`

	_, err := db.Pool.Exec(context.Background(), query,
		req.WholesellerID, req.ProductID, req.UnitID,
		req.Price, req.Currency, req.Remarks,
	)
	return err
}

// Update price data for a wholesaler in a specific branch
func updatePriceData(req PriceData) error {
	query := `SELECT business_schema.update_daily_price($1, $2, $3, $4, $5, $6, $7)`

	_, err := db.Pool.Exec(context.Background(), query,
		req.WholesellerID, req.ProductID, req.UnitID,
		req.Price, req.Currency, time.Now(), req.Remarks,
	)
	return err
}

// Get Price Handler
func GetPriceHandler(c *fiber.Ctx) error {
	productID := c.Query("product_id")
	unitID := c.Query("unit_id")
	wholesellerID := c.Query("wholeseller_id")

	if productID == "" || unitID == "" || wholesellerID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Missing query parameters"})
	}

	priceData, err := fetchPriceData(productID, unitID, wholesellerID)
	if err == pgx.ErrNoRows {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Price data not found"})
	} else if err != nil {
		log.Println("Database error:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal server error"})
	}

	return c.JSON(priceData)
}

// Insert Price Handler
func InsertPriceHandler(c *fiber.Ctx) error {
	var req PriceData
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	if err := insertPriceData(req); err != nil {
		log.Println("Database insert error:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to insert price"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Price inserted successfully"})
}

// Update Price Handler
func UpdatePriceHandler(c *fiber.Ctx) error {
	var req PriceData
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	if err := updatePriceData(req); err != nil {
		log.Println("Database update error:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update price"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Price updated successfully"})
}
