package delivery

import (
	"github.com/gofiber/fiber/v2"
)

type DeliveryRouter struct {
	handler *DeliveryHandler
}

func InitializeDeliveryModule() *DeliveryRouter {
	repo := &DeliveryRepository{}
	service := &DeliveryService{repository: repo}
	handler := &DeliveryHandler{service: service}
	return &DeliveryRouter{handler: handler}
}

func (router *DeliveryRouter) RegisterRoutes(r fiber.Router) {

	g := r.Group("/delivery")
	g.Get("/active-deliveries", router.handler.GetActiveDeliveries)
	g.Get("/upcoming-deliveries", router.handler.GetUpcomingDeliveries)
	g.Get("/completed-deliveries", router.handler.GetCompletedDeliveries)
}
