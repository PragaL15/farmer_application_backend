package handlers

import (
	"context"

	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/PragaL15/go_newBackend/go_backend/db"
)

func GetInvoiceDetails(c *fiber.Ctx) error {
	rows, err := db.Pool.Query(context.Background(), "SELECT * FROM get_invoice_details_with_business_info();")
	if err != nil {
		log.Printf("Failed to fetch invoice details: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch invoice details"})
	}
	defer rows.Close()

	invoicesMap := make(map[int]map[string]interface{})
	productsMap := make(map[int][]map[string]interface{})
	var invoices []map[string]interface{}

	for rows.Next() {
		var id, orderID, retailerID, retailerLocationID, retailerStateID, wholesellerID, wholesellerLocationID, wholesellerStateID, payMode, payType, orderItemID, productID int
		var totalAmount, discountAmount, finalAmount float64
		var invoiceNumber, retailerName, retailerEmail, retailerPhone, retailerAddress string
		var wholesellerName, wholesellerEmail, wholesellerPhone, wholesellerAddress, payModeName, payTypeName, productName string
		var invoiceDate, dueDate time.Time

		err := rows.Scan(
			&id, &invoiceNumber, &orderID, &retailerID, &retailerName, &retailerEmail, &retailerPhone, &retailerAddress,
			&retailerLocationID, &retailerStateID, &wholesellerID, &wholesellerName, &wholesellerEmail, &wholesellerPhone,
			&wholesellerAddress, &wholesellerLocationID, &wholesellerStateID, &totalAmount, &discountAmount, &invoiceDate,
			&dueDate, &payMode, &payModeName, &payType, &payTypeName, &finalAmount, &orderItemID, &productID, &productName,
		)

		if err != nil {
			log.Printf("Error scanning row: %v", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error processing data"})
		}

		if _, exists := invoicesMap[id]; !exists {
			invoicesMap[id] = map[string]interface{}{
				"id":                   id,
				"invoice_number":       invoiceNumber,
				"order_id":             orderID,
				"retailer_id":          retailerID,
				"retailer_name":        retailerName,
				"retailer_email":       retailerEmail,
				"retailer_phone":       retailerPhone,
				"retailer_address":     retailerAddress,
				"retailer_location_id": retailerLocationID,
				"retailer_state_id":    retailerStateID,
				"wholeseller_id":       wholesellerID,
				"wholeseller_name":     wholesellerName,
				"wholeseller_email":    wholesellerEmail,
				"wholeseller_phone":    wholesellerPhone,
				"wholeseller_address":  wholesellerAddress,
				"wholeseller_location_id": wholesellerLocationID,
				"wholeseller_state_id":    wholesellerStateID,
				"total_amount":         totalAmount,
				"discount_amount":      discountAmount,
				"invoice_date":         invoiceDate.Format(time.RFC3339),
				"due_date":             dueDate.Format("2006-01-02"),
				"pay_mode":             payMode,
				"pay_mode_name":        payModeName,
				"pay_type":             payType,
				"pay_type_name":        payTypeName,
				"final_amount":         finalAmount,
			}
			productsMap[id] = []map[string]interface{}{}
		}

		if orderItemID != 0 {
			product := map[string]interface{}{
				"order_item_id": orderItemID,
				"product_id":    productID,
				"product_name":  productName,
			}
			productsMap[id] = append(productsMap[id], product)
		}
	}

	for id, invoice := range invoicesMap {
		invoice["products"] = productsMap[id]
		invoices = append(invoices, invoice)
	}

	return c.JSON(invoices)
}
