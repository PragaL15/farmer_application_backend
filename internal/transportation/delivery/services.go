package delivery

import "fmt"

type DeliveryServiceInterface interface {
	GetActiveDeliveries(transporterId int) ([]Delivery, error)
	GetUpcomingDeliveries(transporterId int) ([]Delivery, error)
	GetCompletedDeliveries(transporterId int) ([]Delivery, error)
	GetDeliveryHistory(transporterId int) ([]Delivery, error)
	ConfirmDelivery(otp int) error
}

type DeliveryService struct {
	repository DeliveryRepositoryInterface
}

func NewDeliveryService(repository DeliveryRepositoryInterface) *DeliveryService {
	return &DeliveryService{repository: repository}
}

func (service *DeliveryService) GetActiveDeliveries(transporterId int) ([]Delivery, error) {

	deliveries, err := service.repository.GetDeliveries("active", transporterId)
	if err != nil {
		return nil, fmt.Errorf("service failed to get active deliveries: %w", err)
	}
	return deliveries, nil
}

func (service *DeliveryService) GetUpcomingDeliveries(transporterId int) ([]Delivery, error) {

	deliveries, err := service.repository.GetDeliveries("upcoming", transporterId)
	if err != nil {
		return nil, fmt.Errorf("service failed to get upcoming deliveries: %w", err)
	}
	return deliveries, nil

}

func (service *DeliveryService) GetCompletedDeliveries(transporterId int) ([]Delivery, error) {

	deliveries, err := service.repository.GetDeliveries("completed", transporterId)
	if err != nil {
		return nil, fmt.Errorf("service failed to get completed deliveries: %w", err)
	}
	return deliveries, nil

}

func (service *DeliveryService) GetDeliveryHistory(transporterId int) ([]Delivery, error) {
	deliveries, err := service.repository.GetDeliveries("history", transporterId)
	if err != nil {
		return nil, fmt.Errorf("service failed to get delivery history: %w", err)
	}
	return deliveries, nil
}

func (service *DeliveryService) ConfirmDelivery(otp int) error {
	// OTP based confirmation ?
	return nil
}
