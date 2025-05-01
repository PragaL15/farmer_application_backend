package handlers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/PragaL15/go_newBackend/go_backend/db"
)

type WholesellerEntry struct {
	ProductID     int       `json:"product_id"`
	Quality       string    `json:"quality"`
	Wastage       string    `json:"wastage"`
	Quantity      float64   `json:"quantity"`
	Price         float64   `json:"price"`
	Datetime      time.Time `json:"datetime"`
	WholesellerID int       `json:"wholeseller_id"`
	MandiID       int       `json:"mandi_id"`
	WarehouseID   int       `json:"warehouse_id"`  // ✅ New field
	UnitID        int       `json:"unit_id"`       // ✅ New field
}

// InsertWholesellerEntryHandler godoc
// @Summary      Insert a wholeseller entry
// @Description  Inserts a new entry into wholeseller_entry_table using the stored procedure
// @Tags         Market Opportunities
// @Accept       json
// @Produce      json
// @Param        entry  body      WholesellerEntry  true  "Wholeseller Entry"
// @Success      200    {object}  map[string]any
// @Failure      400    {object}  map[string]string
// @Failure      500    {object}  map[string]string
// @Router       /api/wholeseller-entry [post]
func InsertWholesellerEntryHandler(c *fiber.Ctx) error {
	var entry WholesellerEntry

	if err := c.BodyParser(&entry); err != nil {
		log.Println("Invalid input:", err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error":  "Invalid input",
			"detail": err.Error(),
		})
	}

	var entryID int
	query := `
		SELECT business_schema.insert_wholeseller_entry($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	`
	err := db.Pool.QueryRow(
		context.Background(),
		query,
		entry.ProductID, entry.Quality, entry.Wastage,
		entry.Quantity, entry.Price, entry.Datetime,
		entry.WholesellerID, entry.MandiID,
		entry.WarehouseID, entry.UnitID,
	).Scan(&entryID)

	if err != nil {
		log.Println("Failed to insert entry:", err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error":  "Failed to insert entry",
			"detail": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message":  "Entry inserted successfully",
		"entry_id": entryID,
	})
}
