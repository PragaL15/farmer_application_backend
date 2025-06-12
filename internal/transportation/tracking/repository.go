package tracking

import (
	"context"
	"farmerapp/go_backend/db"
	"fmt"
)

type TrackingRepositoryInterface interface {
	ValidateQrCode(jobid int, qrCode string, transporterId int) error
}

type TrackingRepository struct {
}

func (repository *TrackingRepository) ValidateQrCode(jobid int, qrCode string, transporterId int) error {
	var query string
	// TODO:- Table used is just a place holder, replace with actual table name
	query = `SELECT EXISTS (
				SELECT
					1
				FROM
					transport_schema.trip_assignments
				WHERE
					job_id = $1
				AND
					transporter_id = $2
				AND
					qr_code = $3
			)`
	var exists bool
	// Check if the QR code is valid for the given job and transporter
	err := db.Pool.QueryRow(context.Background(), query, jobid, transporterId, qrCode).Scan(&exists)
	if err != nil {
		return fmt.Errorf("Query failed: %w", err)
	}

	if !exists {
		return fmt.Errorf("Scanned QR code is either invalid or incorrect")
	}
	return nil

}
