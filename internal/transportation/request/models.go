package request

import "time"

type TransportRequest struct {
	JobID         int    `json:"job_id"`
	PickupAddress string `json:"pickup_address"`
	DropAddress   string `json:"drop_address"`
}

type AcceptJobRequest struct {
	JobID         int `json:"job_id"`
	TransporterID int `json:"transporter_id"`
	VehicleID     int `json:"vehicle_id"`
	DriverID      int `json:"driver_id"`
}

type RejectJobRequest struct {
	JobID         int `json:"job_id"`
	TransporterID int `json:"transporter_id"`
}

type OrderRejection struct {
	ID              int       `json:"id"`
	OrderID         string    `json:"order_id"`  // UUID
	UserID          string    `json:"user_id"`   // UUID
	UserRole        string    `json:"user_role"` // Enum: aadthi, transporter, retailer
	ReasonID        int       `json:"reason_id"`
	OtherReasonText string    `json:"other_reason_text"`
	Timestamp       time.Time `json:"timestamp"`
	RetryAllowed    bool      `json:"retry_allowed"`
	RejectionSource string    `json:"rejection_source"` // Enum: manual, auto
}

type RejectionReason struct {
	ID           int      `json:"id"`
	ReasonText   string   `json:"reason_text"`
	ApplicableTo []string `json:"applicable_to"` // Roles array
	IsActive     bool     `json:"is_active"`
}
