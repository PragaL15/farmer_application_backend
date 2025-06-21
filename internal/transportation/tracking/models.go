package tracking

import "time"

type ValidateQrCode struct {
	JobID         int    `json:"job_id"`
	TransporterID int    `json:"transporter_id"`
	QrCode        string `json:"qr_code"`
}

type QRCode struct {
	QRID              string    `json:"qr_id"` // UUID
	IssueDate         time.Time `json:"issue_date"`
	Status            string    `json:"status"` // Enum: unused, sealed, etc.
	LastUpdatedAt     time.Time `json:"last_updated_at"`
	LastKnownLocation string    `json:"last_known_location"`
}

type QRScanEvent struct {
	ScanID        int       `json:"scan_id"`
	QRID          string    `json:"qr_id"`           // UUID
	ScannedByUser string    `json:"scanned_by_user"` // UUID
	ScanType      string    `json:"scan_type"`       // Enum
	ScannedAt     time.Time `json:"scanned_at"`
	Location      string    `json:"location"`
	Remarks       string    `json:"remarks"`
}

type OTPVerification struct {
	OTPID          int       `json:"otp_id"`
	ReferenceID    string    `json:"reference_id"` // UUID
	OTPCode        string    `json:"otp_code"`
	CreatedAt      time.Time `json:"created_at"`
	ExpiresAt      time.Time `json:"expires_at"`
	Verified       bool      `json:"verified"`
	DeliveryMethod string    `json:"delivery_method"` // Enum: in_app, sms, both
}

type QROrderLink struct {
	LinkID      int    `json:"link_id"`
	QRID        string `json:"qr_id"`    // UUID
	OrderID     string `json:"order_id"` // UUID
	LinkedOTPID int    `json:"linked_otp_id"`
}
