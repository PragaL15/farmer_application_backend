package delivery

import "fmt"

type DeliveryServiceInterface interface {
	GetActiveDeliveries(transporterId string) ([]Delivery, error)
	GetUpcomingDeliveries(transporterId string) ([]Delivery, error)
	GetCompletedDeliveries(transporterId string) ([]Delivery, error)
	GetDeliveryHistory(transporterId string) ([]Delivery, error)
	ConfirmDelivery(otp int) error
}

type DeliveryService struct {
	repository DeliveryRepositoryInterface
}

func NewDeliveryService(repository DeliveryRepositoryInterface) *DeliveryService {
	return &DeliveryService{repository: repository}
}

func (service *DeliveryService) GetActiveDeliveries(transporterId string) ([]Delivery, error) {

	deliveries, err := service.repository.GetDeliveries("active", transporterId)
	if err != nil {
		return nil, fmt.Errorf("service failed to get active deliveries: %w", err)
	}
	return deliveries, nil
}

func (service *DeliveryService) GetUpcomingDeliveries(transporterId string) ([]Delivery, error) {

	deliveries, err := service.repository.GetDeliveries("upcoming", transporterId)
	if err != nil {
		return nil, fmt.Errorf("service failed to get upcoming deliveries: %w", err)
	}
	return deliveries, nil

}

func (service *DeliveryService) GetCompletedDeliveries(transporterId string) ([]Delivery, error) {

	deliveries, err := service.repository.GetDeliveries("completed", transporterId)
	if err != nil {
		return nil, fmt.Errorf("service failed to get completed deliveries: %w", err)
	}
	return deliveries, nil

}

func (service *DeliveryService) GetDeliveryHistory(transporterId string) ([]Delivery, error) {
	deliveries, err := service.repository.GetDeliveries("history", transporterId)
	if err != nil {
		return nil, fmt.Errorf("service failed to get delivery history: %w", err)
	}
	if deliveries != nil {
		return deliveries, nil

	}
	return []Delivery{}, nil
}

func (service *DeliveryService) ConfirmDelivery(otp int) error {
	// OTP based confirmation ?
	return nil
}
