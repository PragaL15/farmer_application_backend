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
var mockOrderStatuse = []OrderStatus{
	{1, "Processing"},
	{2, "Confirmed"},
	{3, "Payment"},
	{4, "Out for Delivery"},
	{5, "Successful"},
	{6, "Cancellation"},
	{7, "Returned"},
}

func mockQuerys(ctx context.Context, query string) ([]OrderStatus, error) {

	return mockOrderStatuse, nil
}

func MockGetOrderStatuse(c *fiber.Ctx) error {
	results, err := mockQuery(context.Background(), "SELECT * FROM order_status_table;")
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch order statuses"})
	}
	return c.JSON(results)
}

func TestGetOrderStatuse(t *testing.T) {

	app := fiber.New()
	app.Get("/order-statuses", MockGetOrderStatuses)
	req := httptest.NewRequest(http.MethodGet, "/order-statuses", nil)
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)
	assert.NoError(t, err, "Expected no error from app.Test()")
	assert.Equal(t, http.StatusOK, resp.StatusCode, "Expected HTTP status 200 OK")

	var result []OrderStatus
	err = json.NewDecoder(resp.Body).Decode(&result)
	assert.NoError(t, err, "Expected no error while decoding JSON response")

	assert.NotEmpty(t, result, "Expected non-empty response")

	assert.Len(t, result, len(mockOrderStatuses), "Expected response length to match mock data")

	expectedStatuses := []string{"Processing", "Confirmed", "Payment", "Out for Delivery", "Successful", "Cancellation", "Returned"}
	for i, status := range expectedStatuses {
		assert.Equal(t, status, result[i].OrderStatus, "Expected order status to match")
	}
}
