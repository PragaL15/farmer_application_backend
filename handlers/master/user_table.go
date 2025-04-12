package Masterhandlers

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/PragaL15/go_newBackend/go_backend/db"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v4"
)

type User struct {
	UserID       int    `json:"user_id"`
	Name         string `json:"name"`
	MobileNum    string `json:"mobile_num"`
	Email        string `json:"email"`
	Address      string `json:"address"`
	Pincode      string `json:"pincode"`
	Location     int    `json:"location_id"`
	LocationName string `json:"location_name"`
	State        int    `json:"state_id"`
	StateName    string `json:"state_name"`
	ActiveStatus int    `json:"active_status"`
	RoleID       int    `json:"role_id"`
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
			&u.UserID,
			&u.Name,
			&u.MobileNum,
			&u.Email,
			&u.Address,
			&u.Pincode,
			&u.Location,
			&u.LocationName,
			&u.State,
			&u.StateName,
			&u.ActiveStatus,
			&u.RoleID,
		)

	if err != nil {
		log.Println("Error fetching user:", err)
		if err == pgx.ErrNoRows || err == sql.ErrNoRows {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
		}
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}

	return c.JSON(u)
}


func InsertUser(c *fiber.Ctx) error {
	var u User
	if err := c.BodyParser(&u); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	_, err := db.Pool.Exec(context.Background(),
    "SELECT admin_schema.insert_user($1, $2, $3, $4, $5, $6, $7, $8, $9)",
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
	type UpdateUser struct {
		Name         string `json:"name"`
		MobileNum    string `json:"mobile_num"`
		Email        string `json:"email"`
		Address      string `json:"address"`
		Pincode      string `json:"pincode"`
		Location     int    `json:"location"` // FIXED: was string
		State        int    `json:"state"`    // FIXED: was string
		ActiveStatus bool   `json:"active_status"`
		RoleID       int64  `json:"role_id"`
	}

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	var u UpdateUser
	if err := c.BodyParser(&u); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	activeInt := 0
	if u.ActiveStatus {
		activeInt = 1
	}

	_, err = db.Pool.Exec(context.Background(),
		"SELECT admin_schema.update_user($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)",
		id, u.Name, u.MobileNum, u.Email, u.Address, u.Pincode,
		u.Location, u.State, activeInt, u.RoleID,
	)
	if err != nil {
		log.Println("Error updating user:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update user"})
	}

	return c.JSON(fiber.Map{"message": "User updated successfully"})
}
