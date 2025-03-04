package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/pashagolub/pgxmock"
	"github.com/stretchr/testify/assert"

	"github.com/PragaL15/go_newBackend/go_backend/db"
handlers	"github.com/PragaL15/go_newBackend/handlers"
)

func TestGetInvoiceDetails(t *testing.T) {
	app := fiber.New()
	
	mockPool, err := pgxmock.NewPool()
	assert.NoError(t, err)
	db.Pool = mockPool 
	invoiceDate := time.Now()
	dueDate := invoiceDate.AddDate(0, 0, 30)

	rows := mockPool.NewRows([]string{
		"id", "invoice_number", "order_id", "retailer_id", "retailer_name", "retailer_email",
		"retailer_phone", "retailer_address", "retailer_location_id", "retailer_state_id",
		"wholeseller_id", "wholeseller_name", "wholeseller_email", "wholeseller_phone",
		"wholeseller_address", "wholeseller_location_id", "wholeseller_state_id",
		"total_amount", "discount_amount", "invoice_date", "due_date", "pay_mode", "pay_mode_name",
		"pay_type", "pay_type_name", "final_amount", "order_item_id", "product_id", "product_name",
		"status_name", "wholeseller_state_name", "retailer_state_name", "wholeseller_location_name", "retailer_location_name",
	}).AddRow(
		1, "INV-1001", 101, 201, "Retailer A", "retailer@example.com",
		"9876543210", "Retailer Address", 301, 401,
		501, "Wholesaler A", "wholesaler@example.com", "9876543211",
		"Wholesaler Address", 601, 701,
		1000.00, 50.00, invoiceDate, dueDate, 1, "Credit Card",
		2, "Online", 950.00, 1001, 2001, "Product X",
		"Pending", "Wholesaler State", "Retailer State", "Wholesaler Location", "Retailer Location",
	)

	mockPool.ExpectQuery(`(?i)SELECT\s+\*\s+FROM\s+get_invoice_details_with_business_info\s*\(\);?`).WillReturnRows(rows)

	app.Get("/invoice", handlers.GetInvoiceDetails)

	req := httptest.NewRequest(http.MethodGet, "/invoice", nil)
	resp, err := app.Test(req, -1)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var response []handlers.Invoice
	err = json.NewDecoder(resp.Body).Decode(&response)
	assert.NoError(t, err)

	assert.Len(t, response, 1)
	assert.Equal(t, "INV-1001", response[0].InvoiceNumber)
	assert.Len(t, response[0].Products, 1)
	assert.Equal(t, "Product X", response[0].Products[0].ProductName)
	assert.NoError(t, mockPool.ExpectationsWereMet())

}
