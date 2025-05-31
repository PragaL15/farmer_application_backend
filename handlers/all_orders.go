package handlers

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"time"

	"farmerapp/go_backend/db"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v4"
)

// Order represents the structure of an order request
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

// OrderResponse represents the response structure for order operations
type OrderResponse struct {
	OrderID int64  `json:"order_id,omitempty"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

// ProductDetail represents detailed information about a product in an order
type ProductDetail struct {
	ProductID    int64   `json:"product_id"`
	ProductName  string  `json:"product_name"`
	CategoryID   int     `json:"category_id"`
	CategoryName string  `json:"category_name"`
	Quantity     float64 `json:"quantity"`
	UnitID       int     `json:"unit_id"`
	UnitName     string  `json:"unit_name"`
	MaxPrice     float64 `json:"max_price"`
}

// OrderItemDetail represents detailed information about an order item
type OrderItemDetail struct {
	OrderItemID  int64   `json:"order_item_id"`
	ProductID    int64   `json:"product_id"`
	ProductName  string  `json:"product_name"`
	Quantity     float64 `json:"quantity"`
	UnitID       int     `json:"unit_id"`
	UnitName     string  `json:"unit_name"`
	MaxItemPrice float64 `json:"max_item_price"`
}

// OrderDetailResponse represents the detailed response for an order
type OrderDetailResponse struct {
	OrderID            int64           `json:"order_id"`
	DateOfOrder        time.Time       `json:"date_of_order"`
	OrderStatus        int             `json:"order_status"`
	ActualDeliveryDate *time.Time      `json:"actual_delivery_date,omitempty"`
	RetailerID         int             `json:"retailer_id"`
	ShopName           string          `json:"shop_name"`
	WholesellerIDs     []int           `json:"wholeseller_ids,omitempty"`
	TotalOrderAmount   float64         `json:"total_order_amount"`
	DiscountAmount     float64         `json:"discount_amount"`
	TaxAmount          float64         `json:"tax_amount"`
	FinalAmount        float64         `json:"final_amount"`
	Products           []ProductDetail `json:"products"`
}

// CompletedOrderDetail represents a completed order with its items
type CompletedOrderDetail struct {
	OrderID          int64             `json:"order_id"`
	TotalOrderAmount float64           `json:"total_order_amount"`
	OrderItems       []OrderItemDetail `json:"order_items"`
}

// CreateRetailerOrderHandler handles the creation of a new retailer order
func CreateRetailerOrderHandler(c *fiber.Ctx) error {
	var req Order
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(OrderResponse{
			Status:  "error",
			Message: "Invalid request body: " + err.Error(),
		})
	}

	// Validate required fields
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

	var orderID int64
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

	return c.Status(fiber.StatusCreated).JSON(OrderResponse{
		OrderID: orderID,
		Status:  status,
		Message: message,
	})
}

// GetOrderDetailsHandler retrieves detailed information about a specific order
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

	// Get basic order information
	var order struct {
		OrderID            int64
		DateOfOrder        time.Time
		OrderStatus        int
		ActualDeliveryDate *time.Time
		RetailerID         int
		TotalOrderAmount   *float64
		DiscountAmount     *float64
		TaxAmount          *float64
		FinalAmount        *float64
	}

	orderQuery := `
		SELECT
			order_id, date_of_order, order_status,
			actual_delivery_date, retailer_id,
			total_order_amount, discount_amount,
			tax_amount, final_amount
		FROM business_schema.order_table
		WHERE order_id = $1`

	err = db.Pool.QueryRow(context.Background(), orderQuery, orderID).Scan(
		&order.OrderID, &order.DateOfOrder, &order.OrderStatus,
		&order.ActualDeliveryDate, &order.RetailerID,
		&order.TotalOrderAmount, &order.DiscountAmount,
		&order.TaxAmount, &order.FinalAmount,
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

	// Get shop name
	var shopName string
	err = db.Pool.QueryRow(context.Background(),
		"SELECT b_shop_name FROM admin_schema.business_branch_table WHERE bid = $1",
		order.RetailerID).Scan(&shopName)
	if err != nil {
		log.Printf("Failed to fetch shop name: %v", err)
		shopName = ""
	}

	// Get wholeseller IDs
	wholesellerIDs, err := getWholesellerIDs(orderID)
	if err != nil {
		log.Printf("Failed to fetch wholesellers: %v", err)
	}

	// Get order products
	products, err := getOrderProducts(orderID)
	if err != nil {
		log.Printf("Failed to fetch order items: %v", err)
	}

	// Prepare response
	response := OrderDetailResponse{
		OrderID:            order.OrderID,
		DateOfOrder:        order.DateOfOrder,
		OrderStatus:        order.OrderStatus,
		ActualDeliveryDate: order.ActualDeliveryDate,
		RetailerID:         order.RetailerID,
		ShopName:           shopName,
		WholesellerIDs:     wholesellerIDs,
		TotalOrderAmount:   derefFloat(order.TotalOrderAmount),
		DiscountAmount:     derefFloat(order.DiscountAmount),
		TaxAmount:          derefFloat(order.TaxAmount),
		FinalAmount:        derefFloat(order.FinalAmount),
		Products:           products,
	}

	return c.JSON(response)
}

// GetAllCompletedOrderItemHandler retrieves all completed order items
func GetAllCompletedOrderItemHandler(c *fiber.Ctx) error {
	query := `SELECT * FROM business_schema.get_order_details_by_status_6();`

	rows, err := db.Pool.Query(context.Background(), query)
	if err != nil {
		log.Println("Error executing order details query:", err)
		return c.Status(http.StatusInternalServerError).JSON(OrderResponse{
			Status:  "error",
			Message: "Database query failed: " + err.Error(),
		})
	}
	defer rows.Close()

	orderDetailsMap := make(map[int64]*CompletedOrderDetail)

	for rows.Next() {
		var orderID int64
		var totalAmount float64
		var item OrderItemDetail

		if err := rows.Scan(
			&orderID,
			&totalAmount,
			&item.OrderItemID,
			&item.ProductID,
			&item.ProductName,
			&item.Quantity,
			&item.UnitID,
			&item.MaxItemPrice,
			&item.UnitName,
		); err != nil {
			log.Println("Error scanning order row:", err)
			return c.Status(http.StatusInternalServerError).JSON(OrderResponse{
				Status:  "error",
				Message: "Failed to scan row: " + err.Error(),
			})
		}

		if existingOrder, exists := orderDetailsMap[orderID]; exists {
			existingOrder.OrderItems = append(existingOrder.OrderItems, item)
		} else {
			orderDetailsMap[orderID] = &CompletedOrderDetail{
				OrderID:          orderID,
				TotalOrderAmount: totalAmount,
				OrderItems:       []OrderItemDetail{item},
			}
		}
	}

	// Convert map to slice
	var orderDetails []CompletedOrderDetail
	for _, order := range orderDetailsMap {
		orderDetails = append(orderDetails, *order)
	}

	return c.Status(http.StatusOK).JSON(orderDetails)
}

// Helper function to get wholeseller IDs for an order
func getWholesellerIDs(orderID int64) ([]int, error) {
	var wholesellerIDs []int
	query := `
		SELECT wholeseller_id
		FROM business_schema.order_wholeseller_mapping
		WHERE order_id = $1 AND status = 1`

	rows, err := db.Pool.Query(context.Background(), query, orderID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		wholesellerIDs = append(wholesellerIDs, id)
	}

	return wholesellerIDs, rows.Err()
}

// Helper function to get products for an order
func getOrderProducts(orderID int64) ([]ProductDetail, error) {
	var products []ProductDetail
	query := `
		SELECT
			oi.product_id, mp.product_name, mp.category_id,
			mpc.category_name, oi.quantity,
			oi.unit_id, ut.unit_name, oi.max_item_price
		FROM business_schema.order_item_table oi
		JOIN admin_schema.master_product mp ON oi.product_id = mp.product_id
		LEFT JOIN admin_schema.master_product_category_table mpc ON mp.category_id = mpc.category_id
		LEFT JOIN admin_schema.units_table ut ON oi.unit_id = ut.id
		WHERE oi.order_id = $1`

	rows, err := db.Pool.Query(context.Background(), query, orderID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var p ProductDetail
		if err := rows.Scan(
			&p.ProductID, &p.ProductName, &p.CategoryID, &p.CategoryName,
			&p.Quantity, &p.UnitID, &p.UnitName, &p.MaxPrice,
		); err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	return products, rows.Err()
}

// Helper function to safely dereference float pointers
func derefFloat(val *float64) float64 {
	if val != nil {
		return *val
	}
	return 0.0
}
