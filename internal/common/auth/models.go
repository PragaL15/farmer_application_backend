package auth

import "time"

type IdentifierType int

const (
	IdentifierEmail IdentifierType = iota
	IdentifierPhone
)

type User struct {
	UserID      string `json:"user_id"` // UUID
	Email       string `json:"email"`
	PhoneNumber string `json:"mobile_num"`
	Password    string `json:"password"`
}

type RefreshToken struct {
	ID         int       `json:"id"`
	UserID     string    `json:"user_id"` // UUID
	Token      string    `json:"token"`
	DeviceInfo string    `json:"device_info"`
	IPAddress  string    `json:"ip_address"`
	IssuedAt   time.Time `json:"issued_at"`
	ExpiresAt  time.Time `json:"expires_at"`
	IsRevoked  bool      `json:"is_revoked"`
}
