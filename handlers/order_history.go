package handlers

import (
	"context"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/PragaL15/go_newBackend/go_backend/db"
)

func GetOrderHistory(c *fiber.Ctx) error {
	rows, err := db.Pool.Query(context.Background(), "SELECT * FROM get_order_history()")
	if err != nil {
		log.Printf("Failed to fetch order history: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch order history"})
	}
	defer rows.Close()
	var orders []map[string]interface{}

	for rows.Next() {
		var (
			orderID, orderStatus, retailerID, wholesellerID, locationID, stateID, historyID int
			retailerName, wholesellerName, locationName, stateName                          *string
			dateOfOrder, expectedDeliveryDate, actualDeliveryDate, deliveryCompletedDate   *time.Time
			pincode, address                                                              *string
		)
		if err := rows.Scan(
			&orderID, &dateOfOrder, &orderStatus, &expectedDeliveryDate, &actualDeliveryDate,
			&retailerID, &retailerName, &wholesellerID, &wholesellerName, &locationID, &locationName,
			&stateID, &stateName, &pincode, &address, &deliveryCompletedDate, &historyID,
		); err != nil {
			log.Printf("Error scanning row: %v", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error processing data"})
		}

		order := map[string]interface{}{
			"order_id":                orderID,
			"date_of_order":           dateOfOrder,
			"order_status":            orderStatus,
			"expected_delivery_date":  expectedDeliveryDate,
			"actual_delivery_date":    actualDeliveryDate,
			"retailer_id":             retailerID,
			"retailer_name":           nullSafeString(retailerName),
			"wholeseller_id":          wholesellerID,
			"wholeseller_name":        nullSafeString(wholesellerName),
			"location_id":             locationID,
			"location_name":           nullSafeString(locationName),
			"state_id":                stateID,
			"state_name":              nullSafeString(stateName),
			"pincode":                 nullSafeString(pincode),
			"address":                 nullSafeString(address),
			"delivery_completed_date": deliveryCompletedDate,
			"history_id":              historyID,
		}
		orders = append(orders, order)
	}
	return c.JSON(orders)
}
func nullSafeString(s *string) string {
	if s == nil {
		return "Unknown"
	}
	return *s
}
