package handlers

import (
	"context"
	"farmerapp/go_backend/db"
	"log"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func InsertUserBankDetail(c *fiber.Ctx) error {
	type Request struct {
		UserID            int    `json:"user_id" validate:"required,min=1"`
		CardNumber        string `json:"card_number" validate:"required,len=16,numeric"`
		UpiID             string `json:"upi_id" validate:"required,email"`
		IFSCCode          string `json:"ifsc_code" validate:"required,len=11"`
		AccountNumber     string `json:"account_number" validate:"required,len=12,numeric"`
		AccountHolderName string `json:"account_holder_name" validate:"required,max=100"`
		BankName          string `json:"bank_name" validate:"required,max=100"`
		BranchName        string `json:"branch_name" validate:"required"`
		Status            bool   `json:"status" validate:"required"`
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
        SELECT manage_user_bank_details(
            'INSERT',
            NULL,
            $1, $2, $3, $4, $5, $6, $7, $8,
            $9
        );
    `, req.UserID, req.CardNumber, req.UpiID, req.IFSCCode, req.AccountNumber, req.AccountHolderName, req.BankName, req.BranchName, req.Status)
	if err != nil {
		log.Printf("Failed to insert user bank detail: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to insert user bank detail"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "User bank detail added successfully"})
}

func UpdateUserBankDetail(c *fiber.Ctx) error {
	type Request struct {
		ID                int    `json:"id"`
		CardNumber        string `json:"card_number"`
		UpiID             string `json:"upi_id"`
		IFSCCode          string `json:"ifsc_code"`
		AccountNumber     string `json:"account_number"`
		AccountHolderName string `json:"account_holder_name"`
		BankName          string `json:"bank_name"`
		BranchName        string `json:"branch_name"`
		Status            bool   `json:"status"`
	}
	var req Request
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}
	_, err := db.Pool.Exec(context.Background(), `
        CALL manage_user_bank_details(
            'UPDATE',
            $1,
            NULL, $2, $3, $4, $5, $6, $7, $8,
            $9
        );
    `, req.ID, req.CardNumber, req.UpiID, req.IFSCCode, req.AccountNumber, req.AccountHolderName, req.BankName, req.BranchName, req.Status)

	if err != nil {
		log.Printf("Failed to update user bank detail: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update user bank detail"})
	}
	return c.JSON(fiber.Map{"message": "User bank detail updated successfully"})
}

func DeleteUserBankDetail(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ID is required"})
	}
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID format"})
	}
	_, err = db.Pool.Exec(context.Background(), `
        CALL manage_user_bank_details(
            'DELETE',
            $1,
            NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL,
            NULL
        );
    `, idInt)
	if err != nil {
		log.Printf("Failed to delete user bank detail: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete record"})
	}
	return c.JSON(fiber.Map{"message": "Record deleted successfully"})
}
