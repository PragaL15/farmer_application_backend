package handlers

import (
    "log"
    "strconv"

    "github.com/gofiber/fiber/v2"
    "github.com/PragaL15/go_newBackend/go_backend/db"
)

func GetUserBankDetails(c *fiber.Ctx) error {
    // Get user_id from the query parameters
    userID := c.Query("user_id")
    if userID == "" {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "user_id is required"})
    }

    // Convert user_id from string to integer
    parsedUserID, err := strconv.Atoi(userID)
    if err != nil {
        log.Printf("Invalid user_id: %v", err)
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user_id"})
    }

    // Query the database
    query := `
        SELECT 
            id, user_id, card_number, upi_id, ifsc_code, account_number, 
            account_holder_name, bank_name, branch_name, date_of_creation, status
        FROM user_bank_details 
        WHERE user_id = $1;
    `
    rows, err := db.DB.Query(query, parsedUserID)
    if err != nil {
        log.Printf("Failed to execute query: %v", err)
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch bank details"})
    }
    defer rows.Close()

    // Parse rows into a slice of bank details
    bankDetails := []map[string]interface{}{}
    for rows.Next() {
        var (
            id                  int
            userID              int
            cardNumber          *string
            upiID               *string
            ifscCode            *string
            accountNumber       *string
            accountHolderName   *string
            bankName            *string
            branchName          *string
            dateOfCreation      string
            status              bool
        )

        err := rows.Scan(
            &id,
            &userID,
            &cardNumber,
            &upiID,
            &ifscCode,
            &accountNumber,
            &accountHolderName,
            &bankName,
            &branchName,
            &dateOfCreation,
            &status,
        )
        if err != nil {
            log.Printf("Failed to scan row: %v", err)
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error processing bank details"})
        }

        // Handle potential NULL values in the database
        bankDetails = append(bankDetails, map[string]interface{}{
            "id":                  id,
            "user_id":             userID,
            "card_number":         nullToString(cardNumber),
            "upi_id":              nullToString(upiID),
            "ifsc_code":           nullToString(ifscCode),
            "account_number":      nullToString(accountNumber),
            "account_holder_name": nullToString(accountHolderName),
            "bank_name":           nullToString(bankName),
            "branch_name":         nullToString(branchName),
            "date_of_creation":    dateOfCreation,
            "status":              status,
        })
    }

    // Return the bank details as JSON
    if len(bankDetails) == 0 {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "No bank details found for the given user_id"})
    }
    return c.JSON(bankDetails)
}

// Helper function to handle NULL values from the database
func nullToString(value *string) string {
    if value == nil {
        return ""
    }
    return *value
}
