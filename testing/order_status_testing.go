// package testing

// import (
// 	"bytes"
// 	"context"
// 	"encoding/json"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/gofiber/fiber/v2"
// 	"github.com/stretchr/testify/assert"
// )

// var mockDB = map[int]string{
// 	1: "Processing",
// 	2: "Shipped",
// }

// func mockQuery(ctx context.Context, query string) ([]map[string]interface{}, error) {
// 	results := []map[string]interface{}{}
// 	for id, status := range mockDB {
// 		results = append(results, map[string]interface{}{
// 			"order_id":     id,
// 			"order_status": status,
// 		})
// 	}
// 	return results, nil
// }

// func TestGetOrderStatuses(t *testing.T) {
// 	app := fiber.New()

// 	app.Get("/order-statuses", GetOrderStatuses)

// 	req := httptest.NewRequest(http.MethodGet, "/order-statuses", nil)
// 	resp, err := app.Test(req)

// 	assert.NoError(t, err)
// 	assert.Equal(t, http.StatusOK, resp.StatusCode)

// 	var result []map[string]interface{}
// 	json.NewDecoder(resp.Body).Decode(&result)

// 	assert.NotEmpty(t, result)
// 	assert.Equal(t, "Processing", result[0]["order_status"])
// }
