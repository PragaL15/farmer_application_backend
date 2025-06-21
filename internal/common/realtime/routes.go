package realtime

import (
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

type RealtimeRouter struct {
	handler *RealtimeHandler
}

func InitializeRealtimeModule() *RealtimeRouter {
	handler := &RealtimeHandler{}
	return &RealtimeRouter{handler: handler}
}

func (router *RealtimeRouter) RegisterRoutes(r fiber.Router) {

	g := r.Group("")
	g.Get("/ws/:userId", websocket.New(router.handler.HandleWebSocket))
	g.Post("/new-thread", router.handler.CreateChatThread)
}
