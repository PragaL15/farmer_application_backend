package handlers

import (
    "context"
    "encoding/json"
    "log"
    "strconv"
    "time"

    "github.com/gofiber/fiber/v2"
   "github.com/PragaL15/go_newBackend/go_backend/db"
)

// Structs
type Cart struct {
    CartID         int64           `json:"cart_id"`
    RetailerID     int64           `json:"retailer_id"`
    WholesellerID  *int64          `json:"wholeseller_id,omitempty"`
    Products       json.RawMessage `json:"products"`
    DeviceInfo     json.RawMessage `json:"device_info,omitempty"`
    CartStatus     int             `json:"cart_status"`
    CreatedAt      time.Time       `json:"created_at"`
    UpdatedAt      time.Time       `json:"updated_at"`
}

type CartProduct struct {
    ProductID             int64     `json:"product_id"`
    Quantity              float64   `json:"quantity"`
    UnitID                int       `json:"unit_id"`
    PriceWhileAdded       float64   `json:"price_while_added"`
    LatestWholesalerPrice float64   `json:"latest_wholesaler_price"`
    PriceUpdatedAt        time.Time `json:"price_updated_at"`
    WholesellerID         *int64    `json:"wholeseller_id,omitempty"`
    IsActive              bool      `json:"is_active"`
}

type CartResponse struct {
    Status  string      `json:"status"`
    Message string      `json:"message"`
    Data    interface{} `json:"data,omitempty"`
}

// Handler: GetCart
func GetCart(c *fiber.Ctx) error {
    cartID := c.Params("id")

    var cart Cart
    var productsJSON []byte

    err := db.Pool.QueryRow(
        context.Background(),
        "SELECT * FROM business_schema.get_cart_details($1)",
        cartID,
    ).Scan(
        &cart.CartID,
        &cart.RetailerID,
        &cart.WholesellerID,
        &productsJSON,
        &cart.CartStatus,
        &cart.CreatedAt,
        &cart.UpdatedAt,
    )

    if err != nil {
        log.Printf("Error fetching cart: %v", err)
        return c.Status(fiber.StatusNotFound).JSON(CartResponse{
            Status:  "error",
            Message: "Cart not found",
        })
    }

    var products []CartProduct
    if err := json.Unmarshal(productsJSON, &products); err != nil {
        log.Printf("Error parsing products: %v", err)
        return c.Status(fiber.StatusInternalServerError).JSON(CartResponse{
            Status:  "error",
            Message: "Failed to parse cart products",
        })
    }

    return c.JSON(CartResponse{
        Status:  "success",
        Message: "Cart retrieved successfully",
        Data: fiber.Map{
            "cart":     cart,
            "products": products,
        },
    })
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
        return c.Status(fiber.StatusBadRequest).JSON(CartResponse{
            Status:  "error",
            Message: "Invalid request body",
        })
    }

    var products []CartProduct
    if err := json.Unmarshal(request.Products, &products); err != nil || len(products) == 0 {
        return c.Status(fiber.StatusBadRequest).JSON(CartResponse{
            Status:  "error",
            Message: "At least one valid product is required",
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
        return c.Status(fiber.StatusInternalServerError).JSON(CartResponse{
            Status:  "error",
            Message: "Failed to create cart",
        })
    }

    var result CartResponse
    if err := json.Unmarshal(resultJSON, &result); err != nil {
        log.Printf("Error parsing insert result: %v", err)
        return c.Status(fiber.StatusInternalServerError).JSON(CartResponse{
            Status:  "error",
            Message: "Failed to process cart creation",
        })
    }

    return c.Status(fiber.StatusCreated).JSON(result)
}

// Handler: DeleteCartItem
func DeleteCartItem(c *fiber.Ctx) error {
    cartID := c.Params("cart_id")
    productID := c.Params("product_id")

    cartIDInt, err := strconv.ParseInt(cartID, 10, 64)
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(CartResponse{
            Status:  "error",
            Message: "Invalid cart ID",
        })
    }

    productIDInt, err := strconv.ParseInt(productID, 10, 64)
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(CartResponse{
            Status:  "error",
            Message: "Invalid product ID",
        })
    }

    var wholesellerID *int64
    if wholesellerParam := c.Query("wholeseller_id"); wholesellerParam != "" {
        wholesellerIDInt, err := strconv.ParseInt(wholesellerParam, 10, 64)
        if err != nil {
            return c.Status(fiber.StatusBadRequest).JSON(CartResponse{
                Status:  "error",
                Message: "Invalid wholeseller ID",
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
        return c.Status(fiber.StatusInternalServerError).JSON(CartResponse{
            Status:  "error",
            Message: "Failed to delete cart item",
        })
    }

    var result CartResponse
    if err := json.Unmarshal(resultJSON, &result); err != nil {
        log.Printf("Error parsing delete result: %v", err)
        return c.Status(fiber.StatusInternalServerError).JSON(CartResponse{
            Status:  "error",
            Message: "Failed to process item deletion",
        })
    }

    return c.JSON(result)
}
