package request

import (
	"context"
	"farmerapp/go_backend/db"
	"fmt"
)

type RequestRepositoryInterface interface {
	GetTransportRequests(region string) ([]TransportRequest, error)
	AcceptTransportRequest(transporterId int, jobId int, vehicleId int, driverId int) error
	RejectTransportRequest(transporterId int, jobId int) error
}

type RequestRepository struct {
}

func (repository *RequestRepository) GetTransportRequests(region string) ([]TransportRequest, error) {
	//DISCUSSION:- What data to return and based on what(location served etc.)

	//  TODO:- Also check if request hasnt been already rejected by the transporter
	var query string
	query = `
			SELECT
    				job_id,
    				pickup_address,
				drop_address
			FROM
    				transport_schema.transport_jobs
			WHERE
    				status = 'open'
			AND
				pickup_address = $1;
			`

	rows, err := db.Pool.Query(context.Background(), query, region)
	if err != nil {
		return nil, fmt.Errorf("Query failed: %w", err)
	}
	defer rows.Close()

	var deliveryRequests []TransportRequest
	for rows.Next() {
		var d TransportRequest
		err := rows.Scan(
			&d.JobID, &d.PickupAddress, &d.DropAddress,
		)
		if err != nil {
			return nil, fmt.Errorf("Row scan failed: %w", err)
		}
		deliveryRequests = append(deliveryRequests, d)
	}
	return deliveryRequests, nil
}

func (repository *RequestRepository) AcceptTransportRequest(transporterId int, jobId int, vehicleId int, driverId int) error {
	//DISCUSSION:- Do we need data after its updated?
	tx, err := db.Pool.Begin(context.Background())
	if err != nil {
		return fmt.Errorf("Transaction begin failed: %w", err)
	}
	defer tx.Rollback(context.Background())

	updateJobStatusQuery := `
		UPDATE
			transport_schema.transport_jobs
		SET
			status = 'accepted',
		WHERE
			job_id = $1
		AND
			status = 'open';
		`
	_, err = tx.Exec(context.Background(), updateJobStatusQuery, jobId)

	if err != nil {
		return fmt.Errorf("Failed to update job status: %w", err)
	}

	insertTripAssignmentQuery := `
		INSERT INTO
			transport_schema.trip_assignments (job_id, transporter_id, vehicle_id, driver_id, delivery_status)
		VALUES
			($1, $2, $3, $4, 'pending')
		`
	_, err = tx.Exec(context.Background(), insertTripAssignmentQuery, jobId, transporterId, vehicleId, driverId)
	if err != nil {
		return fmt.Errorf("Failed to insert trip assignment: %w", err)
	}
	if err := tx.Commit(context.Background()); err != nil {
		return fmt.Errorf("Transaction commit failed: %w", err)
	}
	return nil
}

func (repository *RequestRepository) RejectTransportRequest(transporterId int, jobId int) error {
	// transport_schema.request_rejection is just a place holder table, actual may vary
	query := `INSERT INTO
			transport_schema.request_rejection (job_id, transporter_id)
		VALUES
			($1, $2)
		`
	_, err := db.Pool.Exec(context.Background(), query, jobId, transporterId)
	if err != nil {
		return fmt.Errorf("Failed to reject transport request: %w", err)
	}
	return nil
}
