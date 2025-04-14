package handlers

import (
    "context"
    "encoding/json"
    "log"
    "strconv"
  "database/sql"
    "github.com/gofiber/fiber/v2"
    "github.com/PragaL15/go_newBackend/go_backend/db"
)

// Structs
type CartDetails struct {
	CartID                 int64  `json:"cart_id"`
	RetailerID             int64  `json:"retailer_id"`
	RetailerName           string `json:"retailer_name"`
	RetailerAddress        string `json:"retailer_address"`
	RetailerStateName      string `json:"retailer_state_name"`
	RetailerStateShortname string `json:"retailer_state_shortname"`
	RetailerLocationName   string `json:"retailer_location_name"`
	WholesellerID          *int64 `json:"wholeseller_id"`
	WholesellerName        string `json:"wholeseller_name"`
	CartStatus             int    `json:"cart_status"`
}

type CartProduct struct {
	ProductID             int64          `json:"product_id"`
	ProductName           string         `json:"product_name"`
	Quantity              int            `json:"quantity"`
	UnitID                int64          `json:"unit_id"`
	UnitName              string         `json:"unit_name"`
	PriceWhileAdded       float64        `json:"price_while_added"`
	LatestWholesalerPrice float64        `json:"latest_wholesaler_price"`
	PriceUpdatedAt        sql.NullString `json:"price_updated_at"`
	IsActive              bool           `json:"is_active"`
}

type CartResponse struct {
	CartDetails CartDetails   `json:"cart_details"`
	Products    []CartProduct `json:"products"`
}

func GetCart(c *fiber.Ctx) error {
	cartID := c.Params("id")

	rows, err := db.Pool.Query(
			context.Background(),
			"SELECT * FROM business_schema.get_cart_details($1)",
			cartID,
	)
	if err != nil {
			log.Printf("Error fetching cart: %v", err)
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
					"status":  "error",
					"message": "Cart not found",
			})
	}
	defer rows.Close()

	var response CartResponse
	var products []CartProduct

	// Nullable temp vars
	var retailerName, retailerAddress, stateName, shortname, locationName, wholesellerName sql.NullString
	var wholesellerID sql.NullInt64

	for rows.Next() {
			var product CartProduct
			err := rows.Scan(
					&response.CartDetails.CartID,
					&response.CartDetails.RetailerID,
					&retailerName,
					&retailerAddress,
					&stateName,
					&shortname,
					&locationName,
					&wholesellerID,
					&wholesellerName,
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
			)
			if err != nil {
					log.Printf("Error scanning row: %v", err)
					return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
							"status":  "error",
							"message": "Failed to parse cart data",
					})
			}

			// Assign nullable fields safely
			response.CartDetails.RetailerName = nullToString(retailerName)
			response.CartDetails.RetailerAddress = nullToString(retailerAddress)
			response.CartDetails.RetailerStateName = nullToString(stateName)
			response.CartDetails.RetailerStateShortname = nullToString(shortname)
			response.CartDetails.RetailerLocationName = nullToString(locationName)
			response.CartDetails.WholesellerName = nullToString(wholesellerName)
			if wholesellerID.Valid {
					id := wholesellerID.Int64
					response.CartDetails.WholesellerID = &id
			} else {
					response.CartDetails.WholesellerID = nil
			}

			products = append(products, product)
	}

	if len(products) == 0 {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
					"status":  "error",
					"message": "No cart found with that ID",
			})
	}

	response.Products = products
	return c.JSON(fiber.Map{
			"status":  "success",
			"message": "Cart retrieved successfully",
			"data":    response,
	})
}

func nullToString(ns sql.NullString) string {
	if ns.Valid {
			return ns.String
	}
	return ""
}
// Handler: InsertCart
func InsertCart(c *fiber.Ctx) error {
    var request struct {
        RetailerID    int64           `json:"retailer_id"`
        WholesellerID *int64          `json:"wholeseller_id,omitempty"`
        Products      json.RawMessage `json:"products"`
        DeviceInfo    json.RawMessage `json:"device_info,omitempty"`
    }

    if err := c.BodyParser(&request); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "status":  "error",
            "message": "Invalid request body",
        })
    }

    var products []CartProduct
    if err := json.Unmarshal(request.Products, &products); err != nil || len(products) == 0 {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "status":  "error",
            "message": "At least one valid product is required",
        })
    }

    var resultJSON []byte
    err := db.Pool.QueryRow(
        context.Background(),
        "SELECT * FROM business_schema.insert_cart($1, $2, $3, $4)",
        request.RetailerID,
        request.Products,
        request.WholesellerID,
        request.DeviceInfo,
    ).Scan(&resultJSON)

    if err != nil {
        log.Printf("Error inserting cart: %v", err)
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "status":  "error",
            "message": "Failed to create cart",
        })
    }

    var result CartResponse
    if err := json.Unmarshal(resultJSON, &result); err != nil {
        log.Printf("Error parsing insert result: %v", err)
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "status":  "error",
            "message": "Failed to process cart creation",
        })
    }

    return c.Status(fiber.StatusCreated).JSON(fiber.Map{
        "status":  "success",
        "message": "Cart created successfully",
        "data":    result,
    })
}

// Handler: DeleteCartItem
func DeleteCartItem(c *fiber.Ctx) error {
    cartID := c.Params("cart_id")
    productID := c.Params("product_id")

    cartIDInt, err := strconv.ParseInt(cartID, 10, 64)
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "status":  "error",
            "message": "Invalid cart ID",
        })
    }

    productIDInt, err := strconv.ParseInt(productID, 10, 64)
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "status":  "error",
            "message": "Invalid product ID",
        })
    }

    var wholesellerID *int64
    if wholesellerParam := c.Query("wholeseller_id"); wholesellerParam != "" {
        wholesellerIDInt, err := strconv.ParseInt(wholesellerParam, 10, 64)
        if err != nil {
            return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
                "status":  "error",
                "message": "Invalid wholeseller ID",
            })
        }
        wholesellerID = &wholesellerIDInt
    }

    var resultJSON []byte
    err = db.Pool.QueryRow(
        context.Background(),
        "SELECT * FROM business_schema.soft_delete_cart_item($1, $2, $3)",
        cartIDInt,
        productIDInt,
        wholesellerID,
    ).Scan(&resultJSON)

    if err != nil {
        log.Printf("Error soft deleting cart item: %v", err)
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "status":  "error",
            "message": "Failed to delete cart item",
        })
    }

    var result CartResponse
    if err := json.Unmarshal(resultJSON, &result); err != nil {
        log.Printf("Error parsing delete result: %v", err)
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "status":  "error",
            "message": "Failed to process item deletion",
        })
    }

    return c.JSON(fiber.Map{
        "status":  "success",
        "message": "Cart item deleted successfully",
        "data":    result,
    })
}
