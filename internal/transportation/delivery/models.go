package delivery

import "time"

type Delivery struct {
	JobID             int        `json:"job_id"`
	TransporterID     int        `json:"transporter_id"`
	VehicleID         int        `json:"vehicle_id"`
	DriverID          int        `json:"driver_id"`
	AcceptedAt        *time.Time `json:"accepted_at"`
	PickupPhotoURL    *string    `json:"pickup_photo_url"`
	PickupConfirmedAt *time.Time `json:"pickup_confirmed_at"`
	DeliveryPhotoURL  *string    `json:"delivery_photo_url"`
	DeliveredAt       *time.Time `json:"delivered_at"`
	DeliveryStatus    string     `json:"delivery_status"`
}
