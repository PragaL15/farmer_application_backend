package Masterhandlers

import (
	"context"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/PragaL15/go_newBackend/go_backend/db"
)

type BusinessBranch struct {
	BranchID        int    `json:"b_branch_id"`
	BID             int    `json:"bid"`
	ShopName        string `json:"b_shop_name"`
	TypeID          int    `json:"b_type_id"`
	Location        int    `json:"b_location"`
	State           int    `json:"b_state"`
	MandiID         int    `json:"b_mandi_id"`
	Address         string `json:"b_address"`
	Email           string `json:"b_email"`
	Number          string `json:"b_number"`
	GSTNum          string `json:"b_gst_num"`
	PANNum          string `json:"b_pan_num"`
	PrivilegeUser   int    `json:"b_privilege_user"`
	EstablishedYear string `json:"b_established_year"`
	ActiveStatus    int    `json:"active_status"`
}

func GetAllBusinessBranches(c *fiber.Ctx) error {
	rows, err := db.Pool.Query(context.Background(), "SELECT * FROM admin_schema.get_all_business_branches()")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch business branches", "details": err.Error()})
	}
	defer rows.Close()

	var branches []BusinessBranch
	for rows.Next() {
		var b BusinessBranch
		if err := rows.Scan(&b.BranchID, &b.BID, &b.ShopName, &b.TypeID, &b.Location, &b.State, &b.MandiID, &b.Address, &b.Email, &b.Number, &b.GSTNum, &b.PANNum, &b.PrivilegeUser, &b.EstablishedYear, &b.ActiveStatus); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error scanning branch data", "details": err.Error()})
		}
		branches = append(branches, b)
	}
	return c.JSON(branches)
}

func GetBusinessBranchByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID format"})
	}

	var b BusinessBranch
	err = db.Pool.QueryRow(context.Background(), "SELECT * FROM admin_schema.get_business_branch_by_id($1)", id).
		Scan(&b.BranchID, &b.BID, &b.ShopName, &b.TypeID, &b.Location, &b.State, &b.MandiID, &b.Address, &b.Email, &b.Number, &b.GSTNum, &b.PANNum, &b.PrivilegeUser, &b.EstablishedYear, &b.ActiveStatus)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Business branch not found"})
	}
	return c.JSON(b)
}

func InsertBusinessBranch(c *fiber.Ctx) error {
	var req BusinessBranch
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request format"})
	}

	var newID int
	err := db.Pool.QueryRow(context.Background(), "SELECT admin_schema.insert_business_branch($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14) RETURNING b_branch_id",
		req.BID, req.ShopName, req.TypeID, req.Location, req.State, req.MandiID, req.Address, req.Email, req.Number, req.GSTNum, req.PANNum, req.PrivilegeUser, req.EstablishedYear, req.ActiveStatus).Scan(&newID)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to insert business branch", "details": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "Business branch inserted successfully", "b_branch_id": newID})
}

func UpdateBusinessBranch(c *fiber.Ctx) error {
	var req BusinessBranch
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request format"})
	}

	_, err := db.Pool.Exec(context.Background(), "SELECT admin_schema.update_business_branch($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)",
		req.BranchID, req.BID, req.ShopName, req.TypeID, req.Location, req.State, req.MandiID, req.Address, req.Email, req.Number, req.GSTNum, req.PANNum, req.PrivilegeUser, req.EstablishedYear, req.ActiveStatus)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update business branch", "details": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "Business branch updated successfully"})
}
