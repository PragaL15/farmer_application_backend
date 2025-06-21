package unit_test

import (
	"encoding/json"
	"errors"
	"farmerapp/internal/transportation/delivery"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockDeliveryService struct {
	mock.Mock
}

func (m *MockDeliveryService) GetActiveDeliveries(transporterId int) ([]delivery.Delivery, error) {
	args := m.Called()
	return args.Get(0).([]delivery.Delivery), args.Error(1)
}
func (m *MockDeliveryService) GetUpcomingDeliveries(transporterId int) ([]delivery.Delivery, error) {
	args := m.Called()
	return args.Get(0).([]delivery.Delivery), args.Error(1)
}
func (m *MockDeliveryService) GetCompletedDeliveries(transporterId int) ([]delivery.Delivery, error) {
	args := m.Called()
	return args.Get(0).([]delivery.Delivery), args.Error(1)
}
func (m *MockDeliveryService) GetDeliveryHistory(transporterId int) ([]delivery.Delivery, error) {
	args := m.Called()
	return []delivery.Delivery{}, args.Error(1)
}
func (m *MockDeliveryService) ConfirmDelivery(otp int) error {
	args := m.Called()
	return args.Error(1)
}

// Test for the GetActiveDeliveries handler
func TestGetActiveDeliveries_Handler_Success(t *testing.T) {

	app := fiber.New()

	mockService := new(MockDeliveryService)
	mockService.On("GetActiveDeliveries").Return([]delivery.Delivery{
		{JobID: 1, PickupAddress: "p_address_1", DropAddress: "d_address_1"},
		{JobID: 2, PickupAddress: "p_address_2", DropAddress: "d_address_2"},
	}, nil)

	handler := delivery.NewDeliveryHandler(mockService)

	app.Get("transportation/delivery/active-deliveries", handler.GetActiveDeliveries)

	req := httptest.NewRequest("GET", "http://localhost:3000/transportation/delivery/active-deliveries", nil)
	resp, err := app.Test(req)

	assert.NoError(t, err)

	assert.Equal(t, fiber.StatusOK, resp.StatusCode)

	var body map[string][]delivery.Delivery
	err = json.NewDecoder(resp.Body).Decode(&body)
	assert.NoError(t, err)

	expectedDeliveries := []delivery.Delivery{
		{JobID: 1, PickupAddress: "p_address_1", DropAddress: "d_address_1"},
		{JobID: 2, PickupAddress: "p_address_2", DropAddress: "d_address_2"},
	}

	assert.Equal(t, expectedDeliveries, body["deliveries"])

	// Assert that the mock service was called
	mockService.AssertExpectations(t)
}

func TestGetActiveDeliveries_Handler_Error(t *testing.T) {

	app := fiber.New()

	mockService := new(MockDeliveryService)
	mockService.On("GetActiveDeliveries").Return([]delivery.Delivery{}, errors.New("Could not fetch active deliveries. Please try again later."))

	handler := delivery.NewDeliveryHandler(mockService)

	app.Get("transportation/delivery/active-deliveries", handler.GetActiveDeliveries)

	req := httptest.NewRequest("GET", "http://localhost:3000/transportation/delivery/active-deliveries", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)

}

// Test for the GetUpcomingDeliveries handler
func TestGetUpcomingDeliveries_Handler_Success(t *testing.T) {

	app := fiber.New()

	mockService := new(MockDeliveryService)
	mockService.On("GetUpcomingDeliveries").Return([]delivery.Delivery{
		{JobID: 1, PickupAddress: "p_address_1", DropAddress: "d_address_1"},
		{JobID: 2, PickupAddress: "p_address_2", DropAddress: "d_address_2"},
	}, nil)

	handler := delivery.NewDeliveryHandler(mockService)

	app.Get("transportation/delivery/upcoming-deliveries", handler.GetUpcomingDeliveries)

	req := httptest.NewRequest("GET", "http://localhost:3000/transportation/delivery/upcoming-deliveries", nil)
	resp, err := app.Test(req)

	assert.NoError(t, err)

	assert.Equal(t, fiber.StatusOK, resp.StatusCode)

	var body map[string][]delivery.Delivery
	err = json.NewDecoder(resp.Body).Decode(&body)
	assert.NoError(t, err)

	expectedDeliveries := []delivery.Delivery{
		{JobID: 1, PickupAddress: "p_address_1", DropAddress: "d_address_1"},
		{JobID: 2, PickupAddress: "p_address_2", DropAddress: "d_address_2"},
	}

	assert.Equal(t, expectedDeliveries, body["deliveries"])

	// Assert that the mock service was called
	mockService.AssertExpectations(t)
}

func TestGetUpcomingDeliveries_Handler_Error(t *testing.T) {

	app := fiber.New()

	mockService := new(MockDeliveryService)
	mockService.On("GetUpcomingDeliveries").Return([]delivery.Delivery{}, errors.New("Could not fetch upcoming deliveries. Please try again later."))

	handler := delivery.NewDeliveryHandler(mockService)

	app.Get("transportation/delivery/upcoming-deliveries", handler.GetUpcomingDeliveries)

	req := httptest.NewRequest("GET", "http://localhost:3000/transportation/delivery/upcoming-deliveries", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)

}

func TestGetCompletedDeliveries_Handler_Success(t *testing.T) {

	app := fiber.New()

	mockService := new(MockDeliveryService)
	mockService.On("GetCompletedDeliveries").Return([]delivery.Delivery{
		{JobID: 1, PickupAddress: "p_address_1", DropAddress: "d_address_1"},
		{JobID: 2, PickupAddress: "p_address_2", DropAddress: "d_address_2"},
	}, nil)

	handler := delivery.NewDeliveryHandler(mockService)

	app.Get("transportation/delivery/completed-deliveries", handler.GetCompletedDeliveries)

	req := httptest.NewRequest("GET", "http://localhost:3000/transportation/delivery/completed-deliveries", nil)
	resp, err := app.Test(req)

	assert.NoError(t, err)

	assert.Equal(t, fiber.StatusOK, resp.StatusCode)

	var body map[string][]delivery.Delivery
	err = json.NewDecoder(resp.Body).Decode(&body)
	assert.NoError(t, err)

	expectedDeliveries := []delivery.Delivery{
		{JobID: 1, PickupAddress: "p_address_1", DropAddress: "d_address_1"},
		{JobID: 2, PickupAddress: "p_address_2", DropAddress: "d_address_2"},
	}

	assert.Equal(t, expectedDeliveries, body["deliveries"])

	// Assert that the mock service was called
	mockService.AssertExpectations(t)
}

func TestGetCompletedDeliveries_Handler_Error(t *testing.T) {

	app := fiber.New()

	mockService := new(MockDeliveryService)
	mockService.On("GetCompletedDeliveries").Return([]delivery.Delivery{}, errors.New("Could not fetch completed deliveries. Please try again later."))

	handler := delivery.NewDeliveryHandler(mockService)

	app.Get("transportation/delivery/completed-deliveries", handler.GetCompletedDeliveries)

	req := httptest.NewRequest("GET", "http://localhost:3000/transportation/delivery/completed-deliveries", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)

}
