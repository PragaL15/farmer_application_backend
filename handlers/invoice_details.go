package handlers

import (
	"context"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/PragaL15/go_newBackend/go_backend/db"
)

type Invoice struct {
	ID                    int       `json:"id"`
	InvoiceNumber         string    `json:"invoice_number"`
	OrderID               int       `json:"order_id"`
	TotalAmount           float64   `json:"total_amount"`
	RetailerID            int       `json:"retailer_id"`
	RetailerName          string    `json:"retailer_name"`
	RetailerEmail         string    `json:"retailer_email"`
	RetailerPhone         string    `json:"retailer_phone"`
	RetailerAddress       string    `json:"retailer_address"`
	RetailerLocationID    int       `json:"retailer_location_id"`
	RetailerStateID       int       `json:"retailer_state_id"`
	RetailerStateName     string    `json:"retailer_state_name"`
	RetailerLocationName  string    `json:"retailer_location_name"`
	WholesellerID         int       `json:"wholeseller_id"`
	WholesellerName       string    `json:"wholeseller_name"`
	WholesellerEmail      string    `json:"wholeseller_email"`
	WholesellerPhone      string    `json:"wholeseller_phone"`
	WholesellerAddress    string    `json:"wholeseller_address"`
	WholesellerLocationID int       `json:"wholeseller_location_id"`
	WholesellerStateID    int       `json:"wholeseller_state_id"`
	WholesellerStateName  string    `json:"wholeseller_state_name"`
	WholesellerLocationName string  `json:"wholeseller_location_name"`
	DiscountAmount        float64   `json:"discount_amount"`
	InvoiceDate           string    `json:"invoice_date"`
	DueDate               string    `json:"due_date"`
	PayMode               int       `json:"pay_mode"`
	PayModeName           string    `json:"pay_mode_name"`
	PayType               int       `json:"pay_type"`
	PayTypeName           string    `json:"pay_type_name"`
	FinalAmount           float64   `json:"final_amount"`
	StatusName            string    `json:"status_name"`
	Products              []Product `json:"products"`
}

type Product struct {
	OrderItemID int    `json:"order_item_id"`
	ProductID   int    `json:"product_id"`
	ProductName string `json:"product_name"`
}

func GetInvoiceDetails(c *fiber.Ctx) error {
	rows, err := db.Pool.Query(context.Background(), "SELECT * FROM get_invoice_details_with_business_info();")
	if err != nil {
		log.Printf("Failed to fetch invoice details: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch invoice details"})
	}
	defer rows.Close()

	var invoices []Invoice

	for rows.Next() {
		var invoice Invoice
		var product Product
		var invoiceDate, dueDate time.Time

		err := rows.Scan(
			&invoice.ID, &invoice.InvoiceNumber, &invoice.OrderID, 
			&invoice.RetailerID, &invoice.RetailerName, &invoice.RetailerEmail, 
			&invoice.RetailerPhone, &invoice.RetailerAddress, &invoice.RetailerLocationID, 
			&invoice.RetailerStateID, &invoice.WholesellerID, &invoice.WholesellerName, 
			&invoice.WholesellerEmail, &invoice.WholesellerPhone, &invoice.WholesellerAddress, 
			&invoice.WholesellerLocationID, &invoice.WholesellerStateID, 
			&invoice.TotalAmount, &invoice.DiscountAmount, &invoiceDate, 
			&dueDate, &invoice.PayMode, &invoice.PayModeName, 
			&invoice.PayType, &invoice.PayTypeName, &invoice.FinalAmount, 
			&product.OrderItemID, &product.ProductID, &product.ProductName, 
			&invoice.StatusName, &invoice.WholesellerStateName, &invoice.RetailerStateName, 
			&invoice.WholesellerLocationName, &invoice.RetailerLocationName,
		)
		if err != nil {
			log.Printf("Error scanning row: %v", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error processing data"})
		}

		invoice.InvoiceDate = invoiceDate.Format(time.RFC3339)
		invoice.DueDate = dueDate.Format("2006-01-02")

		found := false
		for i := range invoices {
			if invoices[i].ID == invoice.ID {
				invoices[i].Products = append(invoices[i].Products, product)
				found = true
				break
			}
		}
		if !found {
			invoice.Products = []Product{product}
			invoices = append(invoices, invoice)
		}
	}
	return c.JSON(invoices)
}
