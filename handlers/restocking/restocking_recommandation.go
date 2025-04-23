package RestockingStock

import (
    "context"
    "log"
    "net/http"

    "github.com/gofiber/fiber/v2"
    "github.com/PragaL15/go_newBackend/go_backend/db"
)

type LowStockProduct struct {
    ProductID   int      `json:"product_id"`
    ProductName string   `json:"product_name"`
    Mandi       []Mandi `json:"mandi"`
}
type Mandi struct {
    MandiID    int     `json:"mandi_id"`
    MandiName  string  `json:"mandi_name"`
    MandiStock float64 `json:"mandi_stock"`
}

func GetLowStockProductsHandler(c *fiber.Ctx) error {
    query := "SELECT * FROM business_schema.get_low_stock_products();"
    rows, err := db.Pool.Query(context.Background(), query)
    if err != nil {
        log.Println("Error executing low stock product query:", err)
        return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
            "error":  "Database query failed",
            "detail": err.Error(),
        })
    }
    defer rows.Close()

    var lowStockProducts []LowStockProduct
    for rows.Next() {
        var productID int
        var productName string
        var mandiData []Mandi

        if err := rows.Scan(&productID, &productName, &mandiData); err != nil {
            log.Println("Error scanning low stock product row:", err)
            return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
                "error":  "Failed to scan row",
                "detail": err.Error(),
            })
        }

        lowStockProducts = append(lowStockProducts, LowStockProduct{
            ProductID:   productID,
            ProductName: productName,
            Mandi:       mandiData,
        })
    }

    return c.Status(http.StatusOK).JSON(lowStockProducts)
}
