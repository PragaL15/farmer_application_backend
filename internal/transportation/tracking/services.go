package tracking

import "fmt"

type TrackingServiceInterface interface {
	ValidateQrCode(jobId int, transporterId int, qrCode string) error
}

type TrackingService struct {
	repository TrackingRepositoryInterface
}

func NewRequestService(repository TrackingRepositoryInterface) *TrackingService {
	return &TrackingService{repository: repository}
}

func (service *TrackingService) ValidateQrCode(jobId int, transporterId int, qrCode string) error {

	err := service.repository.ValidateQrCode(jobId, qrCode, transporterId)
	if err != nil {
		return fmt.Errorf("service failed to get validate QR code: %w", err)
	}
	return nil
}
