package handlers

import (
	"context"
	"database/sql"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/PragaL15/go_newBackend/go_backend/db"
	"github.com/go-playground/validator/v10"
)

func InsertBusiness(c *fiber.Ctx) error {
	type Request struct {
		BTypeID     int     `json:"b_typeid" validate:"required"`
		BName       string  `json:"b_name" validate:"required"`
		BLocationID int     `json:"b_location_id" validate:"required"`
		BStateID    int     `json:"b_state_id" validate:"required"`
		BMandiID    *int    `json:"b_mandiid"`            
		BAddress    string  `json:"b_address" validate:"required"`
		BPhoneNum   *string `json:"b_phone_num"`         
		BEmail      string  `json:"b_email" validate:"required,email"`
		BGstNum     string  `json:"b_gstnum" validate:"required,max=20"`
		BPanNum     string  `json:"b_pannum" validate:"required,max=10"`
	}
	var req Request
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}
	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	bMandiID := interface{}(nil)
	if req.BMandiID != nil {
		bMandiID = *req.BMandiID
	}
	bPhoneNum := interface{}(nil)
	if req.BPhoneNum != nil && *req.BPhoneNum != "" {
		bPhoneNum = *req.BPhoneNum
	}
 _, err := db.Pool.Exec(context.Background(),
		`SELECT insert_business($1, $2, $3, $4, $5, $6, $7, $8, $9, $10);`,
		req.BTypeID, req.BName, req.BLocationID, req.BStateID, req.BAddress, req.BEmail, req.BGstNum, req.BPanNum, bMandiID, bPhoneNum)
	if err != nil {
		log.Printf("Failed to insert business: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to insert business"})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Business added successfully"})
}

func UpdateBusiness(c *fiber.Ctx) error {
	type Request struct {
		BID         int     `json:"bid" validate:"required"`
		BTypeID     int     `json:"b_typeid" validate:"required"`
		BLocationID int     `json:"b_location_id" validate:"required"`
		BStateID    int     `json:"b_state_id" validate:"required"`
		BMandiID    *int    `json:"b_mandiid"`
		BAddress    string  `json:"b_address" validate:"required"`
		BPhoneNum   *string `json:"b_phone_num"`
		BEmail      string  `json:"b_email" validate:"email"`
		BGstNum     string  `json:"b_gstnum" validate:"max=20"`
		BPanNum     string  `json:"b_pan" validate:"max=10"`
	}
	var req Request
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}
	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	_, err := db.Pool.Exec(context.Background(),
		`CALL update_business($1, $2, $3, $4, $5, $6, $7, $8, $9, $10);`,
		req.BID, req.BTypeID, req.BLocationID, req.BStateID, req.BMandiID, req.BAddress, req.BPhoneNum, req.BEmail, req.BGstNum, req.BPanNum)
	if err != nil {
		log.Printf("Failed to update business: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update business"})
	}
	return c.JSON(fiber.Map{"message": "Business updated successfully"})
}

func DeleteBusiness(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ID is required"})
	}
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID format"})
	}
	_, err = db.Pool.Exec(context.Background(), `CALL delete_business($1);`, idInt)
	if err != nil {
		log.Printf("Failed to delete business: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete business"})
	}
	return c.JSON(fiber.Map{"message": "Business deleted successfully"})
}

func GetBusinesses(c *fiber.Ctx) error {
	rows, err := db.Pool.Query(context.Background(), "SELECT * FROM get_all_businesses_func()")
	if err != nil {
		log.Printf("Failed to fetch business records: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch businesses"})
	}
	defer rows.Close()
	var businesses []map[string]interface{}
	for rows.Next() {
		var bid, bTypeID, bLocationID, bStateID, bMandiID sql.NullInt32
		var bTypeName, bName, bAddress, bPhoneNum, bEmail, bGstNum, bPanNum, stateName, location sql.NullString
		if err := rows.Scan(&bid, &bTypeID, &bTypeName, &bName, &bLocationID, &location, &bStateID, &stateName, 
			&bMandiID, &bAddress, &bPhoneNum, &bEmail, &bGstNum, &bPanNum); err != nil {
			log.Printf("Error scanning row: %v", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error processing data"})
		}
		businesses = append(businesses, map[string]interface{}{
			"bid":          bid.Int32,
			"b_type_id":    bTypeID.Int32,
			"b_typename":   bTypeName.String,
			"b_name":       bName.String,
			"b_location_id": bLocationID.Int32,
			"location":     location.String,
			"b_state_id":   bStateID.Int32,
			"state_name":   stateName.String,
			"b_mandiid":    bMandiID.Int32,
			"b_address":    bAddress.String,
			"b_phone_num":  bPhoneNum.String,
			"b_email":      bEmail.String,
			"b_gstnum":     bGstNum.String,
			"b_pan":        bPanNum.String,
		})
	}
	return c.JSON(businesses)
}

