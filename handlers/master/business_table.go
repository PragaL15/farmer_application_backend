package Masterhandlers

import (
	"context"
	"log"
	"strconv"

	"github.com/PragaL15/go_newBackend/go_backend/db"
	"github.com/gofiber/fiber/v2"
)

type Business struct {
	BID              int64  `json:"bid"`
	BRegistrationNum string `json:"b_registration_num"`
	BOwnerName       string `json:"b_owner_name"`
	BCategoryID      int64  `json:"b_category_id"`
	BTypeID          int64  `json:"b_type_id"`
	IsActive         bool   `json:"is_active"`
}

// GetAllBusinesses godoc
// @Summary Get all businesses
// @Description Retrieve all business entries from the database
// @Tags Business
// @Accept json
// @Produce json
// @Success 200 {array} Masterhandlers.Business
// @Failure 500 {object} map[string]string
// @Router /business [get]
func GetAllBusinesses(c *fiber.Ctx) error {
	rows, err := db.Pool.Query(context.Background(), "SELECT * FROM admin_schema.get_all_businesses()")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to fetch businesses",
			"details": err.Error(),
		})
	}
	defer rows.Close()

	var businesses []Business
	for rows.Next() {
		var b Business
		if err := rows.Scan(&b.BID, &b.BRegistrationNum, &b.BOwnerName, &b.BCategoryID, &b.BTypeID, &b.IsActive); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":   "Error scanning business data",
				"details": err.Error(),
			})
		}
		businesses = append(businesses, b)
	}
	return c.JSON(businesses)
}

// GetBusinessByID godoc
// @Summary Get a business by ID
// @Description Retrieve a business entry by its ID
// @Tags Business
// @Accept json
// @Produce json
// @Param id path int true "Business ID"
// @Success 200 {object} Masterhandlers.Business
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /business/{id} [get]
func GetBusinessByID(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID format"})
	}

	var b Business
	err = db.Pool.QueryRow(context.Background(), "SELECT * FROM admin_schema.get_business_by_id($1)", id).
		Scan(&b.BID, &b.BRegistrationNum, &b.BOwnerName, &b.BCategoryID, &b.BTypeID, &b.IsActive)

	if err != nil {
		log.Println("Error fetching business:", err)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Business not found"})
	}
	return c.JSON(b)
}

// InsertBusiness godoc
// @Summary Insert a new business
// @Description Create a new business record
// @Tags Business
// @Accept json
// @Produce json
// @Param data body Masterhandlers.Business true "Business Data"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /business [post]
func InsertBusiness(c *fiber.Ctx) error {
	var req Business
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request format",
		})
	}

	var newID int64
	err := db.Pool.QueryRow(context.Background(),
		`SELECT admin_schema.insert_business($1, $2, $3, $4, $5)`,
		req.BRegistrationNum, req.BOwnerName, req.BCategoryID, req.BTypeID, req.IsActive,
	).Scan(&newID)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to insert business",
			"details": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Business inserted successfully",
		"bid":     newID,
	})
}

// UpdateBusiness godoc
// @Summary Update a business
// @Description Update an existing business record
// @Tags Business
// @Accept json
// @Produce json
// @Param data body Masterhandlers.Business true "Business Update Data"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /business [put]
func UpdateBusiness(c *fiber.Ctx) error {
	var req Business
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request format",
		})
	}

	_, err := db.Pool.Exec(context.Background(), `
		SELECT admin_schema.update_business($1, $2, $3, $4, $5, $6)
	`, req.BID, req.BRegistrationNum, req.BOwnerName, req.BCategoryID, req.BTypeID, req.IsActive)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to update business",
			"details": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Business updated successfully",
	})
}
