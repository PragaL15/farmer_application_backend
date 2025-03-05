package tests

// import (
// 	"bytes"
// 	"encoding/json"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/gofiber/fiber/v2"
// 	"github.com/pashagolub/pgxmock"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/PragaL15/go_newBackend/go_backend/db"
// 		Masterhandlers "github.com/PragaL15/go_newBackend/handlers/master"
// )

// func setupMockDB(t *testing.T) pgxmock.PgxPoolIface {
// 	mockPool, err := pgxmock.NewPool()
// 	if err != nil {
// 		t.Fatalf("Failed to create mock database: %v", err)
// 	}
// 	db.Pool = mockPool
// 	return mockPool
// }

// func TestInsertMasterMandi(t *testing.T) {
// 	app := fiber.New()
// 	app.Post("/mandi", Masterhandlers.InsertMasterMandi)
// 	mockPool := setupMockDB(t)
// 	defer mockPool.Close()

// 	mockPool.ExpectExec(`CALL insert_master_mandi\(\$1, \$2, \$3, \$4, \$5, \$6, \$7, \$8, \$9\)`).
// 		WithArgs("Delhi", "12345", "John Doe", "9876543210", "110001", "Address", "Remarks", 1, 1).
// 		WillReturnResult(pgxmock.NewResult("CALL", 1))

// 	payload := map[string]interface{}{
// 		"mandi_location":     "Delhi",
// 		"mandi_number":       "12345",
// 		"mandi_incharge":     "John Doe",
// 		"mandi_incharge_num": "9876543210",
// 		"mandi_pincode":      "110001",
// 		"mandi_address":      "Address",
// 		"remarks":            "Remarks",
// 		"mandi_city":         1,
// 		"mandi_state":        1,
// 	}
// 	body, _ := json.Marshal(payload)
// 	req := httptest.NewRequest(http.MethodPost, "/mandi", bytes.NewBuffer(body))
// 	req.Header.Set("Content-Type", "application/json")
// 	resp, _ := app.Test(req)

// 	assert.Equal(t, http.StatusCreated, resp.StatusCode)
// }

// func TestUpdateMasterMandi(t *testing.T) {
// 	app := fiber.New()
// 	app.Put("/mandi", Masterhandlers.UpdateMasterMandi)
// 	mockPool := setupMockDB(t)
// 	defer mockPool.Close()

// 	mockPool.ExpectExec(`CALL update_master_mandi\(\$1, \$2, \$3, \$4, \$5, \$6, \$7, \$8, \$9, \$10\)`).
// 		WithArgs(1, "Updated Location", "54321", "Jane Doe", "9876543210", "110002", "New Address", "New Remarks", 2, 2).
// 		WillReturnResult(pgxmock.NewResult("CALL", 1))

// 	payload := map[string]interface{}{
// 		"id":                1,
// 		"mandi_location":     "Updated Location",
// 		"mandi_number":       "54321",
// 		"mandi_incharge":     "Jane Doe",
// 		"mandi_incharge_num": "9876543210",
// 		"mandi_pincode":      "110002",
// 		"mandi_address":      "New Address",
// 		"remarks":            "New Remarks",
// 		"mandi_city":         2,
// 		"mandi_state":        2,
// 	}
// 	body, _ := json.Marshal(payload)
// 	req := httptest.NewRequest(http.MethodPut, "/mandi", bytes.NewBuffer(body))
// 	req.Header.Set("Content-Type", "application/json")
// 	resp, _ := app.Test(req)

// 	assert.Equal(t, http.StatusOK, resp.StatusCode)
// }

// func TestDeleteMasterMandi(t *testing.T) {
// 	app := fiber.New()
// 	app.Delete("/mandi/:id", Masterhandlers.DeleteMasterMandi)
// 	mockPool := setupMockDB(t)
// 	defer mockPool.Close()

// 	mockPool.ExpectExec(`CALL delete_master_mandi\(\$1\)`).
// 		WithArgs(1).
// 		WillReturnResult(pgxmock.NewResult("CALL", 1))

// 	req := httptest.NewRequest(http.MethodDelete, "/mandi/1", nil)
// 	resp, _ := app.Test(req)

// 	assert.Equal(t, http.StatusOK, resp.StatusCode)
// }

// func TestGetMandi(t *testing.T) {
// 	app := fiber.New()
// 	app.Get("/mandi", Masterhandlers.GetMandi)
// 	mockPool := setupMockDB(t)
// 	defer mockPool.Close()

// 	columns := []string{"mandi_id", "mandi_location", "mandi_number", "mandi_incharge", "mandi_incharge_num", "mandi_pincode", "mandi_address", "remarks", "mandi_city", "mandi_state"}
// 	mockPool.ExpectQuery(`SELECT \* FROM get_master_mandis\(\)`).
// 		WillReturnRows(pgxmock.NewRows(columns).
// 			AddRow(1, "Delhi", "12345", "John Doe", "9876543210", "110001", "Address", "Remarks", 1, 1))

// 	req := httptest.NewRequest(http.MethodGet, "/mandi", nil)
// 	resp, _ := app.Test(req)

// 	assert.Equal(t, http.StatusOK, resp.StatusCode)
// }