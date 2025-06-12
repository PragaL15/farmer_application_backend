package Masterhandlers

import (
	"context"
	"strconv"

	"farmerapp/go_backend/db"

	"github.com/gofiber/fiber/v2"
)

// BusinessUser struct
type BusinessUser struct {
	BID          int64  `json:"b_id"`
	UserName     string `json:"user_name"`
	ActiveStatus int    `json:"active_status"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
	IsLocked     bool   `json:"is_locked"`
}

// Get all business users
func GetAllBusinessUsers(c *fiber.Ctx) error {
	rows, err := db.Pool.Query(context.Background(), "SELECT * FROM admin_schema.get_all_business_users()")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch business users", "details": err.Error()})
	}
	defer rows.Close()

	var users []BusinessUser
	for rows.Next() {
		var user BusinessUser
		if err := rows.Scan(&user.BID, &user.UserName, &user.ActiveStatus, &user.CreatedAt, &user.UpdatedAt, &user.IsLocked); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error scanning business user data", "details": err.Error()})
		}
		users = append(users, user)
	}
	return c.JSON(users)
}

// Get business user by ID
func GetBusinessUserByID(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID format"})
	}

	var user BusinessUser
	err = db.Pool.QueryRow(context.Background(), "SELECT * FROM admin_schema.get_business_user_by_id($1)", id).
		Scan(&user.BID, &user.UserName, &user.ActiveStatus, &user.CreatedAt, &user.UpdatedAt, &user.IsLocked)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Business user not found"})
	}
	return c.JSON(user)
}

// Update business user
func UpdateBusinessUser(c *fiber.Ctx) error {
	type Request struct {
		BID          int64  `json:"b_id"`
		UserName     string `json:"user_name"`
		ActiveStatus int    `json:"active_status"`
		IsLocked     bool   `json:"is_locked"`
	}

	var req Request
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request format"})
	}

	_, err := db.Pool.Exec(context.Background(), "SELECT admin_schema.update_business_user($1, $2, $3, $4)",
		req.BID, req.UserName, req.ActiveStatus, req.IsLocked)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update business user", "details": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "Business user updated successfully"})
}
