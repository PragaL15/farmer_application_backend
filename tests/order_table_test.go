package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
 handlers "github.com/PragaL15/go_newBackend/handlers"
	"github.com/PragaL15/go_newBackend/go_backend/db"
	"github.com/gofiber/fiber/v2"
	"github.com/pashagolub/pgxmock"
	"github.com/stretchr/testify/assert"
)

func setupApp() *fiber.App {
	app := fiber.New()
	app.Get("/orders", handlers.GetOrders)
	app.Post("/orders", handlers.InsertOrder)
	app.Put("/orders", handlers.UpdateOrder)
	return app
}


func TestGetOrders_Success(t *testing.T) {
	app := setupApp()
	mockDB, err := pgxmock.NewPool()
	assert.NoError(t, err)
	db.Pool = mockDB

	rows := mockDB.NewRows([]string{
		"order_id", "date_of_order", "order_status", "order_status_name",
		"expected_delivery_date", "actual_delivery_date", "retailer_id", "retailer_name",
		"wholeseller_id", "wholeseller_name", "location_id", "location_name",
		"state_id", "state_name", "pincode", "address",
	}).AddRow(
		1, "2024-03-01 10:00:00", 1, "Pending",
		"2024-03-05", "2024-03-06", 101, "Retailer A",
		201, "Wholeseller X", 301, "Location Y",
		401, "State Z", "123456", "Sample Address",
	)
	

	mockDB.ExpectQuery("SELECT \\* FROM sp_get_orders\\(\\)").WillReturnRows(rows)

	req := httptest.NewRequest(http.MethodGet, "/orders", nil)
	resp, err := app.Test(req, -1)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestInsertOrder_Success(t *testing.T) {
	app := setupApp()
	mockDB, err := pgxmock.NewPool()
	assert.NoError(t, err)
	db.Pool = mockDB

	mockDB.ExpectExec("SELECT sp_insert_order").
		WithArgs(1, "2024-03-05", "2024-03-06", 101, 201, 301, 401, "123456", "Sample Address").
		WillReturnResult(pgxmock.NewResult("SELECT", 1))

	requestBody := `{"order_status":1, "expected_delivery_date":"2024-03-05", "actual_delivery_date":"2024-03-06", 
	"retailer_id":101, "wholeseller_id":201, "location_id":301, "state_id":401, "pincode":"123456", "address":"Sample Address"}`

	req := httptest.NewRequest(http.MethodPost, "/orders", strings.NewReader(requestBody))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req, -1)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var responseBody map[string]string
	json.NewDecoder(resp.Body).Decode(&responseBody)
	assert.Equal(t, "Order inserted successfully", responseBody["message"])
}

func TestUpdateOrder_Success(t *testing.T) {
	app := setupApp()
	mockDB, err := pgxmock.NewPool()
	assert.NoError(t, err)
	db.Pool = mockDB

	mockDB.ExpectExec("SELECT sp_update_order").
		WithArgs(1, 2, "2024-03-06").
		WillReturnResult(pgxmock.NewResult("SELECT", 1))

	requestBody := `{"order_id":1, "order_status":2, "actual_delivery_date":"2024-03-06"}`

	req := httptest.NewRequest(http.MethodPut, "/orders", strings.NewReader(requestBody))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req, -1)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var responseBody map[string]string
	json.NewDecoder(resp.Body).Decode(&responseBody)
	assert.Equal(t, "Order updated successfully", responseBody["message"])
}
