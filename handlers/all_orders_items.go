package handlers

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/PragaL15/go_newBackend/go_backend/db"
)

func GetOrderDetails(c *fiber.Ctx) error {
	rows, err := db.Pool.Query(context.Background(), "SELECT * FROM get_order_details()")
	if err != nil {
		log.Printf("Failed to fetch order details: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch order details"})
	}
	defer rows.Close()

	ordersMap := make(map[int]map[string]interface{}) 
	productsMap := make(map[int][]map[string]interface{}) 

	var orders []map[string]interface{} 

	for rows.Next() {
		var orderID, retailerID, wholesellerID, unitID, orderItemID int
		var productID int64
		var dateOfOrder time.Time
		var expectedDeliveryDate, actualDeliveryDate sql.NullTime
		var orderStatus, retailerName, wholesellerName, locationName, stateName, productName, address, pincode string
		var orderItemStatus sql.NullString
		var quantity, totalOrderAmount float64
		var amtOfOrderItem sql.NullFloat64

		err := rows.Scan(
			&orderID, &orderItemID, &dateOfOrder, &orderStatus, &retailerID, &retailerName,
			&wholesellerID, &wholesellerName, &locationName, &stateName, &pincode, &address, &totalOrderAmount,
			&productID, &productName, &quantity, &unitID, &amtOfOrderItem, &expectedDeliveryDate,
			&actualDeliveryDate, &orderItemStatus,
		)
		if err != nil {
			log.Printf("Error scanning row: %v", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error processing data"})
		}

		if _, exists := ordersMap[orderID]; !exists {
			ordersMap[orderID] = map[string]interface{}{
				"order_id":           orderID,
				"date_of_order":      dateOfOrder.Format(time.RFC3339),
				"order_status_name":  orderStatus,
				"retailer_id":        retailerID,
				"retailer_name":      retailerName,
				"wholeseller_id":     wholesellerID,
				"wholeseller_name":   wholesellerName,
				"location_name":      locationName,
				"state_name":         stateName,
				"pincode":            pincode,
				"address":            address,
				"total_order_amount": totalOrderAmount,
			}
			productsMap[orderID] = []map[string]interface{}{}
		}

		if orderItemID != 0 && orderID > 0 {
			product := map[string]interface{}{
				"order_item_id":         orderItemID,
				"product_id":            productID,
				"product_name":          productName,
				"quantity":              quantity,
				"unit_id":               unitID,
				"amt_of_order_item":     formatNullFloat64(amtOfOrderItem),
				"expected_delivery_date": formatNullTime(expectedDeliveryDate),
				"actual_delivery_date":   formatNullTime(actualDeliveryDate),
				"order_item_status_name": formatNullString(orderItemStatus),
			}
			productsMap[orderID] = append(productsMap[orderID], product)
		}
	}

	for orderID, order := range ordersMap {
		order["products"] = productsMap[orderID]
		orders = append(orders, order)
	}

	return c.JSON(orders)
}

func formatNullTime(nt sql.NullTime) string {
	if nt.Valid {
		return nt.Time.Format(time.RFC3339)
	}
	return ""
}

func formatNullString(ns sql.NullString) string {
	if ns.Valid {
		return ns.String
	}
	return ""
}

func formatNullFloat64(nf sql.NullFloat64) float64 {
	if nf.Valid {
		return nf.Float64
	}
	return 0.0
}
