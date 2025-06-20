package realtime

import "time"

type MessageType string

const (
	TextMessageType      MessageType = "text"
	OfferMessageType     MessageType = "offer"
	AnswerMessageType    MessageType = "answer"
	CandidateMessageType MessageType = "candidate"
	HangUpMessageType    MessageType = "hangup"
	RejectedMessageType  MessageType = "reject"
)

type Message struct {
	Type    MessageType `json:"type"`
	Content interface{} `json:"content"`
}

type ChatMessage struct {
	ThreadID string    `json:"thread_id"`
	SenderID string    `json:"sender_id"`
	Message  string    `json:"message"`
	SentAt   time.Time `json:"sent_at"`
}

type InitiateVoiceCall struct {
	ThreadID    string `json:"thread_id"`
	InitiatorID string `json:"initiator_id"`
}

type ChatThread struct {
	ThreadId string `json:"thread_id"`
	OrderId  string `json:"order_id"`
	User1Id  string `json:"user_1_id"`
	User2Id  string `json:"user_2_id"`
	Role1    int    `json:"role_1"`
	Role2    int    `json:"role_2"`
}

type CallLog struct {
	CallID        string    `json:"call_id"`
	ThreadID      string    `json:"thread_id"`
	InitiatorType string    `json:"initiator_type"`
	InitiatorID   string    `json:"initiator_id"`
	ReceiverID    string    `json:"receiver_id"`
	CallStart     time.Time `json:"call_start"`
	CallEnd       time.Time `json:"call_end"`
	CallDuration  int       `json:"call_duration"`
	CallStatus    string    `json:"call_status"`
}

type EndVoiceCall struct {
	CallID string `json:"call_id"`
}

type ChatThreadModel struct {
	ThreadID int    `json:"thread_id"`
	OrderID  int    `json:"order_id"`
	User1ID  int    `json:"user_1_id"`
	User2ID  int    `json:"user_2_id"`
	Role1    string `json:"role_1"` // Enum: retailer
	Role2    string `json:"role_2"` // Enum: aadthi, transporter
}

type ChatMessageModel struct {
	MessageID int       `json:"message_id"`
	ThreadID  int       `json:"thread_id"`
	SenderID  int       `json:"sender_id"`
	Message   string    `json:"message"`
	SentAt    time.Time `json:"sent_at"`
}

type Call struct {
	CallID        int       `json:"call_id"`
	ThreadID      int       `json:"thread_id"`
	InitiatorType string    `json:"initiator_type"` // Enum: aadthi, transporter
	InitiatorID   int       `json:"initiator_id"`
	ReceiverID    int       `json:"receiver_id"`
	CallStart     time.Time `json:"call_start"`
	CallEnd       time.Time `json:"call_end"`
	CallDuration  int       `json:"call_duration"` // In seconds
	CallStatus    string    `json:"call_status"`   // Enum: completed, missed, rejected
}
