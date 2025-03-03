package tests

import (
	"encoding/json"
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

func setupApp() *fiber.App {
	app := fiber.New()
	app.Post("/insert", Masterhandlers.InsertMasterProduct)
	app.Put("/update", Masterhandlers.UpdateMasterProduct)
	app.Delete("/delete/:id", Masterhandlers.DeleteMasterProduct)
	app.Get("/products", Masterhandlers.GetProducts)
	return app
}

func TestInsertMasterProduct_Success(t *testing.T) {
	app := setupApp()
	mockDB, err := pgxmock.NewPool()
	assert.NoError(t, err)
	db.Pool = mockDB

	mockDB.ExpectExec(`(?i)CALL insert_master_product\(\$1, \$2, \$3\)`).
		WithArgs(1, "Product A", 1).
		WillReturnResult(pgxmock.NewResult("CALL", 1))

	requestBody := `{"category_id":1, "product_name":"Product A", "status":1}`
	req := httptest.NewRequest(http.MethodPost, "/insert", strings.NewReader(requestBody))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req, -1)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, resp.StatusCode)
}



func TestUpdateMasterProduct_Success(t *testing.T) {
	app := setupApp()
	mockDB, err := pgxmock.NewPool()
	assert.NoError(t, err)
	db.Pool = mockDB
	mockDB.ExpectExec(`(?i)CALL update_master_product\(\$1, \$2, \$3, \$4\)`).
    WithArgs(int32(1), int32(1), "Updated Product", int32(1)). 
    WillReturnResult(pgxmock.NewResult("CALL", 1))

	requestBody := `{"product_id":1, "category_id":1, "product_name":"Updated Product", "status":1}`
	req := httptest.NewRequest(http.MethodPut, "/update", strings.NewReader(requestBody))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req, -1)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}



func TestDeleteMasterProduct_Success(t *testing.T) {
	app := setupApp()
	mockDB, err := pgxmock.NewPool()
	assert.NoError(t, err)
	db.Pool = mockDB

	mockDB.ExpectExec(`(?i)CALL delete_master_product\(\$1\)`).
		WithArgs(1).
		WillReturnResult(pgxmock.NewResult("CALL", 1))

	req := httptest.NewRequest(http.MethodDelete, "/delete/1", nil)
	resp, err := app.Test(req, -1)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}


func TestGetProducts_Success(t *testing.T) {
	app := setupApp()
	mockDB, err := pgxmock.NewPool()
	assert.NoError(t, err)
	db.Pool = mockDB

	prodID, catID, status := 1, 1, 1
	catName, prodName := "Category A", "Product A"

	rows := mockDB.NewRows([]string{"product_id", "category_id", "category_name", "product_name", "status"}).
			AddRow(&prodID, &catID, &catName, &prodName, &status) 

	mockDB.ExpectQuery(`SELECT \* FROM get_master_products\(\)`).WillReturnRows(rows)

	req := httptest.NewRequest(http.MethodGet, "/products", nil)
	resp, err := app.Test(req, -1)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var products []map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&products)
	assert.NoError(t, err)

	assert.Len(t, products, 1)
	assert.Equal(t, float64(1), products[0]["product_id"])
	assert.Equal(t, float64(1), products[0]["category_id"])
	assert.Equal(t, "Category A", products[0]["category_name"])
	assert.Equal(t, "Product A", products[0]["product_name"])
	assert.Equal(t, float64(1), products[0]["status"])
}

