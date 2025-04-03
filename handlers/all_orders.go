package handlers

import (
	"context"
	"log"
	"time"
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
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"status":  "error",
					"message": "Order ID is required",
			})
	}

	orderID, err := strconv.ParseInt(orderIDStr, 10, 64)
	if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"status":  "error",
					"message": "Invalid order ID format",
			})
	}

	// Define our response structure
	type ProductDetail struct {
			ProductID    int64   `json:"product_id"`
			ProductName string  `json:"product_name"`
			CategoryID   int     `json:"category_id"`
			CategoryName string  `json:"category_name"`
			Quantity     float64 `json:"quantity"`
			UnitID       int     `json:"unit_id"`
			UnitName     string  `json:"unit_name"`
			MaxPrice     float64 `json:"max_price"`
	}

	type OrderResponse struct {
			OrderID             int64          `json:"order_id"`
			DateOfOrder        time.Time      `json:"date_of_order"`
			OrderStatus        int            `json:"order_status"`
			ActualDeliveryDate *time.Time     `json:"actual_delivery_date,omitempty"`
			RetailerID         int            `json:"retailer_id"`
			ShopName           string         `json:"shop_name"`
			WholesellerIDs     []int          `json:"wholeseller_ids,omitempty"`
			TotalOrderAmount   float64        `json:"total_order_amount"`
			DiscountAmount     float64        `json:"discount_amount"`
			TaxAmount          float64        `json:"tax_amount"`
			FinalAmount        float64        `json:"final_amount"`
			Products          []ProductDetail `json:"products"`
	}

	// First get the basic order information
	var order struct {
			OrderID             int64
			DateOfOrder        time.Time
			OrderStatus        int
			ActualDeliveryDate *time.Time
			RetailerID         int
			TotalOrderAmount   float64
			DiscountAmount     float64
			TaxAmount          float64
			FinalAmount        float64
	    
	}

	orderQuery := `
			SELECT 
					order_id, date_of_order, order_status, 
					actual_delivery_date,
					retailer_id, total_order_amount,
					discount_amount, tax_amount, final_amount
			FROM business_schema.order_table
			WHERE order_id = $1`

	err = db.Pool.QueryRow(context.Background(), orderQuery, orderID).Scan(
			&order.OrderID, &order.DateOfOrder, &order.OrderStatus,
			&order.ActualDeliveryDate,
			&order.RetailerID, &order.TotalOrderAmount,
			&order.DiscountAmount, &order.TaxAmount, &order.FinalAmount,
	)

	if err == pgx.ErrNoRows {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
					"status":  "error",
					"message": "Order not found",
			})
	} else if err != nil {
			log.Printf("Failed to fetch order: %v", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"status":  "error",
					"message": "Failed to fetch order details",
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

	var wholesellerIDs []int
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
							wholesellerIDs = append(wholesellerIDs, id)
					}
			}
			if err = rows.Err(); err != nil {
					log.Printf("Error processing wholeseller rows: %v", err)
			}
	}

	var products []ProductDetail
	productQuery := `
			SELECT 
					oi.product_id, 
					mp.product_name,
					mp.category_id,
					mpc.category_name,
					oi.quantity, 
					oi.unit_id,
					ut.unit_name,
					oi.max_item_price
			FROM business_schema.order_item_table oi
			JOIN admin_schema.master_product mp ON oi.product_id = mp.product_id
			LEFT JOIN admin_schema.master_product_category_table mpc ON mp.category_id = mpc.category_id
			LEFT JOIN admin_schema.units_table ut ON oi.unit_id = ut.id
			WHERE oi.order_id = $1`

	productRows, err := db.Pool.Query(context.Background(), productQuery, orderID)
	if err != nil {
			log.Printf("Failed to fetch order items: %v", err)
	} else {
			defer productRows.Close()
			for productRows.Next() {
					var p ProductDetail
					if err := productRows.Scan(
							&p.ProductID, &p.ProductName, &p.CategoryID, &p.CategoryName,
							&p.Quantity, &p.UnitID, &p.UnitName, &p.MaxPrice,
					); err == nil {
							products = append(products, p)
					}
			}
			if err = productRows.Err(); err != nil {
					log.Printf("Error processing product rows: %v", err)
			}
	}

	response := OrderResponse{
			OrderID:             order.OrderID,
			DateOfOrder:        order.DateOfOrder,
			OrderStatus:        order.OrderStatus,
			ActualDeliveryDate: order.ActualDeliveryDate,
			RetailerID:         order.RetailerID,
			ShopName:           shopName,
			WholesellerIDs:     wholesellerIDs,
			TotalOrderAmount:   order.TotalOrderAmount,
			DiscountAmount:     order.DiscountAmount,
			TaxAmount:          order.TaxAmount,
			FinalAmount:        order.FinalAmount,
			Products:          products,
	}

	return c.JSON(response)
}