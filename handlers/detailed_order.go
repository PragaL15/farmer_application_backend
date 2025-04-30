package handlers

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/PragaL15/go_newBackend/go_backend/db"
)

// Structs for DB Scan (use sql.Null*)
type OrderItems struct {
	OrderItemID  sql.NullInt64
	ProductID    sql.NullInt64
	ProductName  sql.NullString
	Quantity     sql.NullFloat64
	UnitID       sql.NullInt64
	MaxItemPrice sql.NullFloat64
	UnitName     sql.NullString
}

type OrderDetail struct {
	OrderID           sql.NullInt64
	RetailerID        sql.NullInt64
	RetailerName      sql.NullString
	RetailerAddress   sql.NullString
	RetailerMobile    sql.NullString
	ActualDeliveryDate sql.NullTime
	OrderStatusID     sql.NullInt64
	OrderStatus       sql.NullString
	TotalOrderAmount  sql.NullFloat64
	OrderItems        []OrderItems
}

// Structs for clean JSON response
type CleanOrderItem struct {
	OrderItemID  int64   `json:"order_item_id"`
	ProductID    int64   `json:"product_id"`
	ProductName  string  `json:"product_name"`
	Quantity     float64 `json:"quantity"`
	UnitID       int64   `json:"unit_id"`
	MaxItemPrice float64 `json:"max_item_price"`
	UnitName     string  `json:"unit_name"`
}

type CleanOrderDetail struct {
	OrderID            int64            `json:"order_id"`
	RetailerID         int64            `json:"retailer_id"`
	RetailerName       string           `json:"retailer_name"`
	RetailerAddress    string           `json:"retailer_address"`
	RetailerMobile     string           `json:"retailer_mobile"`
	ActualDeliveryDate string           `json:"actual_delivery_date"`
	OrderStatusID      int64            `json:"order_status_id"`
	OrderStatus        string           `json:"order_status"`
	TotalOrderAmount   float64          `json:"total_order_amount"`
	OrderItems         []CleanOrderItem `json:"order_items"`
}

// Handler Function
func GetAllOrderItemDetailsHandler(c *fiber.Ctx) error {
	query := `
		SELECT 
			order_id,
			retailer_id,
			retailer_name,
			retailer_address,
			retailer_mobile,
			actual_delivery_date,
			order_status_id,
			order_status,
			total_order_amount,
			order_item_id,
			product_id,
			product_name,
			quantity,
			unit_id,
			max_item_price,
			unit_name
		FROM 
			business_schema.get_order_details_v2();
	`

	rows, err := db.Pool.Query(context.Background(), query)
	if err != nil {
		log.Println("Error executing query:", err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	defer rows.Close()

	orderMap := make(map[int64]*CleanOrderDetail)

	for rows.Next() {
		var od OrderDetail
		var oi OrderItems

		err := rows.Scan(
			&od.OrderID,
			&od.RetailerID,
			&od.RetailerName,
			&od.RetailerAddress,
			&od.RetailerMobile,
			&od.ActualDeliveryDate,
			&od.OrderStatusID,
			&od.OrderStatus,
			&od.TotalOrderAmount,
			&oi.OrderItemID,
			&oi.ProductID,
			&oi.ProductName,
			&oi.Quantity,
			&oi.UnitID,
			&oi.MaxItemPrice,
			&oi.UnitName,
		)

		if err != nil {
			log.Println("Error scanning row:", err)
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		orderID := od.OrderID.Int64

		// Convert order item
		cleanItem := CleanOrderItem{
			OrderItemID:  oi.OrderItemID.Int64,
			ProductID:    oi.ProductID.Int64,
			ProductName:  oi.ProductName.String,
			Quantity:     oi.Quantity.Float64,
			UnitID:       oi.UnitID.Int64,
			MaxItemPrice: oi.MaxItemPrice.Float64,
			UnitName:     oi.UnitName.String,
		}

		if existingOrder, ok := orderMap[orderID]; ok {
			existingOrder.OrderItems = append(existingOrder.OrderItems, cleanItem)
		} else {
			// Convert time to string if valid
			var deliveryDateStr string
			if od.ActualDeliveryDate.Valid {
				deliveryDateStr = od.ActualDeliveryDate.Time.Format(time.RFC3339)
			}

			orderMap[orderID] = &CleanOrderDetail{
				OrderID:            orderID,
				RetailerID:         od.RetailerID.Int64,
				RetailerName:       od.RetailerName.String,
				RetailerAddress:    od.RetailerAddress.String,
				RetailerMobile:     od.RetailerMobile.String,
				ActualDeliveryDate: deliveryDateStr,
				OrderStatusID:      od.OrderStatusID.Int64,
				OrderStatus:        od.OrderStatus.String,
				TotalOrderAmount:   od.TotalOrderAmount.Float64,
				OrderItems:         []CleanOrderItem{cleanItem},
			}
		}
	}

	var cleanOrders []CleanOrderDetail
	for _, v := range orderMap {
		cleanOrders = append(cleanOrders, *v)
	}

	return c.Status(http.StatusOK).JSON(cleanOrders)
}
