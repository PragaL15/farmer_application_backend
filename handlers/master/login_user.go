package Masterhandlers

import (
	"context"
	"log"
	"net/http"
	"github.com/gofiber/fiber/v2"
	"github.com/PragaL15/go_newBackend/go_backend/db"
)

type UserDetails struct {
	UserID      int    `json:"user_id"`
	Name        string `json:"name"`
	MobileNum   string `json:"mobile_num"`
	Email       string `json:"email"`
	Address     string `json:"address"`
	Pincode     string `json:"pincode"`
	Location    int `json:"location"`
	State       int `json:"state"`
	ActiveStatus int   `json:"active_status"`
	RoleID      int    `json:"role_id"`
}

func LoginUser(c *fiber.Ctx) error {
	// Request body containing email and password
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// Parse request body
	if err := c.BodyParser(&req); err != nil {
		log.Println("Error parsing request body:", err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	// Check if the user exists and fetch the stored password
	var storedPassword string
	var userID int
	var username string
	err := db.Pool.QueryRow(context.Background(), "SELECT user_id, username, password FROM admin_schema.users WHERE email = $1", req.Email).Scan(&userID, &username, &storedPassword)
	if err != nil {
		log.Println("Error checking credentials:", err)
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid email or password"})
	}

	// Compare the provided password with the stored password (no hashing)
	if storedPassword != req.Password {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid email or password"})
	}

	// Fetch user details from user_table
	var userDetails UserDetails
	err = db.Pool.QueryRow(context.Background(), "SELECT user_id, name, mobile_num, email, address, pincode, location, state, active_status, role_id FROM admin_schema.user_table WHERE email = $1", req.Email).
		Scan(&userDetails.UserID, &userDetails.Name, &userDetails.MobileNum, &userDetails.Email, &userDetails.Address, &userDetails.Pincode, &userDetails.Location, &userDetails.State, &userDetails.ActiveStatus, &userDetails.RoleID)
	if err != nil {
		log.Println("Error fetching user details:", err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch user details"})
	}

	// Return user details if credentials are valid
	return c.Status(http.StatusOK).JSON(userDetails)
}
