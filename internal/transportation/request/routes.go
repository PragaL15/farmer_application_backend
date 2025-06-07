package request

import (
	"github.com/gofiber/fiber/v2"
)

type RequestRouter struct {
	handler *RequestHandler
}

func InitializeRequestModule() *RequestRouter {
	repo := &RequestRepository{}
	service := &RequestService{repository: repo}
	handler := &RequestHandler{service: service}
	return &RequestRouter{handler: handler}
}

func (router *RequestRouter) RegisterRoutes(r fiber.Router) {

	g := r.Group("/requests")
	g.Get("/transport-requests", router.handler.GetTransportRequests)
	g.Post("/accept-transport-request", router.handler.AcceptTransportRequest)
	g.Post("/reject-transport-request", router.handler.RejectTransportRequest)
}
