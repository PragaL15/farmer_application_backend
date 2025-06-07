package handlers

import (
	"context"
	"database/sql"
	"log"
	"time"

	"farmerapp/go_backend/db"

	"github.com/gofiber/fiber/v2"
	"github.com/lib/pq"
)

type FilterOrder struct {
	OrderID               int64           `json:"order_id"`
	DateOfOrder           time.Time       `json:"date_of_order"`
	OrderStatus           sql.NullInt64   `json:"order_status"` // 0 to 6 or NULL
	ActualDeliveryDate    sql.NullTime    `json:"actual_delivery_date"`
	RetailerID            int             `json:"retailer_id"`
	WholesellerID         []int64         `json:"wholeseller_id"`
	TotalOrderAmount      sql.NullFloat64 `json:"total_order_amount"`
	DiscountAmount        sql.NullFloat64 `json:"discount_amount"`
	TaxAmount             sql.NullFloat64 `json:"tax_amount"`
	FinalAmount           sql.NullFloat64 `json:"final_amount"`
	CreatedBy             time.Time       `json:"created_by"`
	UpdatedBy             time.Time       `json:"updated_by"`
	RetailerContactMobile string          `json:"retailer_contact_mobile"`
	DeliveryPincode       string          `json:"delivery_pincode"`
	DeliveryAddress       string          `json:"delivery_address"`
	MaxPriceLimit         sql.NullFloat64 `json:"max_price_limit"`
	DesiredDeliveryDate   time.Time       `json:"desired_delivery_date"`
	DeliveryDeadline      time.Time       `json:"delivery_deadline"`
	WholesellerOfferID    sql.NullInt64   `json:"wholeseller_offer_id"`
	CancellationReason    sql.NullString  `json:"cancellation_reason"`
	Version               int             `json:"version"`
}

func GetFilteredOrders(c *fiber.Ctx) error {
	filter := c.Query("filter", "DATE")

	var query string
	switch filter {
	case "DATE":
		query = `SELECT * FROM business_schema.order_table ORDER BY order_id DESC`
	case "PRICE_LOW":
		query = `SELECT * FROM business_schema.order_table ORDER BY total_order_amount ASC`
	case "PRICE_HIGH":
		query = `SELECT * FROM business_schema.order_table ORDER BY total_order_amount DESC`
	case "BULK":
		query = `SELECT * FROM business_schema.order_table WHERE total_order_amount > 750 ORDER BY total_order_amount DESC`
	case "PRODUCT":
		query = `
			SELECT o.*
			FROM business_schema.order_table o
			JOIN business_schema.order_item_table oi ON o.order_id = oi.order_id
			JOIN admin_schema.master_product p ON oi.product_id = p.product_id
			ORDER BY p.product_name ASC
		`
	default:
		query = `SELECT * FROM business_schema.order_table ORDER BY order_id DESC`
	}

	rows, err := db.Pool.Query(context.Background(), query)
	if err != nil {
		log.Println("Query error:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	defer rows.Close()

	var orders []FilterOrder
	for rows.Next() {
		var o FilterOrder
		var wholesellerRaw sql.NullString

		err := rows.Scan(
			&o.OrderID, &o.DateOfOrder, &o.OrderStatus, &o.ActualDeliveryDate,
			&o.RetailerID, &wholesellerRaw, &o.TotalOrderAmount, &o.DiscountAmount,
			&o.TaxAmount, &o.FinalAmount, &o.CreatedBy, &o.UpdatedBy, &o.RetailerContactMobile,
			&o.DeliveryPincode, &o.DeliveryAddress, &o.MaxPriceLimit, &o.DesiredDeliveryDate,
			&o.DeliveryDeadline, &o.WholesellerOfferID, &o.CancellationReason, &o.Version,
		)
		if err != nil {
			log.Println("Scan error:", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		// Parse wholeseller_id safely
		var wholesellerIDs pq.Int64Array
		if wholesellerRaw.Valid && wholesellerRaw.String != "" {
			err = wholesellerIDs.Scan(wholesellerRaw.String)
			if err != nil {
				log.Println("Array parse error:", err)
				// you can choose to skip or set empty list
			}
		}
		o.WholesellerID = wholesellerIDs

		orders = append(orders, o)
	}

	return c.JSON(orders)
}
