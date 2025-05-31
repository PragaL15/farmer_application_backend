package delivery

import "fmt"

type DeliveryServiceInterface interface {
	GetActiveDeliveries() ([]Delivery, error)
	GetUpcomingDeliveries() ([]Delivery, error)
	GetCompletedDeliveries() ([]Delivery, error)
	GetDeliveryHistory() ([]string, error)
	ConfirmDelivery(otp int) error
}

type DeliveryService struct {
	repository DeliveryRepositoryInterface
}

func (service *DeliveryService) GetActiveDeliveries() ([]Delivery, error) {

	deliveries, err := service.repository.GetDeliveries("active")
	if err != nil {
		return nil, fmt.Errorf("service failed to get active deliveries: %w", err)
	}
	return deliveries, nil
}

func (service *DeliveryService) GetUpcomingDeliveries() ([]Delivery, error) {

	deliveries, err := service.repository.GetDeliveries("upcoming")
	if err != nil {
		return nil, fmt.Errorf("service failed to get upcoming deliveries: %w", err)
	}
	return deliveries, nil

}

func (service *DeliveryService) GetCompletedDeliveries() ([]Delivery, error) {

	deliveries, err := service.repository.GetDeliveries("completed")
	if err != nil {
		return nil, fmt.Errorf("service failed to get completed deliveries: %w", err)
	}
	return deliveries, nil

}
