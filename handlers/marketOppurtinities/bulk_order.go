package marketoppurtinities

import (
	"context"
	"fmt"
	"time"
	"github.com/jackc/pgx/v4"
)

type OrderDetail struct {
	OrderID          int     `json:"order_id"`
	DateOfOrder      time.Time `json:"date_of_order"`
	TotalOrderAmount float64   `json:"total_order_amount"`
	FinalAmount      float64   `json:"final_amount"`
	ProductName      string    `json:"product_name"`
	WholesellerName  string    `json:"wholeseller_name"`
	RetailerName     string    `json:"retailer_name"`
}

func getAllOrderDetails(conn *pgx.Conn) ([]OrderDetail, error) {
	rows, err := conn.Query(
		context.Background(),
		"SELECT * FROM business_schema.get_all_order_details()",
	)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %v", err)
	}
	defer rows.Close()

	var orderDetails []OrderDetail
	for rows.Next() {
		var od OrderDetail
		if err := rows.Scan(
			&od.OrderID,
			&od.DateOfOrder,
			&od.TotalOrderAmount,
			&od.FinalAmount,
			&od.ProductName,
			&od.WholesellerName,
			&od.RetailerName,
		); err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}
		orderDetails = append(orderDetails, od)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %v", err)
	}
	return orderDetails, nil
}