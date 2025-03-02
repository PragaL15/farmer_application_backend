package tests

import (
	
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
 "github.com/PragaL15/go_newBackend/go_backend/db"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"


	Masterhandlers "github.com/PragaL15/go_newBackend/handlers/master"
	"github.com/pashagolub/pgxmock"
)

type OrderStatus struct {
	OrderID     int    `json:"order_id"`
	OrderStatus string `json:"order_status"`
}

func TestGetOrderStatuses(t *testing.T) {
	app := fiber.New()
	app.Get("/order-statuses", Masterhandlers.GetOrderStatuses)

	// ✅ Create a new pgxmock pool
	mockDB, err := pgxmock.NewPool()
	assert.NoError(t, err)

	// ✅ Assign mockDB to db.Pool
	db.Pool = &db.PgxDB{Pool: mockDB}

	// ✅ Mock database query response
	rows := mockDB.NewRows([]string{"order_id", "order_status"}).
		AddRow(1, "Processing").
		AddRow(2, "Confirmed").
		AddRow(3, "Payment").
		AddRow(4, "Out for Delivery").
		AddRow(5, "Successful").
		AddRow(6, "Cancellation").
		AddRow(7, "Returned")

	mockDB.ExpectQuery(`SELECT \* FROM sp_get_order_status\(\)`).WillReturnRows(rows)

	// ✅ Create a test request
	req := httptest.NewRequest(http.MethodGet, "/order-statuses", nil)
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// ✅ Decode response
	var result []OrderStatus
	err = json.NewDecoder(resp.Body).Decode(&result)
	assert.NoError(t, err)
	assert.NotEmpty(t, result)
	assert.Len(t, result, 7)

	// ✅ Check if returned order statuses match expected values
	expectedStatuses := []string{"Processing", "Confirmed", "Payment", "Out for Delivery", "Successful", "Cancellation", "Returned"}
	for i, status := range expectedStatuses {
		assert.Equal(t, status, result[i].OrderStatus)
	}

	// ✅ Ensure all expectations were met
	assert.NoError(t, mockDB.ExpectationsWereMet())
}
