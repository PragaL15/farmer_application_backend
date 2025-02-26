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
		var orderID, orderItemID, retailerID, wholesellerID, unitID int
		var productID int64
		var dateOfOrder time.Time
		var expectedDeliveryDate, actualDeliveryDate sql.NullTime
		var orderStatus, retailerName, wholesellerName, locationName, stateName, productName, unitName, address, pincode sql.NullString
		var totalOrderAmount, quantity float64
		var amtOfOrderItem sql.NullFloat64
		var orderItemStatus sql.NullString
		err := rows.Scan(
			&orderID, &orderItemID, &dateOfOrder, &expectedDeliveryDate, &actualDeliveryDate, 
			&orderStatus, &retailerID, &retailerName, &wholesellerID, &wholesellerName, 
			&locationName, &stateName, &pincode, &address, &totalOrderAmount, 
			&productID, &productName, &quantity, &unitID, &unitName, 
			&amtOfOrderItem, &orderItemStatus,
		)
		if err != nil {
			log.Printf("Error scanning row: %v", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error processing data"})
		}
		if _, exists := ordersMap[orderID]; !exists {
			ordersMap[orderID] = map[string]interface{}{
				"order_id":               orderID,
				"date_of_order":          dateOfOrder.Format(time.RFC3339),
				"expected_delivery_date": formatNullTime(expectedDeliveryDate), 
				"actual_delivery_date":   formatNullTime(actualDeliveryDate), 
				"order_status_name":      formatNullString(orderStatus),
				"retailer_id":            retailerID,
				"retailer_name":          formatNullString(retailerName),
				"wholeseller_id":         wholesellerID,
				"wholeseller_name":       formatNullString(wholesellerName),
				"location_name":          formatNullString(locationName),
				"state_name":             formatNullString(stateName),
				"pincode":                formatNullString(pincode),
				"address":                formatNullString(address),
				"total_order_amount":     totalOrderAmount,
			}
			productsMap[orderID] = []map[string]interface{}{}
		}
		if orderItemID != 0 {
			product := map[string]interface{}{
				"order_item_id":         orderItemID,
				"product_id":            productID,
				"product_name":          formatNullString(productName),
				"quantity":              quantity,
				"unit_id":               unitID,
				"unit_name":             formatNullString(unitName),
				"amt_of_order_item":     formatNullFloat64(amtOfOrderItem),
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


// package handlers
// import (
// 	"context"
// 	"database/sql"
// 	"log"
// 	"time"

// 	"github.com/gofiber/fiber/v2"
// 	"github.com/PragaL15/go_newBackend/go_backend/db"
// )

// type Order struct {
// 	OrderID             int          `json:"order_id"`
// 	DateOfOrder         string       `json:"date_of_order"`
// 	ExpectedDelivery    string       `json:"expected_delivery_date"`
// 	ActualDelivery      string       `json:"actual_delivery_date"`
// 	OrderStatus         string       `json:"order_status_name"`
// 	RetailerID          int          `json:"retailer_id"`
// 	RetailerName        string       `json:"retailer_name"`
// 	WholesellerID       int          `json:"wholeseller_id"`
// 	WholesellerName     string       `json:"wholeseller_name"`
// 	LocationName        string       `json:"location_name"`
// 	StateName           string       `json:"state_name"`
// 	Pincode            string       `json:"pincode"`
// 	Address            string       `json:"address"`
// 	TotalOrderAmount    float64      `json:"total_order_amount"`
// 	Products            []Product    `json:"products"`
// }

// type Product struct {
// 	OrderItemID         int          `json:"order_item_id"`
// 	ProductID           int64        `json:"product_id"`
// 	ProductName         string       `json:"product_name"`
// 	Quantity           float64      `json:"quantity"`
// 	UnitID              int          `json:"unit_id"`
// 	UnitName            string       `json:"unit_name"`
// 	AmtOfOrderItem      float64      `json:"amt_of_order_item"`
// 	OrderItemStatus     string       `json:"order_item_status_name"`
// }

// func GetOrderDetails(c *fiber.Ctx) error {
// 	rows, err := db.Pool.Query(context.Background(), "SELECT * FROM get_order_details()")
// 	if err != nil {
// 		log.Printf("Failed to fetch order details: %v", err)
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch order details"})
// 	}
// 	defer rows.Close()

// 	ordersMap := make(map[int]*Order)

// 	for rows.Next() {
// 		var orderID, orderItemID, retailerID, wholesellerID, unitID int
// 		var productID int64
// 		var dateOfOrder time.Time
// 		var expectedDeliveryDate, actualDeliveryDate sql.NullTime
// 		var orderStatus, retailerName, wholesellerName, locationName, stateName, productName, unitName, address, pincode sql.NullString
// 		var totalOrderAmount, quantity float64
// 		var amtOfOrderItem sql.NullFloat64
// 		var orderItemStatus sql.NullString

// 		err := rows.Scan(
// 			&orderID, &orderItemID, &dateOfOrder, &expectedDeliveryDate, &actualDeliveryDate,
// 			&orderStatus, &retailerID, &retailerName, &wholesellerID, &wholesellerName,
// 			&locationName, &stateName, &pincode, &address, &totalOrderAmount,
// 			&productID, &productName, &quantity, &unitID, &unitName,
// 			&amtOfOrderItem, &orderItemStatus,
// 		)
// 		if err != nil {
// 			log.Printf("Error scanning row: %v", err)
// 			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error processing data"})
// 		}

// 		if _, exists := ordersMap[orderID]; !exists {
// 			ordersMap[orderID] = &Order{
// 				OrderID:            orderID,
// 				DateOfOrder:        dateOfOrder.Format(time.RFC3339),
// 				ExpectedDelivery:   formatNullTime(expectedDeliveryDate),
// 				ActualDelivery:     formatNullTime(actualDeliveryDate),
// 				OrderStatus:        formatNullString(orderStatus),
// 				RetailerID:         retailerID,
// 				RetailerName:       formatNullString(retailerName),
// 				WholesellerID:      wholesellerID,
// 				WholesellerName:    formatNullString(wholesellerName),
// 				LocationName:       formatNullString(locationName),
// 				StateName:          formatNullString(stateName),
// 				Pincode:            formatNullString(pincode),
// 				Address:            formatNullString(address),
// 				TotalOrderAmount:   totalOrderAmount,
// 				Products:           []Product{},
// 			}
// 		}

// 		if orderItemID != 0 {
// 			product := Product{
// 				OrderItemID:        orderItemID,
// 				ProductID:          productID,
// 				ProductName:        formatNullString(productName),
// 				Quantity:           quantity,
// 				UnitID:             unitID,
// 				UnitName:           formatNullString(unitName),
// 				AmtOfOrderItem:     formatNullFloat64(amtOfOrderItem),
// 				OrderItemStatus:    formatNullString(orderItemStatus),
// 			}
// 			ordersMap[orderID].Products = append(ordersMap[orderID].Products, product)
// 		}
// 	}

// 	var orders []Order
// 	for _, order := range ordersMap {
// 		orders = append(orders, *order)
// 	}

// 	return c.JSON(orders)
// }

// func formatNullTime(nt sql.NullTime) string {
// 	if nt.Valid {
// 		return nt.Time.Format(time.RFC3339)
// 	}
// 	return ""
// }

// func formatNullString(ns sql.NullString) string {
// 	if ns.Valid {
// 		return ns.String
// 	}
// 	return ""
// }

// func formatNullFloat64(nf sql.NullFloat64) float64 {
// 	if nf.Valid {
// 		return nf.Float64
// 	}
// 	return 0.0
// }

