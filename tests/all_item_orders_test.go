
package tests 
// import (
// 	"database/sql" 
// 	"encoding/json"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"
// 	"time"

// 	"github.com/PragaL15/go_newBackend/go_backend/db"
// 	"github.com/PragaL15/go_newBackend/handlers"
// 	"github.com/gofiber/fiber/v2"
// 	"github.com/pashagolub/pgxmock"
// 	"github.com/stretchr/testify/assert"
// )

// func TestGetOrderDetails_Success(t *testing.T) {
	
// 	app := fiber.New()
// 	app.Get("/orders", handlers.GetOrderDetails)

// 	mockDB, err := pgxmock.NewPool()
// 	assert.NoError(t, err, "Failed to create mock DB")
// 	db.Pool = mockDB 
// 	now := time.Now()
// 	expectedDelivery := sql.NullTime{Time: now.AddDate(0, 0, 7), Valid: true} 
// 	actualDelivery := sql.NullTime{Valid: false}                              

// 	expectedRows := mockDB.NewRows([]string{
//     "order_id", "order_item_id", "date_of_order", "expected_delivery_date", "actual_delivery_date",
//     "order_status_name", "retailer_id", "retailer_name", "wholeseller_id", "wholeseller_name",
//     "location_name", "state_name", "pincode", "address", "total_order_amount",
//     "product_id", "product_name", "quantity", "unit_id", "unit_name",
//     "amt_of_order_item", "order_item_status_name",
// }).AddRow(
//     1, 1, 
//     now,
//     expectedDelivery,
//     actualDelivery,
//     "Delivered",
//     2, "Retailer A", 2, "Wholeseller B",
//     "City X", "State Y", "123456", "Street 1",
//     100.50,
//     int64(10), "Product A", 5.0, 1, "Kg",
//     50.25, "Completed",
// )


// 	mockDB.ExpectQuery(`(?i)^SELECT \* FROM get_order_details\(\)$`).WillReturnRows(expectedRows)

// 	// Perform HTTP request
// 	req := httptest.NewRequest(http.MethodGet, "/orders", nil)
// 	req.Header.Set("Content-Type", "application/json")
// 	resp, err := app.Test(req, -1)

// 	// Validate response
// 	assert.NoError(t, err, "Request to API failed")
// 	assert.Equal(t, http.StatusOK, resp.StatusCode, "Expected HTTP 200 OK")

// 	// Decode response body
// 	var orders []map[string]interface{}
// 	err = json.NewDecoder(resp.Body).Decode(&orders)
// 	assert.NoError(t, err, "Failed to decode response body")

// 	// Validate response data
// 	assert.Len(t, orders, 1, "Expected 1 order in response")
// 	assert.Equal(t, float64(1), orders[0]["order_id"], "Order ID mismatch")
// 	assert.Equal(t, "Delivered", orders[0]["order_status_name"], "Order status mismatch")

// 	// Ensure mock expectations are met
// 	assert.NoError(t, mockDB.ExpectationsWereMet(), "Mock expectations were not met")
// }
