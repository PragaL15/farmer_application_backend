package tests

import (
	
	"net/http"
	"net/http/httptest"
	"testing"
"strings"
	"github.com/PragaL15/go_newBackend/go_backend/db"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"

	Masterhandlers "github.com/PragaL15/go_newBackend/handlers/master"
	"github.com/pashagolub/pgxmock"
)

func TestInsertBusiness_Success(t *testing.T) {
	app := fiber.New()
	app.Post("/business", Masterhandlers.InsertBusiness)

	mockDB, err := pgxmock.NewPool()
	assert.NoError(t, err, "Failed to create mock DB")
	db.Pool = &db.MockDB{Mock: mockDB}

	mockDB.ExpectExec(`SELECT insert_business\(.*,.*,.*,.*,.*,.*,.*,.*,.*,.*\)`).WillReturnResult(pgxmock.NewResult("INSERT", 1))

	body := `{"b_typeid":1,"b_name":"Test Business","b_location_id":2,"b_state_id":3,"b_address":"Test Address","b_email":"test@example.com","b_gstnum":"12345678901234567890","b_pannum":"ABCDE1234F"}`
	req := httptest.NewRequest(http.MethodPost, "/business", strings.NewReader(body))

	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)
	assert.NoError(t, err, "Request to API failed")
	assert.Equal(t, http.StatusCreated, resp.StatusCode, "Expected HTTP 201 Created")

	assert.NoError(t, mockDB.ExpectationsWereMet(), "Mock expectations were not met")
}

func TestUpdateBusiness_Success(t *testing.T) {
	app := fiber.New()
	app.Put("/business", Masterhandlers.UpdateBusiness)

	mockDB, err := pgxmock.NewPool()
	assert.NoError(t, err, "Failed to create mock DB")
	db.Pool = &db.MockDB{Mock: mockDB}

	mockDB.ExpectExec(`CALL update_business\(.*,.*,.*,.*,.*,.*,.*,.*,.*,.*\)`).WillReturnResult(pgxmock.NewResult("UPDATE", 1))

	body := `{"bid":1,"b_typeid":2,"b_location_id":3,"b_state_id":4,"b_address":"Updated Address","b_email":"updated@example.com","b_gstnum":"98765432109876543210","b_pan":"XYZAB1234Z"}`
	req := httptest.NewRequest(http.MethodPut, "/business", strings.NewReader(body))

	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)
	assert.NoError(t, err, "Request to API failed")
	assert.Equal(t, http.StatusOK, resp.StatusCode, "Expected HTTP 200 OK")

	assert.NoError(t, mockDB.ExpectationsWereMet(), "Mock expectations were not met")
}

func TestDeleteBusiness_Success(t *testing.T) {
	app := fiber.New()
	app.Delete("/business/:id", Masterhandlers.DeleteBusiness)

	mockDB, err := pgxmock.NewPool()
	assert.NoError(t, err, "Failed to create mock DB")
	db.Pool = &db.MockDB{Mock: mockDB}

	mockDB.ExpectExec(`CALL delete_business\(.*\)`).WillReturnResult(pgxmock.NewResult("DELETE", 1))

	req := httptest.NewRequest(http.MethodDelete, "/business/1", nil)
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)
	assert.NoError(t, err, "Request to API failed")
	assert.Equal(t, http.StatusOK, resp.StatusCode, "Expected HTTP 200 OK")

	assert.NoError(t, mockDB.ExpectationsWereMet(), "Mock expectations were not met")
}

func TestGetBusinesses_Success(t *testing.T) {
	app := fiber.New()
	app.Get("/businesses", Masterhandlers.GetBusinesses)

	mockDB, err := pgxmock.NewPool()
	assert.NoError(t, err, "Failed to create mock DB")
	db.Pool = &db.MockDB{Mock: mockDB}

	rows := mockDB.NewRows([]string{"bid", "b_type_id", "b_typename", "b_name", "b_location_id", "location", "b_state_id", "state_name", "b_mandiid", "b_address", "b_phone_num", "b_email", "b_gstnum", "b_pan"}).
		AddRow(1, 1, "Retail", "Business 1", 2, "City A", 3, "State A", 4, "Address 1", "1234567890", "email1@example.com", "GST123", "PAN123").
		AddRow(2, 2, "Wholesale", "Business 2", 3, "City B", 4, "State B", 5, "Address 2", "0987654321", "email2@example.com", "GST456", "PAN456")

	mockDB.ExpectQuery(`SELECT \* FROM get_all_businesses_func\(\)`).WillReturnRows(rows)

	req := httptest.NewRequest(http.MethodGet, "/businesses", nil)
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)
	assert.NoError(t, err, "Request to API failed")
	assert.Equal(t, http.StatusOK, resp.StatusCode, "Expected HTTP 200 OK")

	assert.NoError(t, mockDB.ExpectationsWereMet(), "Mock expectations were not met")
}
