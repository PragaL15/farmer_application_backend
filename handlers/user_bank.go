package handlers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/PragaL15/go_newBackend/go_backend/db"
)

func GetUserBankDetails(c *fiber.Ctx) error {
	userID := c.Query("user_id") 
	if userID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "user_id is required"})
	}
	query := `
		SELECT * FROM get_user_bank_details($1);
	`
	rows, err := db.DB.Query(query, userID)
	if err != nil {
		log.Printf("Failed to fetch bank details: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch bank details"})
	}
	defer rows.Close()

	bankDetails := []map[string]interface{}{}
	for rows.Next() {
		var (
			id                  int
			userID              int
			cardNumber          string
			upiID               string
			ifscCode            string
			accountNumber       string
			accountHolderName   string
			bankName            string
			branchName          string
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
			continue
		}
		bankDetails = append(bankDetails, map[string]interface{}{
			"id":                  id,
			"user_id":             userID,
			"card_number":         cardNumber,
			"upi_id":              upiID,
			"ifsc_code":           ifscCode,
			"account_number":      accountNumber,
			"account_holder_name": accountHolderName,
			"bank_name":           bankName,
			"branch_name":         branchName,
			"date_of_creation":    dateOfCreation,
			"status":              status,
		})
	}

	return c.JSON(bankDetails)
}
