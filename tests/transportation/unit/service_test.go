package unit_test

import (
	"errors"
	"farmerapp/internal/transportation/delivery"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockDeliveryRepository struct {
	mock.Mock
}

func (m *MockDeliveryRepository) GetDeliveries(deliveryType string, transporterId int) ([]delivery.Delivery, error) {
	args := m.Called(deliveryType)
	return args.Get(0).([]delivery.Delivery), args.Error(1)
}

func TestGetActiveDeliveries_Success(t *testing.T) {
	mockRepo := new(MockDeliveryRepository)
	mockRepo.On("GetDeliveries", "active").Return([]delivery.Delivery{
		{JobID: 1, PickupAddress: "p_address_1", DropAddress: "d_address_1"},
	}, nil)

	service := delivery.NewDeliveryService(mockRepo)
	deliveries, err := service.GetActiveDeliveries(201)

	expectedDeliveries := []delivery.Delivery{
		{JobID: 1, PickupAddress: "p_address_1", DropAddress: "d_address_1"},
	}
	assert.NoError(t, err)
	assert.Len(t, deliveries, 1)
	assert.Equal(t, expectedDeliveries, deliveries)
}

func TestGetActiveDeliveries_Error(t *testing.T) {
	mockRepo := new(MockDeliveryRepository)
	mockRepo.On("GetDeliveries", "active").Return([]delivery.Delivery{}, errors.New("Could not fetch active deliveries. Please try again later."))

	service := delivery.NewDeliveryService(mockRepo)
	deliveries, err := service.GetActiveDeliveries(201)

	assert.Error(t, err)
	assert.Len(t, deliveries, 0)
}

func TestGetUpcomingDeliveries_Success(t *testing.T) {
	mockRepo := new(MockDeliveryRepository)
	mockRepo.On("GetDeliveries", "upcoming").Return([]delivery.Delivery{
		{JobID: 1, PickupAddress: "p_address_1", DropAddress: "d_address_1"},
	}, nil)

	service := delivery.NewDeliveryService(mockRepo)
	deliveries, err := service.GetUpcomingDeliveries(201)

	expectedDeliveries := []delivery.Delivery{
		{JobID: 1, PickupAddress: "p_address_1", DropAddress: "d_address_1"},
	}
	assert.NoError(t, err)
	assert.Len(t, deliveries, 1)
	assert.Equal(t, expectedDeliveries, deliveries)
}

func TestGetUpcomingDeliveries_Error(t *testing.T) {
	mockRepo := new(MockDeliveryRepository)
	mockRepo.On("GetDeliveries", "upcoming").Return([]delivery.Delivery{}, errors.New("Could not fetch upcoming deliveries. Please try again later."))

	service := delivery.NewDeliveryService(mockRepo)
	deliveries, err := service.GetUpcomingDeliveries(201)

	assert.Error(t, err)
	assert.Len(t, deliveries, 0)
}

func TestGetCompletedDeliveries_Success(t *testing.T) {
	mockRepo := new(MockDeliveryRepository)
	mockRepo.On("GetDeliveries", "completed").Return([]delivery.Delivery{
		{JobID: 1, PickupAddress: "p_address_1", DropAddress: "d_address_1"},
	}, nil)

	service := delivery.NewDeliveryService(mockRepo)
	deliveries, err := service.GetCompletedDeliveries(201)

	expectedDeliveries := []delivery.Delivery{
		{JobID: 1, PickupAddress: "p_address_1", DropAddress: "d_address_1"},
	}
	assert.NoError(t, err)
	assert.Len(t, deliveries, 1)
	assert.Equal(t, expectedDeliveries, deliveries)
}

func TestGetCompletedDeliveries_Error(t *testing.T) {
	mockRepo := new(MockDeliveryRepository)
	mockRepo.On("GetDeliveries", "completed").Return([]delivery.Delivery{}, errors.New("Could not fetch completed deliveries. Please try again later."))

	service := delivery.NewDeliveryService(mockRepo)
	deliveries, err := service.GetCompletedDeliveries(201)

	assert.Error(t, err)
	assert.Len(t, deliveries, 0)
}
