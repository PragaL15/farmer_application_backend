package tracking

import (
	"github.com/gofiber/fiber/v2"
)

type TrackingRouter struct {
	handler *TrackingHandler
}

func InitializeTrackingModule() *TrackingRouter {
	repo := &TrackingRepository{}
	service := &TrackingService{repository: repo}
	handler := &TrackingHandler{service: service}
	return &TrackingRouter{handler: handler}
}

func (router *TrackingRouter) RegisterRoutes(r fiber.Router) {

	g := r.Group("/tracking")
	g.Post("/validate-qrcode", router.handler.ValidateQrCode)
}
