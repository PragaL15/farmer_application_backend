package Masterhandlers

import (
	"context"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/PragaL15/go_newBackend/go_backend/db"
)

type Business struct {
	BID              int64  `json:"bid"`
	BPersonName      string `json:"b_person_name"`
	BRegistrationNum string `json:"b_registration_num"`
	BOwnerName       string `json:"b_owner_name"`
	BCategoryID      int64  `json:"b_category_id"`
	BTypeID          int64  `json:"b_type_id"`
	ActiveStatus     int    `json:"active_status"`
	CreatedAt        string `json:"created_at"`
	UpdatedAt        string `json:"updated_at"`
}

func GetAllBusinesses(c *fiber.Ctx) error {
	rows, err := db.Pool.Query(context.Background(), "SELECT * FROM admin_schema.get_all_businesses()")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch businesses", "details": err.Error()})
	}
	defer rows.Close()

	var businesses []Business
	for rows.Next() {
		var b Business
		if err := rows.Scan(&b.BID, &b.BPersonName, &b.BRegistrationNum, &b.BOwnerName, &b.BCategoryID, &b.BTypeID, &b.ActiveStatus, &b.CreatedAt, &b.UpdatedAt); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error scanning business data", "details": err.Error()})
		}
		businesses = append(businesses, b)
	}
	return c.JSON(businesses)
}

func GetBusinessByID(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID format"})
	}

	var b Business
	err = db.Pool.QueryRow(context.Background(), "SELECT * FROM admin_schema.get_business_by_id($1)", id).
		Scan(&b.BID, &b.BPersonName, &b.BRegistrationNum, &b.BOwnerName, &b.BCategoryID, &b.BTypeID, &b.ActiveStatus, &b.CreatedAt, &b.UpdatedAt)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Business not found"})
	}
	return c.JSON(b)
}

func InsertBusiness(c *fiber.Ctx) error {
	type Request struct {
		BPersonName      string `json:"b_person_name"`
		BRegistrationNum string `json:"b_registration_num"`
		BOwnerName       string `json:"b_owner_name"`
		BCategoryID      int64  `json:"b_category_id"`
		BTypeID          int64  `json:"b_type_id"`
		ActiveStatus     int    `json:"active_status"`
	}

	var req Request
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request format"})
	}

	var newID int64
	err := db.Pool.QueryRow(context.Background(), "SELECT admin_schema.insert_business($1, $2, $3, $4, $5, $6) RETURNING bid",
		req.BPersonName, req.BRegistrationNum, req.BOwnerName, req.BCategoryID, req.BTypeID, req.ActiveStatus).Scan(&newID)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to insert business", "details": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "Business inserted successfully", "bid": newID})
}

func UpdateBusiness(c *fiber.Ctx) error {
	type Request struct {
		BID              int64  `json:"bid"`
		BPersonName      string `json:"b_person_name"`
		BRegistrationNum string `json:"b_registration_num"`
		BOwnerName       string `json:"b_owner_name"`
		BCategoryID      int64  `json:"b_category_id"`
		BTypeID          int64  `json:"b_type_id"`
		ActiveStatus     int    `json:"active_status"`
	}

	var req Request
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request format"})
	}

	_, err := db.Pool.Exec(context.Background(), "SELECT admin_schema.update_business($1, $2, $3, $4, $5, $6, $7)",
		req.BID, req.BPersonName, req.BRegistrationNum, req.BOwnerName, req.BCategoryID, req.BTypeID, req.ActiveStatus)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update business", "details": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "Business updated successfully"})
}
