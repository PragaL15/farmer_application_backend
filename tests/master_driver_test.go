package tests

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/PragaL15/go_newBackend/go_backend/db"
	Masterhandlers "github.com/PragaL15/go_newBackend/handlers/master"
	"github.com/gofiber/fiber/v2"
	"github.com/pashagolub/pgxmock"
	"github.com/stretchr/testify/assert"
)

func TestInsertDriver(t *testing.T) {
	mockDB, err := pgxmock.NewPool()
	assert.NoError(t, err)
	db.Pool = mockDB

	app := fiber.New()
	app.Post("/insert-driver", Masterhandlers.InsertDriver)

	requestBody := `{"driver_name":"John Doe","driver_license":"XYZ123","driver_number":"1234567890","license_expiry_date":"2030-01-01"}`
	mockDB.ExpectExec("CALL insert_driver").WillReturnResult(pgxmock.NewResult("INSERT", 1))

	req := httptest.NewRequest(http.MethodPost, "/insert-driver", strings.NewReader(requestBody))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req, -1)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, resp.StatusCode)
}


func TestDeleteDriver(t *testing.T) {
	mockDB, err := pgxmock.NewPool()
	assert.NoError(t, err)
	db.Pool = mockDB

	app := fiber.New()
	app.Delete("/delete-driver/:id", Masterhandlers.DeleteDriver)

	mockDB.ExpectExec("CALL delete_driver").WithArgs(1).WillReturnResult(pgxmock.NewResult("DELETE", 1))

	req := httptest.NewRequest(http.MethodDelete, "/delete-driver/1", nil)
	resp, err := app.Test(req, -1)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

