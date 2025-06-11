package tracking

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

type TrackingHandler struct {
	service TrackingServiceInterface
}

func NewRequestHandler(service TrackingServiceInterface) *TrackingHandler {
	return &TrackingHandler{service: service}
}

func (handler *TrackingHandler) ValidateQrCode(c *fiber.Ctx) error {

	var req ValidateQrCode

	/*THOUGHT:- TransporterId field to be taken from header instead of body  */

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request payload",
		})
	}

	err := handler.service.ValidateQrCode(req.JobID, req.TransporterID, req.QrCode)
	if err != nil {
		log.Println("Handler error:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not validate QR code. Please try again later.",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "QR code validated successfully",
	})
}
