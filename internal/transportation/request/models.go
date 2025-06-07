package request

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
