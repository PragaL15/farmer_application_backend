package handlers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/PragaL15/go_newBackend/go_backend/db"
)

// Request body structure
type OfferRequest struct {
	OrderID              int     `json:"order_id"`
	WholesellerID        int     `json:"wholeseller_id"`
	OfferedPrice         float64 `json:"offered_price"`
	ProposedDeliveryDate string  `json:"proposed_delivery_date"` // "YYYY-MM-DD"
	Message              string  `json:"message,omitempty"`
}

// Response structure
type OfferResponse struct {
	OfferID int `json:"offer_id"`
}

// CreateWholesellerOfferHandler godoc
// @Summary      Make a new price offer on an existing bulk order
// @Description  Inserts a new offer for the given order by the wholeseller
// @Tags         Offers
// @Accept       json
// @Produce      json
// @Param        offer  body  OfferRequest  true  "Offer Details"
// @Success      200    {object}  OfferResponse
// @Failure      400    {object}  map[string]string
// @Failure      500    {object}  map[string]string
// @Router       /api/wholeseller/offer [post]
func CreateWholesellerOfferHandler(c *fiber.Ctx) error {
	var req OfferRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request format",
		})
	}

	// Validate date
	deliveryDate, err := time.Parse("2006-01-02", req.ProposedDeliveryDate)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid delivery date format. Use YYYY-MM-DD",
		})
	}

	query := `
		SELECT business_schema.insert_wholeseller_offer($1, $2, $3, $4, $5)
	`

	var offerID int
	err = db.Pool.QueryRow(context.Background(), query,
		req.OrderID,
		req.WholesellerID,
		req.OfferedPrice,
		deliveryDate,
		req.Message,
	).Scan(&offerID)

	if err != nil {
		log.Println("Error inserting offer:", err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error":  "Failed to insert offer",
			"detail": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(OfferResponse{OfferID: offerID})
}
