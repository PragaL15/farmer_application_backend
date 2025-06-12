package tracking

type ValidateQrCode struct {
	JobID         int    `json:"job_id"`
	TransporterID int    `json:"transporter_id"`
	QrCode        string `json:"qr_code"`
}
