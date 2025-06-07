package delivery

import (
	"context"
	"farmerapp/go_backend/db"
	"fmt"
)

type DeliveryRepositoryInterface interface {
	GetDeliveries(deliveryType string, transporterId int) ([]Delivery, error)
}

type DeliveryRepository struct {
}

func (repository *DeliveryRepository) GetDeliveries(deliveryType string, transporterId int) ([]Delivery, error) {
	// THOUGHT:- deliveryType can be an enum
	// Returning orders related to a certain user(transporter)
	var query string

	switch deliveryType {
	case "active":
		query = `

			SELECT
    				A.job_id,
    				B.pickup_address,
				B.drop_address
			FROM
    				transport_schema.trip_assignments A
			JOIN
    				transport_schema.transport_jobs B ON A.job_id = B.job_id
			WHERE
    				A.delivery_status = 'pending'
			AND
				A.transporter_id = $1
    			AND
				B.status = 'picked';
			`

	case "upcoming":
		query = `
			SELECT
    				A.job_id,
    				B.pickup_address,
				B.drop_address
			FROM
    				transport_schema.trip_assignments A
			JOIN
    				transport_schema.transport_jobs B ON A.job_id = B.job_id
			WHERE
    				A.delivery_status = 'pending'
			AND
				A.transporter_id = $1
			AND
				B.status = 'accepted';

			`
	case "completed":
		query = `
			SELECT
    				A.job_id,
    				B.pickup_address,
				B.drop_address
			FROM
    				transport_schema.trip_assignments A
			JOIN
    				transport_schema.transport_jobs B ON A.job_id = B.job_id
			WHERE
    				A.delivery_status = 'delivered'
			AND
				A.transporter_id = $1
    			AND
				B.status = 'delivered';

			`
	case "history":
		query = `
			SELECT
    				A.job_id,
    				B.pickup_address,
				B.drop_address
			FROM
    				transport_schema.trip_assignments A
			JOIN
    				transport_schema.transport_jobs B ON A.job_id = B.job_id
			WHERE
    				A.transporter_id = $1;

			`
	default:
		return nil, fmt.Errorf("Invalid delivery type: %s", deliveryType)
	}
	rows, err := db.Pool.Query(context.Background(), query, transporterId)
	if err != nil {
		return nil, fmt.Errorf("Query failed: %w", err)
	}
	defer rows.Close()

	var deliveries []Delivery
	for rows.Next() {
		var d Delivery
		err := rows.Scan(
			&d.JobID, &d.PickupAddress, &d.DropAddress,
		)
		if err != nil {
			return nil, fmt.Errorf("Row scan failed: %w", err)
		}
		deliveries = append(deliveries, d)
	}
	return deliveries, nil
}
