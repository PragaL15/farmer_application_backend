package handlers

import (
    "github.com/gofiber/fiber/v2"
    "github.com/PragaL15/go_newBackend/go_backend/db"
    "log"
    "database/sql"
    "github.com/go-playground/validator/v10"
    "context"
    "strconv"
)

func GetAllUsers(c *fiber.Ctx) error {
    rows, err := db.Pool.Query(context.Background(), "SELECT * FROM public.user_table;")
    if err != nil {
        log.Printf("Failed to fetch users: %v", err)
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch users"})
    }
    defer rows.Close()

    users := []map[string]interface{}{}
    for rows.Next() {
        var user = make(map[string]interface{})
        var userID, userTypeID, mandiID sql.NullInt64
        var name, mobileNum, email, address, pincode, location, businessLicenseNo, validity, gstNo, businessName,state, businessType, mandiTypeID, remarks, col1 , col2 sql.NullString
        var dtOfCommenceBusiness, expiryDt sql.NullTime

        err = rows.Scan(&userID, &userTypeID, &name, &dtOfCommenceBusiness, &mobileNum, &email, &address,
            &pincode, &location, &businessLicenseNo, &validity, &gstNo, &expiryDt, &businessName, &businessType, &mandiID, &mandiTypeID, &remarks, &col1, &col2, &state)
        if err != nil {
            log.Printf("Failed to scan user: %v", err)
            continue
        }

        user["user_id"] = userID.Int64
        user["user_type_id"] = userTypeID.Int64
        user["name"] = name.String
        user["dt_of_commence_business"] = dtOfCommenceBusiness.Time
        user["mobile_num"] = mobileNum.String
        user["email"] = email.String
        user["address"] = address.String
        user["pincode"] = pincode.String
        user["location"] = location.String
        user["state"] = location.String
        user["business_license_no"] = businessLicenseNo.String
        user["validity"] = validity.String
        user["gst_no"] = gstNo.String
        user["expiry_dt"] = expiryDt.Time
        user["business_name"] = businessName.String
        user["business_type"] = businessType.String
        user["mandi_id"] = mandiID.Int64
        user["mandi_type_id"] = mandiTypeID.String
        user["remarks"] = remarks.String
        user["col1"] = col1.String
        user["col2"] = col2.String

        users = append(users, user)
    }

    return c.JSON(users)
}


func InsertUser(c *fiber.Ctx) error {
	type Request struct {
		UserTypeID            int    `json:"user_type_id"`
		Name                  string `json:"name" validate:"required,max=255"`
		DtOfCommenceBusiness  string `json:"dt_of_commence_business"`
		MobileNum             string `json:"mobile_num" validate:"required,max=15"`
		Email                 string `json:"email" validate:"required,email,max=255"`
		Address               string `json:"address"`
		Pincode               string `json:"pincode" validate:"max=10"`
		Location              int    `json:"location"`
		BusinessLicenseNo     string `json:"business_license_no"`
		Validity              string `json:"validity"`
		GstNo                 string `json:"gst_no" validate:"max=15"`
		ExpiryDt              string `json:"expiry_dt"`
		BusinessName          string `json:"business_name" validate:"max=255"`
		BusinessType          string `json:"business_type" validate:"max=255"`
		MandiID               int    `json:"mandi_id"`
		MandiTypeID           string `json:"mandi_type_id"`
		Remarks               string `json:"remarks"`
		Col1                  string `json:"col1"`
		Col2                  string `json:"col2"`
		State                 int    `json:"state"`
	}

	var req Request
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	_, err := db.Pool.Exec(context.Background(), `
		CALL insert_user($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20);
	`, req.UserTypeID, req.Name, req.DtOfCommenceBusiness, req.MobileNum, req.Email, req.Address, req.Pincode, req.Location, req.BusinessLicenseNo, req.Validity, req.GstNo, req.ExpiryDt, req.BusinessName, req.BusinessType, req.MandiID, req.MandiTypeID, req.Remarks, req.Col1, req.Col2, req.State)

	if err != nil {
		log.Printf("Failed to insert user: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to insert user"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "User added successfully"})
}
func UpdateUser(c *fiber.Ctx) error {
	type Request struct {
		UserID               int    `json:"user_id" validate:"required"`
		UserTypeID           int    `json:"user_type_id"`
		Name                 string `json:"name" validate:"required,max=255"`
		DtOfCommenceBusiness string `json:"dt_of_commence_business"`
		MobileNum            string `json:"mobile_num" validate:"required,max=15"`
		Email                string `json:"email" validate:"required,email,max=255"`
		Address              string `json:"address"`
		Pincode              string `json:"pincode" validate:"max=10"`
		Location             int    `json:"location"`
		BusinessLicenseNo    string `json:"business_license_no"`
		Validity             string `json:"validity"`
		GstNo                string `json:"gst_no" validate:"max=15"`
		ExpiryDt             string `json:"expiry_dt"`
		BusinessName         string `json:"business_name" validate:"max=255"`
		BusinessType         string `json:"business_type" validate:"max=255"`
		MandiID              int    `json:"mandi_id"`
		MandiTypeID          string `json:"mandi_type_id"`
		Remarks              string `json:"remarks"`
		Col1                 string `json:"col1"`
		Col2                 string `json:"col2"`
		State                int    `json:"state"`
	}

	var req Request
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	_, err := db.Pool.Exec(context.Background(), `
		CALL update_user($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21);
	`, req.UserID, req.UserTypeID, req.Name, req.DtOfCommenceBusiness, req.MobileNum, req.Email, req.Address, req.Pincode, req.Location, req.BusinessLicenseNo, req.Validity, req.GstNo, req.ExpiryDt, req.BusinessName, req.BusinessType, req.MandiID, req.MandiTypeID, req.Remarks, req.Col1, req.Col2, req.State)

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

	_, err = db.Pool.Exec(context.Background(), `
		CALL delete_user($1);
	`, idInt)

	if err != nil {
		log.Printf("Failed to delete user: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete user"})
	}

	return c.JSON(fiber.Map{"message": "User deleted successfully"})
}
