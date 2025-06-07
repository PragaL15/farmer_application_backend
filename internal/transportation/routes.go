package transportation

import (
	"farmerapp/internal/transportation/delivery"
	"farmerapp/internal/transportation/request"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(r fiber.Router) {
	deliveryRouter := delivery.InitializeDeliveryModule()
	deliveryRouter.RegisterRoutes(r)

	requestRouter := request.InitializeRequestModule()
	requestRouter.RegisterRoutes(r)
}
