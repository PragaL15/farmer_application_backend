package realtime

import (
	"context"
	"farmerapp/go_backend/db"
	"fmt"
	"time"

	"github.com/jackc/pgconn"
)

type RealtimeRepositoryInterface interface {
	StoreChatMessage(threadId string, senderId string, message string, sentAt time.Time) error
	GetChatThread(threadId string) (ChatThread, error)
	CreateChatThread(input ChatThread) error
	InsertCallStart(threadID, receiverID, initiatorID string, callStartTime time.Time) (*string, error)
	UpdateCallEnd(callID string, callEndTime time.Time) error
}

type RealtimeRepository struct {
}

func (repository *RealtimeRepository) StoreChatMessage(threadId string, senderId string, message string, sentAt time.Time) error {

	// TODO: replace table below with real table name
	var query string
	query = `INSERT INTO admin.chats (thread_id,sender_id,message,sent_at) VALUES ($1,$2,$3,$4)`

	_, err := db.Pool.Exec(context.Background(), query, threadId, senderId, message, sentAt)

	if err != nil {
		return fmt.Errorf("Failed to save user chat: %w", err)
	}
	return nil
}

func (repository *RealtimeRepository) GetChatThread(threadId string) (ChatThread, error) {

	var chatThread ChatThread
	var query string
	query = `
		SELECT
			thread_id,user_1_id,user_2_id
		FROM
			admin.chat_thread_table
		WHERE
			thread_id = $1
		`
	err := db.Pool.QueryRow(context.Background(), query, threadId).
		Scan(&chatThread.ThreadId, &chatThread.User1Id, &chatThread.User2Id)
	if err != nil {
		return chatThread, fmt.Errorf("Query failed: %w", err)
	}

	return chatThread, nil

}

func (repository *RealtimeRepository) CreateChatThread(input ChatThread) error {

	var query string
	query = `INSERT INTO admin.chat_thread_table VALUES ($1,$2.$3,$4,$5,$6)`

	_, err := db.Pool.Exec(context.Background(), query, input.ThreadId, input.OrderId, input.User1Id, input.User2Id, input.Role1, input.Role2)

	if err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			if pgErr.Code == "23505" {
				return fmt.Errorf("Thread realated to this order already exists")
			}
		}
		return fmt.Errorf("Failed to create new thread: %w", err)
	}
	return nil

}

func (repository *RealtimeRepository) InsertCallStart(threadID, receiverID, initiatorID string, callStartTime time.Time) (*string, error) {

	var callID *string

	query := `
			INSERT INTO call_logs (thread_id, initiator_id, receiver_id, call_start, call_end, call_duration, call_status)
			VALUES ($1, $2, $3, $4, NULL, NULL, 'initiated')
			RETURNING call_id
	        `

	err := db.Pool.QueryRow(context.Background(), query, threadID, initiatorID, receiverID, callStartTime).Scan(&callID)

	if err != nil {
		return nil, fmt.Errorf("Failed to insert call record: %w", err)
	}
	return callID, nil
}

func (repository *RealtimeRepository) UpdateCallEnd(callID string, callEndTime time.Time, callStatus string) error {

	var callStartTime time.Time

	query := `
			SELECT call_start FROM call_logs WHERE call_id = $1
		`

	err := db.Pool.QueryRow(context.Background(), query, callID).Scan(&callStartTime)

	if err != nil {
		return fmt.Errorf("Failed to fetch call record: %w", err)

	}

	callDuration := int(callEndTime.Sub(callStartTime).Seconds())

	switch callStatus {
	case "missed":
		query = `
			UPDATE call_logs
			SET call_end = $1, call_duration = $2, call_status = 'missed'
			WHERE call_id = $3
		`
	case "completed":
		query = `
			UPDATE call_logs
			SET call_end = $1, call_duration = $2, call_status = 'completed'
			WHERE call_id = $3
		`
	case "rejected":
		query = `
			UPDATE call_logs
			SET call_end = $1, call_duration = $2, call_status = 'rejected'
			WHERE call_id = $3
		`
	default:
		return fmt.Errorf("Invalid Call Status")

	}

	_, err = db.Pool.Exec(context.Background(), query, callEndTime, callDuration, callID)

	if err != nil {
		return fmt.Errorf("Failed to update call record: %w", err)
	}

	return nil

}
