package request

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

type RequestHandler struct {
	service RequestServiceInterface
}

func NewRequestHandler(service RequestServiceInterface) *RequestHandler {
	return &RequestHandler{service: service}
}

func (handler *RequestHandler) GetTransportRequests(c *fiber.Ctx) error {

	region := c.Query("region")
	deliveryRequests, err := handler.service.GetTransportRequests(region)
	if err != nil {
		log.Println("Handler error:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not fetch delivery requests. Please try again later.",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"delivery_requests": deliveryRequests,
	})
}

func (handler *RequestHandler) AcceptTransportRequest(c *fiber.Ctx) error {

	var req AcceptJobRequest

	/*THOUGHT:- TransporterId field to be taken from header instead of body  */

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request payload",
		})
	}

	err := handler.service.AcceptTransportRequest(req.TransporterID, req.JobID, req.VehicleID, req.DriverID)
	if err != nil {
		log.Println("Handler error:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not process transport request acceptance. Please try again later.",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Request accepted successfully",
	})
}

func (handler *RequestHandler) RejectTransportRequest(c *fiber.Ctx) error {

	var req RejectJobRequest

	/*THOUGHT:- TransporterId field to be taken from header instead of body  */

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request payload",
		})
	}

	err := handler.service.RejectTransportRequest(req.TransporterID, req.JobID)
	if err != nil {
		log.Println("Handler error:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not process transport request rejection. Please try again later.",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Request rejected successfully",
	})
}
