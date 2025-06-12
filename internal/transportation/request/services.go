package request

import "fmt"

type RequestServiceInterface interface {
	GetTransportRequests(region string) ([]TransportRequest, error)
	AcceptTransportRequest(transporterId int, jobId int, vehicleId int, driverId int) error
	RejectTransportRequest(transporterId int, jobId int) error
}

type RequestService struct {
	repository RequestRepositoryInterface
}

func NewRequestService(repository RequestRepositoryInterface) *RequestService {
	return &RequestService{repository: repository}
}

func (service *RequestService) GetTransportRequests(region string) ([]TransportRequest, error) {

	deliveryRequests, err := service.repository.GetTransportRequests(region)
	if err != nil {
		return nil, fmt.Errorf("service failed to get delivery requests: %w", err)
	}
	return deliveryRequests, nil
}

func (service *RequestService) AcceptTransportRequest(transporterId int, jobId int, vehicleId int, driverId int) error {

	err := service.repository.AcceptTransportRequest(transporterId, jobId, vehicleId, driverId)
	if err != nil {
		return fmt.Errorf("service failed to accept transport request: %w", err)
	}
	return nil
}

func (service *RequestService) RejectTransportRequest(transporterId int, jobId int) error {

	err := service.repository.RejectTransportRequest(transporterId, jobId)
	if err != nil {
		return fmt.Errorf("service failed to reject transport request: %w", err)
	}
	return nil
}
