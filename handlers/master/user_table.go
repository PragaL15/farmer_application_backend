package Masterhandlers

import (
	"context"
	"database/sql"
	"log"
	"strconv"

	"github.com/PragaL15/go_newBackend/go_backend/db"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func GetAllUsers(c *fiber.Ctx) error {
	rows, err := db.Pool.Query(context.Background(), "SELECT * FROM get_all_users()")
	if err != nil {
		log.Printf("Failed to fetch users: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch users"})
	}
	defer rows.Close()

	users := []map[string]interface{}{}
	for rows.Next() {
		var user = make(map[string]interface{})
		var userID, userTypeID, location, state, status sql.NullInt64
		var userTypeName, name, mobileNum, email, address, pincode, locationName, stateName sql.NullString

		err = rows.Scan(&userID, &userTypeID, &userTypeName, &name, &mobileNum, &email, &address, &pincode, 
			&location, &locationName, &state, &stateName, &status)
		if err != nil {
			log.Printf("Failed to scan user: %v", err)
			continue
		}

		user["user_id"] = userID.Int64
		user["user_type_id"] = userTypeID.Int64
		user["user_type_name"] = userTypeName.String
		user["name"] = name.String
		user["mobile_num"] = mobileNum.String
		user["email"] = email.String
		user["address"] = address.String
		user["pincode"] = pincode.String
		user["location"] = location.Int64
		user["location_name"] = locationName.String
		user["state"] = state.Int64
		user["state_name"] = stateName.String
		user["status"] = status.Int64

		users = append(users, user)
	}

	return c.JSON(users)
}


func InsertUser(c *fiber.Ctx) error {
	type Request struct {
		UserTypeID int    `json:"user_type_id" validate:"required"`
		Name       string `json:"name" validate:"required,max=255"`
		MobileNum  string `json:"mobile_num" validate:"required,max=15"`
		Email      string `json:"email" validate:"required,email,max=255"`
		Address    string `json:"address"`
		Pincode    string `json:"pincode" validate:"max=10"`
		Location   int    `json:"location"`
		State      int    `json:"state"`
		Status     int    `json:"status"`
	}

	var req Request
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	_, err := db.Pool.Exec(context.Background(), "CALL insert_user($1, $2, $3, $4, $5, $6, $7, $8, $9)",
		req.UserTypeID, req.Name, req.MobileNum, req.Email, req.Address, req.Pincode, req.Location, req.State, req.Status)

	if err != nil {
		log.Printf("Failed to insert user: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to insert user"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "User added successfully"})
}

func UpdateUser(c *fiber.Ctx) error {
	type Request struct {
		UserID     int    `json:"user_id" validate:"required"`
		UserTypeID int    `json:"user_type_id"`
		Name       string `json:"name" validate:"required,max=255"`
		MobileNum  string `json:"mobile_num" validate:"required,max=15"`
		Email      string `json:"email" validate:"required,email,max=255"`
		Address    string `json:"address"`
		Pincode    string `json:"pincode" validate:"max=10"`
		Location   int    `json:"location"`
		State      int    `json:"state"`
		Status     int    `json:"status"`
	}

	var req Request
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	_, err := db.Pool.Exec(context.Background(), "CALL update_user($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)",
		req.UserID, req.UserTypeID, req.Name, req.MobileNum, req.Email, req.Address, req.Pincode, req.Location, req.State, req.Status)

	if err != nil {
		log.Printf("Failed to update user: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update user"})
	}

	return c.JSON(fiber.Map{"message": "User updated successfully"})
}

func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ID is required"})
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID format"})
	}

	_, err = db.Pool.Exec(context.Background(), "CALL delete_user($1)", idInt)
	if err != nil {
		log.Printf("Failed to delete user: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete user"})
	}

	return c.JSON(fiber.Map{"message": "User deleted successfully"})
}
