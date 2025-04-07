package handlers

import (
	"context"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/PragaL15/go_newBackend/go_backend/db"
)

func GetOrderHistory(c *fiber.Ctx) error {
	rows, err := db.Pool.Query(context.Background(), "SELECT * FROM business_schema.get_order_history_with_details()")
	if err != nil {
		log.Printf("Failed to fetch order history: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch order history"})
	}
	defer rows.Close()

	var orders []map[string]interface{}

	for rows.Next() {
		var (
			historyID, orderID, orderStatus, retailerID, wholesellerID, deliveryLocationID, deliveryStateID, userID int
			retailerName, wholesellerName, locationName, cityName, stateName                                     *string
			pincode, deliveryAddress, retailerContactMobile, cancellationReason, command                         *string
			dateOfOrder, expectedDeliveryDate, actualDeliveryDate, deliveryCompletedDate                         *time.Time
			totalOrderAmount, discountAmount, taxAmount, finalAmount, maxPriceLimit                              *float64
			deliveryDeadline, desiredDeliveryDate                                                                *time.Time
			wholesellerOfferID                                                                                   *int64
			createdBy, updatedBy                                                                                 *time.Time
		)

		if err := rows.Scan(
			&historyID,
			&orderID,
			&dateOfOrder,
			&orderStatus,
			&expectedDeliveryDate,
			&actualDeliveryDate,
			&deliveryCompletedDate,
			&retailerID,
			&retailerName,
			&wholesellerID,
			&wholesellerName,
			&deliveryLocationID,
			&locationName,
			&cityName,
			&deliveryStateID,
			&stateName,
			&pincode,
			&deliveryAddress,
			&retailerContactMobile,
			&totalOrderAmount,
			&discountAmount,
			&taxAmount,
			&finalAmount,
			&deliveryDeadline,
			&desiredDeliveryDate,
			&wholesellerOfferID,
			&cancellationReason,
			&maxPriceLimit,
			&userID,
			&command,
			&createdBy,
			&updatedBy,
		); err != nil {
			log.Printf("Error scanning row: %v", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error processing data"})
		}

		order := map[string]interface{}{
			"history_id":               historyID,
			"order_id":                 orderID,
			"date_of_order":            dateOfOrder,
			"order_status":             orderStatus,
			"expected_delivery_date":   expectedDeliveryDate,
			"actual_delivery_date":     actualDeliveryDate,
			"delivery_completed_date":  deliveryCompletedDate,
			"retailer_id":              retailerID,
			"retailer_name":            nullSafeString(retailerName),
			"wholeseller_id":           wholesellerID,
			"wholeseller_name":         nullSafeString(wholesellerName),
			"delivery_location_id":     deliveryLocationID,
			"location_name":            nullSafeString(locationName),
			"city_name":                nullSafeString(cityName),
			"delivery_state_id":        deliveryStateID,
			"state_name":               nullSafeString(stateName),
			"pincode":                  nullSafeString(pincode),
			"delivery_address":         nullSafeString(deliveryAddress),
			"retailer_contact_mobile":  nullSafeString(retailerContactMobile),
			"total_order_amount":       totalOrderAmount,
			"discount_amount":          discountAmount,
			"tax_amount":               taxAmount,
			"final_amount":             finalAmount,
			"delivery_deadline":        deliveryDeadline,
			"desired_delivery_date":    desiredDeliveryDate,
			"wholeseller_offer_id":     wholesellerOfferID,
			"cancellation_reason":      nullSafeString(cancellationReason),
			"max_price_limit":          maxPriceLimit,
			"user_id":                  userID,
			"command":                  nullSafeString(command),
			"created_by":               createdBy,
			"updated_by":               updatedBy,
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
