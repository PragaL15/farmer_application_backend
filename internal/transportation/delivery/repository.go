package delivery

import (
	"context"
	"farmerapp/go_backend/db"
	"fmt"
)

type DeliveryRepositoryInterface interface {
	GetDeliveries(deliveryType string) ([]Delivery, error)
}

type DeliveryRepository struct {
}

func (repository *DeliveryRepository) GetDeliveries(deliveryType string) ([]Delivery, error) {
	// THOUGHT:- deliveryType can be an enum
	// Returning orders related to a certain user(transporter)
	var query string

	switch deliveryType {
	case "active":
		query = `SELECT * FROM transport_schema.trip_assignments WHERE delivery_status = 'active'`

	case "upcoming":
		query = `SELECT * FROM transport_schema.trip_assignments WHERE delivery_status = 'upcoming'`
	case "completed":
		query = `SELECT job_id,transporter_id FROM transport_schema.trip_assignments WHERE delivery_status = 'delivered'`
	default:
		return nil, fmt.Errorf("Invalid delivery type: %s", deliveryType)
	}
	rows, err := db.Pool.Query(context.Background(), query)
	if err != nil {
		return nil, fmt.Errorf("Query failed: %w", err)
	}
	defer rows.Close()

	var deliveries []Delivery
	for rows.Next() {
		var d Delivery
		err := rows.Scan(
			&d.JobID, &d.TransporterID,
		)
		if err != nil {
			return nil, fmt.Errorf("Row scan failed: %w", err)
		}
		deliveries = append(deliveries, d)
	}
	return deliveries, nil
}
