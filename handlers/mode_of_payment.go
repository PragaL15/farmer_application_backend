package handlers
import (
	"context"
	"log"
	"github.com/gofiber/fiber/v2"
	"github.com/PragaL15/go_newBackend/go_backend/db"
)
func GetModeOfPayments(c *fiber.Ctx) error {
	rows, err := db.Pool.Query(context.Background(), "SELECT * FROM get_mode_of_payment_details()")
	if err != nil {
		log.Printf("Failed to fetch payment details: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch payment details"})
	}
	defer rows.Close()

	var payments []map[string]interface{}

	for rows.Next() {
		var (
			id, payMode, payType int
			payModeName, payTypeName *string
		)

		if err := rows.Scan(&id, &payMode, &payType, &payModeName, &payTypeName); err != nil {
			log.Printf("Error scanning row: %v", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error processing data"})
		}

		payment := map[string]interface{}{
			"id":             id,
			"pay_mode":       payMode,
			"pay_type":       payType,
			"pay_mode_name":  payModeName,
			"pay_type_name":  payTypeName,
		}
		payments = append(payments, payment)
	}

	return c.JSON(payments)
}