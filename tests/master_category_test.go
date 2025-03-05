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
// 	Masterhandlers "github.com/PragaL15/go_newBackend/handlers/master"
// )

// func setupMockDB(t *testing.T) pgxmock.PgxPoolIface {
// 	mockPool, err := pgxmock.NewPool()
// 	if err != nil {
// 		t.Fatalf("Failed to create mock pool: %v", err)
// 	}
// 	db.Pool = mockPool
// 	return mockPool
// }

// func TestInsertCategory(t *testing.T) {
// 	app := fiber.New()
// 	app.Post("/category", Masterhandlers.InsertCategory)
// 	mockPool := setupMockDB(t)
// 	defer mockPool.Close()

// 	mockPool.ExpectExec(`CALL insert_category\(\$1, \$2\)`).
// 		WithArgs("Electronics", pgxmock.AnyArg()). 
// 		WillReturnResult(pgxmock.NewResult("CALL", 1))

// 	payload := map[string]interface{}{
// 		"category_name": "Electronics",
// 		"super_cat_id":  nil, 
// 	}
// 	body, _ := json.Marshal(payload)
// 	req := httptest.NewRequest(http.MethodPost, "/category", bytes.NewBuffer(body))
// 	req.Header.Set("Content-Type", "application/json")
// 	resp, _ := app.Test(req)
// 	assert.Equal(t, http.StatusCreated, resp.StatusCode)
// }

// func TestUpdateCategory(t *testing.T) {
// 	app := fiber.New()
// 	app.Put("/category", Masterhandlers.UpdateCategory)
// 	mockPool := setupMockDB(t)
// 	defer mockPool.Close()

// 	mockPool.ExpectExec(`CALL update_category\(\$1, \$2, \$3\)`).
// 		WithArgs(1, "Updated Category", pgxmock.AnyArg()). 
// 		WillReturnResult(pgxmock.NewResult("CALL", 1))

// 	payload := map[string]interface{}{
// 		"category_id":   1,
// 		"category_name": "Updated Category",
// 		"super_cat_id":  nil, 
// 	}
// 	body, _ := json.Marshal(payload)
// 	req := httptest.NewRequest(http.MethodPut, "/category", bytes.NewBuffer(body))
// 	req.Header.Set("Content-Type", "application/json")
// 	resp, _ := app.Test(req)
// 	assert.Equal(t, http.StatusOK, resp.StatusCode)
// }


// func TestDeleteCategory(t *testing.T) {
// 	app := fiber.New()
// 	app.Delete("/category/:id", Masterhandlers.DeleteCategory)
// 	mockPool := setupMockDB(t)
// 	defer mockPool.Close()

// 	mockPool.ExpectExec(`CALL delete_category\(\$1\)`).
// 		WithArgs(1).
// 		WillReturnResult(pgxmock.NewResult("CALL", 1))

// 	req := httptest.NewRequest(http.MethodDelete, "/category/1", nil)
// 	resp, _ := app.Test(req)
// 	assert.Equal(t, http.StatusOK, resp.StatusCode)
// }

// func TestGetCategories(t *testing.T) {
// 	app := fiber.New()
// 	app.Get("/categories", Masterhandlers.GetCategories)
// 	mockPool := setupMockDB(t)
// 	defer mockPool.Close()

// 	rows := pgxmock.NewRows([]string{"category_id", "category_name", "super_cat_id", "remarks"}).
// 		AddRow(1, "Electronics", nil, "Main category")

// 	mockPool.ExpectQuery(`SELECT \* FROM get_categories\(\);`).
// 		WillReturnRows(rows)

// 	req := httptest.NewRequest(http.MethodGet, "/categories", nil)
// 	resp, _ := app.Test(req)
// 	assert.Equal(t, http.StatusOK, resp.StatusCode)
// }
