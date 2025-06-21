package realtime

import (
	"encoding/json"
	"log"
	"time"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

type RealtimeHandler struct {
	service RealtimeServiceInterface
}

func (handler *RealtimeHandler) HandleWebSocket(c *websocket.Conn) {
	userID := c.Params("userId")
	hub.AddClient(userID, c)
	defer func() {
		hub.RemoveClient(userID)
		c.Close()
	}()

	for {
		var msg Message
		var chatThread ChatThread
		var reciever string
		var err error
		if err := c.ReadJSON(&msg); err != nil {
			log.Println("WebSocket read error:", err)
			break
		}

		contentBytes, _ := json.Marshal(msg.Content)
		log.Printf("Received message of message type %s : %s\n", msg.Type, contentBytes)

		switch msg.Type {
		case TextMessageType:
			textContent, ok := msg.Content.(ChatMessage)
			if !ok {
				log.Printf("Invalid content type for text message: %T\n", msg.Content)
				continue
			}

			chatThread, err = handler.service.GetChatThread(textContent.ThreadID)

			if err != nil {
				log.Printf("Failed to get chat thread: %v\n", err)
				continue
			}

			if chatThread.User1Id != textContent.SenderID && chatThread.User2Id != textContent.SenderID {
				log.Printf("Sender doesn't belong to this thread")
				continue
			}

			if textContent.SenderID == chatThread.User1Id {
				reciever = chatThread.User2Id
			} else {
				reciever = chatThread.User1Id
			}

			textContent.SentAt = time.Now()
			if err := handler.service.StoreChatMessage(textContent.ThreadID, textContent.SenderID, textContent.Message, textContent.SentAt); err != nil {
				log.Printf("Failed to store chat message: %v\n", err)
				continue
			}
		case OfferMessageType:

			offerContent, ok := msg.Content.(InitiateVoiceCall)

			if !ok {
				log.Printf("Invalid content type for voice call: %T\n", msg.Content)
				continue
			}

			chatThread, err = handler.service.GetChatThread(offerContent.ThreadID)

			if err != nil {
				log.Printf("Failed to get chat thread: %v\n", err)
				continue
			}

			if chatThread.User1Id != offerContent.InitiatorID && chatThread.User2Id != offerContent.InitiatorID {
				log.Printf("Call Initiator doesn't belong to this thread")
				continue
			}

			if offerContent.InitiatorID == chatThread.User1Id {
				reciever = chatThread.User2Id
			} else {
				reciever = chatThread.User1Id
			}

			callStartTime := time.Now()

			threadID := offerContent.ThreadID
			initiatorID := offerContent.InitiatorID
			receiverID := reciever

			callID, err := handler.service.InsertCallStart(threadID, initiatorID, receiverID, callStartTime)
			if err != nil {
				log.Printf("Failed to insert call start: %v", err)
				continue
			}

			callTimeoutCh := make(chan bool, 1)

			go func() {
				select {
				case <-time.After(30 * time.Second): // 30 seconds timeout
					// Timeout reached: Mark as missed
					// callStatus = "missed"
					callEndTime := time.Now()
					// callDuration := int(callEndTime.Sub(callStartTime).Seconds())

					if err := handler.service.UpdateCallEnd(*callID, callEndTime, "missed"); err != nil {
						log.Printf("Failed to update missed call log: %v", err)
					}

				case <-callTimeoutCh:
					// Timeout was cancelled because the call was answered
					return
				}
			}()
		case AnswerMessageType:

		case HangUpMessageType:
			hangUpcontent, ok := msg.Content.(EndVoiceCall)

			if !ok {
				log.Printf("Invalid content type for voice call: %T\n", msg.Content)
				continue
			}

			callEndTime := time.Now()
			if err := handler.service.UpdateCallEnd(hangUpcontent.CallID, callEndTime, "completed"); err != nil {
				log.Printf("Failed to update call log: %v", err)
			}
		case RejectedMessageType:

			rejectContent, ok := msg.Content.(EndVoiceCall)

			if !ok {
				log.Printf("Invalid content type for voice call: %T\n", msg.Content)
				continue
			}

			callEndTime := time.Now()
			if err := handler.service.UpdateCallEnd(rejectContent.CallID, callEndTime, "rejected"); err != nil {
				log.Printf("Failed to update call log: %v", err)
			}

		default:
			log.Printf("Unsupported message type: %s\n", msg.Type)
			continue

		}

		if receiverConn, exists := hub.GetClient(reciever); exists {
			if err := receiverConn.WriteJSON(msg); err != nil {
				log.Printf("WebSocket write error to user %s: %v\n", reciever, err)
			}
		} else {
			log.Printf("User %s not connected\n", reciever)
		}
	}
}

func (handler *RealtimeHandler) CreateChatThread(c *fiber.Ctx) error {

	var userInput ChatThread

	if err := c.BodyParser(&userInput); err != nil {
		log.Println("Error parsing request body:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	err := handler.service.CreateChatThread(userInput)
	if err != nil {
		log.Println("Handler error:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not create new chat thread. Please try again later.",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Thread created successfully",
	})
}
