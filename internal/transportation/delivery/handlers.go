package delivery

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

type DeliveryHandler struct {
	service DeliveryServiceInterface
}

func NewDeliveryHandler(service DeliveryServiceInterface) *DeliveryHandler {
	return &DeliveryHandler{service: service}
}

func (handler *DeliveryHandler) GetActiveDeliveries(c *fiber.Ctx) error {

	deliveries, err := handler.service.GetActiveDeliveries()
	if err != nil {
		log.Println("Handler error:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not fetch active deliveries. Please try again later.",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"deliveries": deliveries,
	})
}

func (handler *DeliveryHandler) GetUpcomingDeliveries(c *fiber.Ctx) error {

	deliveries, err := handler.service.GetUpcomingDeliveries()
	if err != nil {
		log.Println("Handler error:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not fetch upcoming deliveries. Please try again later.",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"deliveries": deliveries,
	})
}

func (handler *DeliveryHandler) GetCompletedDeliveries(c *fiber.Ctx) error {

	deliveries, err := handler.service.GetCompletedDeliveries()
	if err != nil {
		log.Println("Handler error:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not fetch completed deliveries. Please try again later.",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"deliveries": deliveries,
	})
}
