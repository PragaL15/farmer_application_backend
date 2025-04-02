package handlers

import (
	"context"
	"log"
	"strconv"
  "database/sql"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v4"
	"github.com/PragaL15/go_newBackend/go_backend/db"
)

type Order struct {
	OrderID              int64     `json:"order_id"`
	DateOfOrder          string    `json:"date_of_order"`
	OrderStatus          int       `json:"order_status"`
	ExpectedDeliveryDate string    `json:"expected_delivery_date"`
	ActualDeliveryDate   string    `json:"actual_delivery_date"`
	RetailerID           int       `json:"retailer_id"`
	WholesellerIDs       []int     `json:"wholeseller_ids"`
	SelectedOfferID      *int64    `json:"selected_offer_id,omitempty"`
	LocationID           int       `json:"location_id"`
	StateID              int       `json:"state_id"`
	Address              string    `json:"address"`
	Pincode              string    `json:"pincode"`
	TotalOrderAmount     float64   `json:"total_order_amount"`
	DiscountAmount       float64   `json:"discount_amount"`
	TaxAmount            float64   `json:"tax_amount"`
	FinalAmount          float64   `json:"final_amount"`
	DesiredDeliveryDate  string    `json:"desired_delivery_date"`
	DeliveryDeadline     string    `json:"delivery_deadline"`
	MaxPriceLimit        float64   `json:"max_price_limit"`
	RetailerContact      string    `json:"retailer_contact"`
	ProductIDs           []int64   `json:"product_ids"`
	Quantities           []float64 `json:"quantities"`
	UnitIDs              []int     `json:"unit_ids"`
	MaxItemPrices        []float64 `json:"max_item_prices"`
}

type OrderResponse struct {
	OrderID int64  `json:"order_id"`
	Status  string `json:"status"`
	Message string `json:"message"`
}
func CreateRetailerOrderHandler(c *fiber.Ctx) error {
	var req Order
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(OrderResponse{
			Status:  "error",
			Message: "Invalid request body: " + err.Error(),
		})
	}

	if req.RetailerID == 0 || req.RetailerContact == "" || req.Pincode == "" || req.Address == "" {
		return c.Status(fiber.StatusBadRequest).JSON(OrderResponse{
			Status:  "error",
			Message: "Missing required fields (retailer_id, retailer_contact, pincode, address)",
		})
	}

	if len(req.ProductIDs) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(OrderResponse{
			Status:  "error",
			Message: "At least one product is required",
		})
	}

	if len(req.ProductIDs) != len(req.Quantities) || 
	   len(req.ProductIDs) != len(req.UnitIDs) || 
	   len(req.ProductIDs) != len(req.MaxItemPrices) {
		return c.Status(fiber.StatusBadRequest).JSON(OrderResponse{
			Status:  "error",
			Message: "Product IDs, quantities, unit IDs, and max prices arrays must have the same length",
		})
	}

	var orderID sql.NullInt64
	var status, message string

	query := `
		SELECT order_id, status, message FROM business_schema.create_retailer_order(
			$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12
		)`

	err := db.Pool.QueryRow(context.Background(), query,
		req.RetailerID,
		req.RetailerContact,
		req.Pincode,
		req.Address,
		req.MaxPriceLimit,
		req.DesiredDeliveryDate,
		req.DeliveryDeadline,
		req.ProductIDs,
		req.Quantities,
		req.UnitIDs,
		req.MaxItemPrices,
		req.WholesellerIDs,
	).Scan(&orderID, &status, &message)

	if err != nil {
		log.Printf("Order creation failed: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(OrderResponse{
			Status:  "error",
			Message: "Database operation failed: " + err.Error(),
		})
	}

	if !orderID.Valid {
		return c.Status(fiber.StatusBadRequest).JSON(OrderResponse{
			Status:  "error",
			Message: "Failed to create order: order ID is NULL",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(OrderResponse{
		OrderID: orderID.Int64,
		Status:  status,
		Message: message,
	})
}



func GetOrderDetailsHandler(c *fiber.Ctx) error {
	orderIDStr := c.Params("id")
	if orderIDStr == "" {
		return c.Status(fiber.StatusBadRequest).JSON(OrderResponse{
			Status:  "error",
			Message: "Order ID is required",
		})
	}

	orderID, err := strconv.ParseInt(orderIDStr, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(OrderResponse{
			Status:  "error",
			Message: "Invalid order ID format",
		})
	}

	var order Order
	orderQuery := `
		SELECT 
			order_id, date_of_order, order_status, 
			expected_delivery_date, actual_delivery_date,
			retailer_id, wholeseller_offer_id, location_id,
			state_id, address, pincode, total_order_amount,
			discount_amount, tax_amount, final_amount,
			desired_delivery_date, delivery_deadline, max_price_limit,
			retailer_contact
		FROM business_schema.order_table
		WHERE order_id = $1`

	err = db.Pool.QueryRow(context.Background(), orderQuery, orderID).Scan(
		&order.OrderID, &order.DateOfOrder, &order.OrderStatus,
		&order.ExpectedDeliveryDate, &order.ActualDeliveryDate,
		&order.RetailerID, &order.SelectedOfferID, &order.LocationID,
		&order.StateID, &order.Address, &order.Pincode, &order.TotalOrderAmount,
		&order.DiscountAmount, &order.TaxAmount, &order.FinalAmount,
		&order.DesiredDeliveryDate, &order.DeliveryDeadline, &order.MaxPriceLimit,
		&order.RetailerContact,
	)

	if err == pgx.ErrNoRows {
		return c.Status(fiber.StatusNotFound).JSON(OrderResponse{
			Status:  "error",
			Message: "Order not found",
		})
	} else if err != nil {
		log.Printf("Failed to fetch order: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(OrderResponse{
			Status:  "error",
			Message: "Failed to fetch order details",
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
		if err = rows.Err(); err != nil {
			log.Printf("Error processing wholeseller rows: %v", err)
		}
	}

	productQuery := `
		SELECT product_id, quantity, unit_id, max_item_price
		FROM business_schema.order_item_table
		WHERE order_id = $1`

	productRows, err := db.Pool.Query(context.Background(), productQuery, orderID)
	if err != nil {
		log.Printf("Failed to fetch order items: %v", err)
	} else {
		defer productRows.Close()
		for productRows.Next() {
			var productID int64
			var quantity float64
			var unitID int
			var maxItemPrice float64

			if err := productRows.Scan(&productID, &quantity, &unitID, &maxItemPrice); err == nil {
				order.ProductIDs = append(order.ProductIDs, productID)
				order.Quantities = append(order.Quantities, quantity)
				order.UnitIDs = append(order.UnitIDs, unitID)
				order.MaxItemPrices = append(order.MaxItemPrices, maxItemPrice)
			}
		}
		if err = productRows.Err(); err != nil {
			log.Printf("Error processing product rows: %v", err)
		}
	}

	return c.JSON(order)
}