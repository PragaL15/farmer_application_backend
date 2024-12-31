package handlers

import (
    "github.com/gofiber/fiber/v2"
    "github.com/PragaL15/go_newBackend/go_backend/db"
    "log"
		"database/sql" 
)

func GetAllUsers(c *fiber.Ctx) error {
	
    rows, err := db.DB.Query("SELECT * FROM user_table")
    if err != nil {
        log.Printf("Failed to fetch users: %v", err)
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch users"})
    }
    defer rows.Close()

    users := []map[string]interface{}{}
    for rows.Next() {
        var user = make(map[string]interface{})
        var userID, userTypeID, mandiID sql.NullInt64
        var name, mobileNum, email, address, pincode, location, businessLicenseNo, validity, gstNo, businessName, businessType, mandiTypeID, remarks, col1 sql.NullString
        var dtOfCommenceBusiness, expiryDt sql.NullTime

        err = rows.Scan(&userID, &userTypeID, &name, &dtOfCommenceBusiness, &mobileNum, &email, &address,
            &pincode, &location, &businessLicenseNo, &validity, &gstNo, &expiryDt, &businessName, &businessType, &mandiID, &mandiTypeID, &remarks, &col1)
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

        users = append(users, user)
    }

    return c.JSON(users)
}