package handlers

import (
	"context"
	"encoding/json"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/guregu/null"
	"github.com/PragaL15/go_newBackend/go_backend/db"
	"database/sql"
)

// Structs with json tags optimized for minimal size
type (
	CartDetails struct {
		CartID                 int64   `json:"cart_id"`
		RetailerID             int64   `json:"retailer_id"`
		RetailerName           string  `json:"retailer_name,omitempty"`
		RetailerAddress        string  `json:"retailer_address,omitempty"`
		RetailerStateName      string  `json:"retailer_state_name,omitempty"`
		RetailerStateShortname string  `json:"retailer_state_shortname,omitempty"`
		RetailerLocationName   string  `json:"retailer_location_name,omitempty"`
		WholesellerID          *int64  `json:"wholeseller_id,omitempty"`
		WholesellerName        string  `json:"wholeseller_name,omitempty"`
		CartStatus             int     `json:"cart_status"`
	}

	CartProduct struct {
		ProductID             int64       `json:"product_id"`
		ProductName           string      `json:"product_name"`
		Quantity              int         `json:"quantity"`
		UnitID                int64       `json:"unit_id"`
		UnitName              string      `json:"unit_name"`
		PriceWhileAdded       float64     `json:"price_while_added"`
		LatestWholesalerPrice float64     `json:"latest_wholesaler_price"`
		PriceUpdatedAt        null.String `json:"price_updated_at,omitempty"`
		IsActive              bool        `json:"is_active"`
	}

	CartInsert struct {
		ProductID             int64       `json:"product_id"`
		Quantity              int         `json:"quantity"`
		UnitID                int64       `json:"unit_id"`
		PriceWhileAdded       float64     `json:"price_while_added"`
		LatestWholesalerPrice float64     `json:"latest_wholesaler_price"`
		PriceUpdatedAt        null.String `json:"price_updated_at,omitempty"`
		WholesellerID         int64       `json:"wholeseller_id"`
		IsActive              bool        `json:"is_active"`
	}

	CartResponse struct {
		CartDetails CartDetails   `json:"cart_details,omitempty"`
		Products    []CartProduct `json:"products,omitempty"`
	}

	cartRequest struct {
		RetailerID    int64           `json:"retailer_id"`
		WholesellerID *int64          `json:"wholeseller_id,omitempty"`
		Products      []CartInsert    `json:"products"`
		DeviceInfo    json.RawMessage `json:"device_info,omitempty"`
		CartStatus    *int            `json:"cart_status,omitempty"`
	}
)

// Helper functions
func nullToString(ns sql.NullString) string {
	if ns.Valid {
		return ns.String
	}
	return ""
}

func errorResponse(c *fiber.Ctx, status int, message string) error {
	return c.Status(status).JSON(fiber.Map{
		"status":  "error",
		"message": message,
	})
}

func successResponse(c *fiber.Ctx, data interface{}) error {
	return c.JSON(fiber.Map{
		"status": "success",
		"data":   data,
	})
}

// GetCart retrieves cart details
func GetCart(c *fiber.Ctx) error {
	cartID, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return errorResponse(c, fiber.StatusBadRequest, "Invalid cart ID")
	}

	rows, err := db.Pool.Query(context.Background(), 
		"SELECT * FROM business_schema.get_cart_details($1)", cartID)
	if err != nil {
		log.Printf("Error fetching cart: %v", err)
		return errorResponse(c, fiber.StatusNotFound, "Cart not found")
	}
	defer rows.Close()

	var (
		response  CartResponse
		products  []CartProduct
		tempNulls struct {
			retailerName, retailerAddress, stateName, 
			shortname, locationName, wholesellerName sql.NullString
			wholesellerID sql.NullInt64
		}
	)

	for rows.Next() {
		var product CartProduct
		if err := rows.Scan(
			&response.CartDetails.CartID,
			&response.CartDetails.RetailerID,
			&tempNulls.retailerName,
			&tempNulls.retailerAddress,
			&tempNulls.stateName,
			&tempNulls.shortname,
			&tempNulls.locationName,
			&tempNulls.wholesellerID,
			&tempNulls.wholesellerName,
			&product.ProductID,
			&product.ProductName,
			&product.Quantity,
			&product.UnitID,
			&product.UnitName,
			&product.PriceWhileAdded,
			&product.LatestWholesalerPrice,
			&product.PriceUpdatedAt,
			&product.IsActive,
			&response.CartDetails.CartStatus,
		); err != nil {
			log.Printf("Error scanning row: %v", err)
			return errorResponse(c, fiber.StatusInternalServerError, "Failed to parse cart data")
		}

		// Assign nullable fields
		response.CartDetails.RetailerName = nullToString(tempNulls.retailerName)
		response.CartDetails.RetailerAddress = nullToString(tempNulls.retailerAddress)
		response.CartDetails.RetailerStateName = nullToString(tempNulls.stateName)
		response.CartDetails.RetailerStateShortname = nullToString(tempNulls.shortname)
		response.CartDetails.RetailerLocationName = nullToString(tempNulls.locationName)
		response.CartDetails.WholesellerName = nullToString(tempNulls.wholesellerName)
		if tempNulls.wholesellerID.Valid {
			id := tempNulls.wholesellerID.Int64
			response.CartDetails.WholesellerID = &id
		}

		products = append(products, product)
	}

	if len(products) == 0 {
		return errorResponse(c, fiber.StatusNotFound, "No cart found with that ID")
	}

	response.Products = products
	return successResponse(c, response)
}

// InsertCart creates a new cart
func InsertCart(c *fiber.Ctx) error {
	var req cartRequest
	if err := c.BodyParser(&req); err != nil {
		log.Printf("Error parsing body: %v", err)
		return errorResponse(c, fiber.StatusBadRequest, "Invalid request body")
	}

	if len(req.Products) == 0 {
		return errorResponse(c, fiber.StatusBadRequest, "At least one product required")
	}

	// Set default cart status if not provided
	cartStatus := 0
	if req.CartStatus != nil {
		cartStatus = *req.CartStatus
	}

	var resultJSON []byte
	err := db.Pool.QueryRow(
		context.Background(),
		"SELECT * FROM business_schema.insert_cart($1, $2, $3, $4, $5)",
		req.RetailerID,
		req.Products,
		req.WholesellerID,
		req.DeviceInfo,
		cartStatus,
	).Scan(&resultJSON)

	if err != nil {
		log.Printf("Error inserting cart: %v", err)
		return errorResponse(c, fiber.StatusInternalServerError, "Failed to create cart")
	}

	var result CartResponse
	if err := json.Unmarshal(resultJSON, &result); err != nil {
		log.Printf("Error parsing insert result: %v", err)
		return errorResponse(c, fiber.StatusInternalServerError, "Failed to process cart creation")
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": "success",
		"data":   result,
	})
}

// DeleteCartItem performs a soft delete of a cart item
func DeleteCartItem(c *fiber.Ctx) error {
	cartID, err := strconv.ParseInt(c.Params("cart_id"), 10, 64)
	if err != nil {
		return errorResponse(c, fiber.StatusBadRequest, "Invalid cart ID")
	}

	productID, err := strconv.ParseInt(c.Params("product_id"), 10, 64)
	if err != nil {
		return errorResponse(c, fiber.StatusBadRequest, "Invalid product ID")
	}

	var wholesellerID *int64
	if wholesellerParam := c.Query("wholeseller_id"); wholesellerParam != "" {
		if id, err := strconv.ParseInt(wholesellerParam, 10, 64); err == nil {
			wholesellerID = &id
		}
	}

	var resultJSON []byte
	err = db.Pool.QueryRow(
		context.Background(),
		"SELECT * FROM business_schema.soft_delete_cart_item($1, $2, $3)",
		cartID, productID, wholesellerID,
	).Scan(&resultJSON)

	if err != nil {
		log.Printf("Error soft deleting cart item: %v", err)
		return errorResponse(c, fiber.StatusInternalServerError, "Failed to delete cart item")
	}

	var result CartResponse
	if err := json.Unmarshal(resultJSON, &result); err != nil {
		log.Printf("Error parsing delete result: %v", err)
		return errorResponse(c, fiber.StatusInternalServerError, "Failed to process item deletion")
	}

	return successResponse(c, result)
}