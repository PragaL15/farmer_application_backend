package tests

import (
	"context"
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

type OrderStatus struct {
	OrderStatusID int
	OrderStatus   string
}

func fetchOrderStatusesFromDB(ctx context.Context, db *sql.DB) ([]OrderStatus, error) {
	query := "SELECT order_status_id, order_status FROM order_status_table"
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orderStatuses []OrderStatus
	for rows.Next() {
		var os OrderStatus
		if err := rows.Scan(&os.OrderStatusID, &os.OrderStatus); err != nil {
			return nil, err
		}
		orderStatuses = append(orderStatuses, os)
	}
	return orderStatuses, nil
}
func TestGetOrderStatuses(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"order_status_id", "order_status"}).
		AddRow(1, "Processing").
		AddRow(2, "Confirmed").
		AddRow(3, "Payment").
		AddRow(4, "Out for Delivery").
		AddRow(5, "Successful").
		AddRow(6, "Cancellation").
		AddRow(7, "Returned")

	mock.ExpectQuery(`SELECT order_status_id, order_status FROM order_status_table`).
		WillReturnRows(rows)

	orderStatuses, err := fetchOrderStatusesFromDB(context.Background(), db)
	assert.NoError(t, err)
	assert.Len(t, orderStatuses, 7)

	assert.Equal(t, "Processing", orderStatuses[0].OrderStatus)
	assert.Equal(t, "Confirmed", orderStatuses[1].OrderStatus)
	assert.NoError(t, mock.ExpectationsWereMet())
}
