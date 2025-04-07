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
	dailyPriceID := c.Query("daily_price_id")

	query := `
		SELECT daily_price_id, product_id, product_name, category_id, category_name, price, unit_id, unit_name, 
		       wholeseller_id, currency, created_at, updated_at, remarks, b_branch_id, b_shop_name
		FROM business_schema.get_all_daily_price_updates()
	`

	var rows pgx.Rows
	var err error

	if dailyPriceID != "" {
		query += " WHERE daily_price_id = $1"
		rows, err = db.Pool.Query(context.Background(), query, dailyPriceID)
	} else {
		rows, err = db.Pool.Query(context.Background(), query)
	}

	if err != nil {
		log.Println("Query error:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Database error"})
	}
	defer rows.Close()

	var prices []struct {
		DailyPriceID   int64     `json:"daily_price_id"`
		ProductID      int       `json:"product_id"`
		ProductName    string    `json:"product_name"`
		CategoryID     int       `json:"category_id"`
		CategoryName   string    `json:"category_name"`
		Price          float64   `json:"price"`
		UnitID         int       `json:"unit_id"`
		UnitName       string    `json:"unit_name"`
		WholesellerID  int       `json:"wholeseller_id"`
		Currency       string    `json:"currency"`
		CreatedAt      time.Time `json:"created_at"`
		UpdatedAt      time.Time `json:"updated_at"`
		Remarks        string    `json:"remarks"`
		BranchID       int       `json:"b_branch_id"`
		BranchShopName string    `json:"b_shop_name"`
	}

	for rows.Next() {
		var p struct {
			DailyPriceID   int64     `json:"daily_price_id"`
			ProductID      int       `json:"product_id"`
			ProductName    string    `json:"product_name"`
			CategoryID     int       `json:"category_id"`
			CategoryName   string    `json:"category_name"`
			Price          float64   `json:"price"`
			UnitID         int       `json:"unit_id"`
			UnitName       string    `json:"unit_name"`
			WholesellerID  int       `json:"wholeseller_id"`
			Currency       string    `json:"currency"`
			CreatedAt      time.Time `json:"created_at"`
			UpdatedAt      time.Time `json:"updated_at"`
			Remarks        string    `json:"remarks"`
			BranchID       int       `json:"b_branch_id"`
			BranchShopName string    `json:"b_shop_name"`
		}

		if err := rows.Scan(
			&p.DailyPriceID,
			&p.ProductID,
			&p.ProductName,
			&p.CategoryID,
			&p.CategoryName,
			&p.Price,
			&p.UnitID,
			&p.UnitName,
			&p.WholesellerID,
			&p.Currency,
			&p.CreatedAt,
			&p.UpdatedAt,
			&p.Remarks,
			&p.BranchID,
			&p.BranchShopName,
		); err != nil {
			log.Println("Row scan error:", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Data scan error"})
		}

		prices = append(prices, p)
	}

	if len(prices) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "No records found"})
	}

	return c.JSON(prices)
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
