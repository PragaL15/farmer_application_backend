package transportation

import (
	"farmerapp/internal/common/auth"
	"farmerapp/internal/transportation/delivery"
	"farmerapp/internal/transportation/request"
	"farmerapp/internal/transportation/tracking"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(r fiber.Router) {
	r.Use(auth.AuthMiddleware())

	deliveryRouter := delivery.InitializeDeliveryModule()
	deliveryRouter.RegisterRoutes(r)

	requestRouter := request.InitializeRequestModule()
	requestRouter.RegisterRoutes(r)

	trackingRouter := tracking.InitializeTrackingModule()
	trackingRouter.RegisterRoutes(r)
}
