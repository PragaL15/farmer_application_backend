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

	/*TODO:-  get transporterId from request header and pass it down right to the repo level  */
	transporterId := 201
	deliveries, err := handler.service.GetActiveDeliveries(transporterId)
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

	/*TODO:-  get transporterId from request header and pass it down right to the repo level  */
	transporterId := 201
	deliveries, err := handler.service.GetUpcomingDeliveries(transporterId)
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
	/*TODO:-  get transporterId from request header and pass it down right to the repo level  */
	transporterId := 201
	deliveries, err := handler.service.GetCompletedDeliveries(transporterId)
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

func (handler *DeliveryHandler) GetDeliveryHistory(c *fiber.Ctx) error {
	/*TODO:-  get transporterId from request header and pass it down right to the repo level  */
	transporterId := 201
	deliveries, err := handler.service.GetDeliveryHistory(transporterId)
	if err != nil {
		log.Println("Handler error:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not fetch delivery history. Please try again later.",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"deliveries": deliveries,
	})
}

func (handler *DeliveryHandler) ConfirmDelivery(c *fiber.Ctx) error {

	enteredOtp := 9329 /*can be encrypted?*/
	err := handler.service.ConfirmDelivery(enteredOtp)
	// err can be request format error, wrong otp error, db error
	if err != nil {
		log.Println("Handler error:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not fetch completed deliveries. Please try again later.",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Delivery Confirmed",
	})
}
