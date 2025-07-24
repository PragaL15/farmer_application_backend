package auth

import (
	"farmerapp/internal/common/utils"
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	service AuthServiceInterface
}

func NewAuthHandler(service AuthServiceInterface) *AuthHandler {
	return &AuthHandler{service: service}
}

func (handler *AuthHandler) RegisterUser(c *fiber.Ctx) error {

	var userInput struct {
		Identifier string `json:"identifier"`
		Password   string `json:"password"`
		Name       string `json:"name"`
		Address    string `json:"address"`
		Pincode    string `json:"pincode"`
		Location   int    `json:"location"`
		State      int    `json:"state"`
		RoleID     int    `json:"role_id"`
		Active     int    `json:"active_status"`
	}

	if err := c.BodyParser(&userInput); err != nil {
		log.Println("Error parsing request body:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	identifier := strings.TrimSpace(userInput.Identifier)

	if identifier == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Email or phone number is required",
		})
	}

	var identifierType IdentifierType

	if isPhone, phoneErr := utils.IsPhoneNumber(identifier); isPhone {
		if phoneErr != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": phoneErr.Error()})
		}
		identifierType = IdentifierPhone
	} else {
		// Then try email
		if err := utils.ValidateEmail(identifier); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid email or phone number"})
		}
		identifierType = IdentifierEmail
	}

	err := handler.service.RegisterUser(identifier, identifierType, userInput.Password, userInput.Name, userInput.Address,
		userInput.Pincode, userInput.Location, userInput.State, userInput.RoleID, userInput.Active)
	if err != nil {
		log.Println("Handler error:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not register user. Please try again later." + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User registered successfully",
	})

}

func (handler *AuthHandler) LoginUser(c *fiber.Ctx) error {
	var userInput struct {
		Identifier string `json:"identifier"`
		Password   string `json:"password"`
	}

	if err := c.BodyParser(&userInput); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	identifier := strings.TrimSpace(userInput.Identifier)
	if identifier == "" || userInput.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Identifier and password are required",
		})
	}

	var identifierType IdentifierType

	if isPhone, phoneErr := utils.IsPhoneNumber(identifier); isPhone {
		if phoneErr != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": phoneErr.Error()})
		}
		identifierType = IdentifierPhone
	} else {
		// Then try email
		if err := utils.ValidateEmail(identifier); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid email or phone number"})
		}
		identifierType = IdentifierEmail
	}

<<<<<<< HEAD
	UserID, roleID, accessToken, refreshToken, err := handler.service.LoginUser(identifier, identifierType, userInput.Password)
=======
	roleID, accessToken, refreshToken, err := handler.service.LoginUser(identifier, identifierType, userInput.Password)
>>>>>>> 89c8837d6c691851780de9cd6d15506fa99fd796
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Login failed: " + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"user_id":       UserID,
		"message":       "User logged in successfully",
		"role_id":       roleID,
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}

func (handler *AuthHandler) RefreshToken(c *fiber.Ctx) error {
	var input struct {
		RefreshToken string `json:"refresh_token"`
	}

	if err := c.BodyParser(&input); err != nil || input.RefreshToken == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid refresh token"})
	}

	userId, err := handler.service.ValidateRefreshToken(input.RefreshToken)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid or expired refresh token"})
	}

	accessToken, err := utils.GenerateAccessToken(userId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not generate access token"})
	}

	return c.JSON(fiber.Map{
		"access_token": accessToken,
	})
}
