package handlers

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v4"
	"github.com/PragaL15/go_newBackend/go_backend/db"
)

type Order struct {
	OrderID              int64    `json:"order_id"`
	DateOfOrder          string   `json:"date_of_order"`
	OrderStatus          int      `json:"order_status"`
	ExpectedDeliveryDate string   `json:"expected_delivery_date"`
	ActualDeliveryDate   string   `json:"actual_delivery_date"`
	RetailerID           int      `json:"retailer_id"`
	WholesellerIDs       []int    `json:"wholeseller_ids"` // Changed to array
	SelectedOfferID      *int64   `json:"selected_offer_id,omitempty"` // Nullable
	LocationID           int      `json:"location_id"`
	StateID              int      `json:"state_id"`
	Address              string   `json:"address"`
	Pincode              string   `json:"pincode"`
	TotalOrderAmount     float64  `json:"total_order_amount"`
	DiscountAmount       float64  `json:"discount_amount"`
	TaxAmount            float64  `json:"tax_amount"`
	FinalAmount          float64  `json:"final_amount"`
	DesiredDeliveryDate  string   `json:"desired_delivery_date"`
	DeliveryDeadline     string   `json:"delivery_deadline"`
	MaxPriceLimit        float64  `json:"max_price_limit"`
	RetailerContact      string   `json:"retailer_contact"`
	ProductIDs           []int64  `json:"product_ids"`
	Quantities           []float64 `json:"quantities"`
	UnitIDs              []int     `json:"unit_ids"`
	MaxItemPrices        []float64 `json:"max_item_prices"`
}

func CreateRetailerOrderHandler(c *fiber.Ctx) error {
	var req Order
	if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"error": "Invalid request body",
					"details": err.Error(),
			})
	}

	if len(req.ProductIDs) != len(req.Quantities) || 
		 len(req.ProductIDs) != len(req.UnitIDs) || 
		 len(req.ProductIDs) != len(req.MaxItemPrices) {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"error": "Product IDs, quantities, unit IDs and max prices arrays must have same length",
			})
	}

	productIDs := convertToInt64Array(req.ProductIDs)
	quantities := convertToFloat64Array(req.Quantities)
	unitIDs := convertToIntArray(req.UnitIDs)
	maxItemPrices := convertToFloat64Array(req.MaxItemPrices)
	wholesellerIDs := convertToIntArray(req.WholesellerIDs)

	var (
			orderID int64
			status  string
			message string
	)

	query := `
			SELECT * FROM business_schema.create_retailer_order(
					$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11
			)`

	err := db.Pool.QueryRow(context.Background(), query,
			req.RetailerID,
			req.RetailerContact,
			req.Pincode,
			req.Address,
			req.MaxPriceLimit,
			req.DesiredDeliveryDate,
			req.DeliveryDeadline,
			productIDs,
			quantities,
			unitIDs,
			maxItemPrices,
	).Scan(&orderID, &status, &message)

	if err != nil {
			log.Printf("Order creation failed: %v", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error": "Failed to create order",
					"details": err.Error(),
			})
	}

	if status != "success" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"error": message,
					"order_id": orderID,
					"status": status,
			})
	}

	if len(req.WholesellerIDs) > 0 {
			_, err = db.Pool.Exec(context.Background(),
					`INSERT INTO business_schema.order_wholeseller_mapping
					 (order_id, wholeseller_id) 
					 SELECT $1, unnest($2::int[])`,
					orderID, wholesellerIDs,
			)
			if err != nil {
					log.Printf("Failed to create wholeseller mappings: %v", err)
			}
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"message": message,
			"order_id": orderID,
			"status": status,
	})
}

func convertToInt64Array(slice []int64) []interface{} {
	result := make([]interface{}, len(slice))
	for i, v := range slice {
		result[i] = v
	}
	return result
}

func convertToIntArray(slice []int) []interface{} {
	result := make([]interface{}, len(slice))
	for i, v := range slice {
		result[i] = v
	}
	return result
}

func convertToFloat64Array(slice []float64) []interface{} {
	result := make([]interface{}, len(slice))
	for i, v := range slice {
		result[i] = v
	}
	return result
}

func GetOrderDetailsHandler(c *fiber.Ctx) error {
	orderID := c.Params("id")
	if orderID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Order ID is required",
		})
	}

	var order Order
	orderQuery := `
		SELECT 
			order_id, date_of_order, order_status, 
			retailer_id, wholeseller_offer_id,
			total_order_amount, discount_amount, tax_amount, final_amount,
			delivery_pincode, delivery_address,
			desired_delivery_date, delivery_deadline
		FROM business_schema.order_table
		WHERE order_id = $1`

	err := db.Pool.QueryRow(context.Background(), orderQuery, orderID).Scan(
		&order.OrderID, &order.DateOfOrder, &order.OrderStatus,
		&order.RetailerID, &order.SelectedOfferID,
		&order.TotalOrderAmount, &order.DiscountAmount, &order.TaxAmount, &order.FinalAmount,
		&order.Pincode, &order.Address,
		&order.DesiredDeliveryDate, &order.DeliveryDeadline,
	)

	if err == pgx.ErrNoRows {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Order not found",
		})
	} else if err != nil {
		log.Printf("Failed to fetch order: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch order details",
		})
	}

	wholesellerQuery := `
		SELECT wholeseller_id 
		FROM business_schema.order_wholeseller_mapping
		WHERE order_id = $1 AND status = 1`

	rows, err := db.Pool.Query(context.Background(), wholesellerQuery, orderID)
	if err != nil {
		log.Printf("Failed to fetch wholesellers: %v", err)
	} else {
		defer rows.Close()
		for rows.Next() {
			var id int
			if err := rows.Scan(&id); err == nil {
				order.WholesellerIDs = append(order.WholesellerIDs, id)
			}
		}
	}

	offersQuery := `
		SELECT 
			offer_id, wholeseller_id, offered_price, 
			proposed_delivery_date, offer_status
		FROM business_schema.wholeseller_offers
		WHERE order_id = $1`

	offersRows, err := db.Pool.Query(context.Background(), offersQuery, orderID)
	if err != nil {
		log.Printf("Failed to fetch offers: %v", err)
	} else {
		defer offersRows.Close()
		
	}

	return c.JSON(order)
}