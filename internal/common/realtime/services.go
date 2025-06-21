package realtime

import (
	"fmt"
	"time"
)

type RealtimeServiceInterface interface {
	StoreChatMessage(threadId string, senderId string, message string, sentAt time.Time) error
	GetChatThread(threadId string) (ChatThread, error)
	CreateChatThread(input ChatThread) error
	InsertCallStart(threadID, receiverID, initiatorID string, callStartTime time.Time) (*string, error)
	UpdateCallEnd(callID string, callEndTime time.Time, callStatus string) error
}

type RealtimeService struct {
	repository RealtimeRepositoryInterface
}

func NewRealtimeService(repository RealtimeRepositoryInterface) *RealtimeService {
	return &RealtimeService{repository: repository}
}

func (service *RealtimeService) StoreChatMessage(threadId string, senderId string, message string, sentAt time.Time) error {

	var err error

	err = service.repository.StoreChatMessage(threadId, senderId, message, sentAt)

	if err != nil {
		return fmt.Errorf("Service failed to register user: %w", err)
	}
	return nil

}

func (service *RealtimeService) GetChatThread(threadId string) (ChatThread, error) {

	chatThread, err := service.repository.GetChatThread(threadId)

	if err != nil {
		return chatThread, fmt.Errorf("Service failed to return reciever:%w", err)

	}

	return chatThread, nil

}

func (service *RealtimeService) CreateChatThread(input ChatThread) error {
	var err error

	err = service.repository.CreateChatThread(input)

	if err != nil {
		return fmt.Errorf("service failed to create chat thread: %w", err)
	}
	return nil
}

func (service *RealtimeService) InsertCallStart(threadID, receiverID, initiatorID string, callStartTime time.Time) (*string, error) {
	return service.repository.InsertCallStart(threadID, receiverID, initiatorID, callStartTime)
}

func (service *RealtimeService) UpdateCallEnd(callID string, callEndTime time.Time) error {
	return service.repository.UpdateCallEnd(callID, callEndTime)
}
