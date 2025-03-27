package Masterhandlers

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/PragaL15/go_newBackend/go_backend/db"
)

type User struct {
	UserID       int64  `json:"user_id"`
	Name         string `json:"name"`
	MobileNum    string `json:"mobile_num"`
	Email        string `json:"email"`
	Address      string `json:"address"`
	Pincode      string `json:"pincode"`
	Location     int64  `json:"location"`
	State        int64  `json:"state"`
	ActiveStatus int64  `json:"active_status"`
	RoleID       int64  `json:"role_id"`
}

func GetAllUsers(c *fiber.Ctx) error {
	rows, err := db.Pool.Query(context.Background(), "SELECT * FROM admin_schema.get_all_users()")
	if err != nil {
		log.Println("Error fetching users:", err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch users"})
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User
		if err := rows.Scan(
			&u.UserID, &u.Name, &u.MobileNum, &u.Email, &u.Address, &u.Pincode,
			&u.Location, &u.State, &u.ActiveStatus, &u.RoleID,
		); err != nil {
			log.Println("Error scanning row:", err)
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to parse data"})
		}
		users = append(users, u)
	}
	return c.JSON(users)
}

func GetUserByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	var u User
	err = db.Pool.QueryRow(context.Background(), "SELECT * FROM admin_schema.get_user_by_id($1)", id).
		Scan(
			&u.UserID, &u.Name, &u.MobileNum, &u.Email, &u.Address, &u.Pincode,
			&u.Location, &u.State, &u.ActiveStatus, &u.RoleID,
		)
	if err != nil {
		log.Println("Error fetching user:", err)
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}
	return c.JSON(u)
}

func InsertUser(c *fiber.Ctx) error {
	var u User
	if err := c.BodyParser(&u); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	_, err := db.Pool.Exec(context.Background(),
		"CALL admin_schema.insert_user($1, $2, $3, $4, $5, $6, $7, $8, $9)",
		u.Name, u.MobileNum, u.Email, u.Address, u.Pincode,
		u.Location, u.State, u.ActiveStatus, u.RoleID,
	)
	if err != nil {
		log.Println("Error inserting user:", err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to insert user"})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{"message": "User created successfully"})
}

func UpdateUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	var u User
	if err := c.BodyParser(&u); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	_, err = db.Pool.Exec(context.Background(),
		"CALL admin_schema.update_user($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)",
		id, u.Name, u.MobileNum, u.Email, u.Address, u.Pincode,
		u.Location, u.State, u.ActiveStatus, u.RoleID,
	)
	if err != nil {
		log.Println("Error updating user:", err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update user"})
	}

	return c.JSON(fiber.Map{"message": "User updated successfully"})
}
