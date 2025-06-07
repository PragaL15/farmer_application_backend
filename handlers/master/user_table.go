package Masterhandlers

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"farmerapp/go_backend/db"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v4"
)

// User represents a user in the system
type User struct {
	UserID       int    `json:"user_id"`
	Name         string `json:"name"`
	MobileNum    string `json:"mobile_num"`
	Email        string `json:"email"`
	Address      string `json:"address"`
	Pincode      string `json:"pincode"`
	Location     int    `json:"location"`
	LocationName string `json:"location_name"`
	State        int    `json:"state"`
	StateName    string `json:"state_name"`
	ActiveStatus int    `json:"active_status"`
	RoleID       int    `json:"role_id"`
}

// GetAllUsers godoc
// @Summary      Get all users
// @Description  Fetches all users from the database
// @Tags         Users
// @Accept       json
// @Produce      json
// @Success      200 {array} Masterhandlers.User
// @Router       /getUsers [get]
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

// GetUserByID godoc
// @Summary      Get user by ID
// @Description  Fetch a single user by ID
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Success      200  {object}  Masterhandlers.User
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Router       /getUsers/{id} [get]
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

// InsertUser godoc
// @Summary      Create a new user
// @Description  Insert a new user into the system
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        user  body      Masterhandlers.User  true  "User data"
// @Success      201   {object}  map[string]string
// @Failure      400   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Router       /userTableDetails [post]
func InsertUser(c *fiber.Ctx) error {
	var u User

	// Step 1: Parse the incoming body
	if err := c.BodyParser(&u); err != nil {
		log.Println("Body parsing failed:", err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	// Step 2: Debug log the received values
	log.Printf("Received User: %+v\n", u)

	// Step 3: Basic validation
	if u.RoleID == 0 {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid or missing role_id"})
	}
	if u.State == 0 {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid or missing state"})
	}

	// Step 4: Call the DB insert function
	_, err := db.Pool.Exec(context.Background(),
		"SELECT admin_schema.insert_user($1, $2, $3, $4, $5, $6, $7, $8, $9)",
		u.Name, u.MobileNum, u.Email, u.Address, u.Pincode,
		u.Location, u.State, u.ActiveStatus, u.RoleID,
	)

	if err != nil {
		log.Println("Error inserting user:", err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to insert user. Please ensure all foreign key values (state, role_id) exist.",
		})
	}

	// Step 5: Success response
	return c.Status(http.StatusCreated).JSON(fiber.Map{"message": "User created successfully"})
}

// UpdateUser godoc
// @Summary      Update an existing user
// @Description  Update user data by ID
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        id    path      int               true  "User ID"
// @Param        user  body      Masterhandlers.User  true  "Updated user data"
// @Success      200   {object}  map[string]string
// @Failure      400   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Router       /userUpdate/{id} [put]
func UpdateUser(c *fiber.Ctx) error {
	type UpdateUser struct {
		Name         string `json:"name"`
		MobileNum    string `json:"mobile_num"`
		Email        string `json:"email"`
		Address      string `json:"address"`
		Pincode      string `json:"pincode"`
		Location     int    `json:"location"`
		State        int    `json:"state"`
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
