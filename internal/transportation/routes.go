package transportation

import (
	"farmerapp/internal/transportation/delivery"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(r fiber.Router) {
	deliveryRouter := delivery.InitializeDeliveryModule()
	deliveryRouter.RegisterRoutes(r)
}
