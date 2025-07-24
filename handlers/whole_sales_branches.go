package handlers

import (
	"context"
	"farmerapp/go_backend/db"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type WholeSellerBranch struct {
	BranchID        int    `json:"branch_id"`
	BusinessID      int    `json:"bid"`
	ShopName        string `json:"b_shop_name"`
	TypeId          int    `json:"b_type_id"`
	Location        string `json:"b_location" `
	State           string `json:"b_state"`
	Address         string `json:"b_address"`
	Email           string `json:"b_email"`
	Number          string `json:"b_number"`
	GSTNumber       string `json:"b_gst_num"`
	PAN             string `json:"b_pan_num"`
	PrivilegedUser  int    `json:"b_privilege_user"`
	EstablishedYear string `json:"b_established_year"`
	ActiveStatus    int    `json:"active_status"`
}

func GetAllBusinessesOfWholeSaler(c *fiber.Ctx) error {

	userId := c.Params("user_Id")

	query := "SELECT * FROM admin_schema.get_all_business_branches(" + userId + ")"

	rows, err := db.Pool.Query(context.Background(), query)
	if err != nil {
		log.Println("Error executing wholeseller branches details query:", err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error":  "Database query failed",
			"detail": err.Error(),
		})
	}
	defer rows.Close()

	wholeSellerBranches := []WholeSellerBranch{}
	for rows.Next() {
		var sd WholeSellerBranch
		if err := rows.Scan(
			&sd.BranchID,
			&sd.BusinessID,
			&sd.ShopName,
			&sd.TypeId,
			&sd.Location,
			&sd.State,
			&sd.Address,
			&sd.Email,
			&sd.GSTNumber,
			&sd.PAN,
			&sd.PrivilegedUser,
			&sd.EstablishedYear,
			&sd.ActiveStatus,
		); err != nil {
			log.Println("Error getting WholeSeller Branch row:", err)
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"error":  "Failed to scan row",
				"detail": err.Error(),
			})
		}
		wholeSellerBranches = append(wholeSellerBranches, sd)
	}

	return c.Status(http.StatusOK).JSON(wholeSellerBranches)

}
