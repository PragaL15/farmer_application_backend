package tests

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

type OrderStatus struct {
	OrderStatusID int    `json:"order_status_id"`
	OrderStatus   string `json:"order_status"`
}

var mockOrderStatuses = []OrderStatus{
	{1, "Processing"},
	{2, "Confirmed"},
	{3, "Payment"},
	{4, "Out for Delivery"},
	{5, "Successful"},
	{6, "Cancellation"},
	{7, "Returned"},
}

func mockQuery(ctx context.Context, query string) ([]OrderStatus, error) {
	return mockOrderStatuses, nil
}

func MockGetOrderStatuses(c *fiber.Ctx) error {
	results, _ := mockQuery(context.Background(), "SELECT * FROM order_status_table;")
	return c.JSON(results)
}

func TestGetOrderStatuses(t *testing.T) {
	app := fiber.New()
	app.Get("/order-statuses", MockGetOrderStatuses)
	req := httptest.NewRequest(http.MethodGet, "/order-statuses", nil)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	var result []OrderStatus
	err = json.NewDecoder(resp.Body).Decode(&result)
	assert.NoError(t, err)
	assert.NotEmpty(t, result)
	assert.Len(t, result, 7)
	expectedStatuses := []string{"Processing", "Confirmed", "Payment", "Out for Delivery", "Successful", "Cancellation", "Returned"}
	for i, status := range expectedStatuses {
		assert.Equal(t, status, result[i].OrderStatus)
	}
}
