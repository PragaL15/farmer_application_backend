package auth

import "time"

type IdentifierType int

const (
	IdentifierEmail IdentifierType = iota
	IdentifierPhone
)

type User struct {
<<<<<<< HEAD
	UserID    int `json:"user_id"`       // UUID
=======
	UserID    string `json:"user_id"`       // UUID
>>>>>>> 89c8837d6c691851780de9cd6d15506fa99fd796
	Name      string `json:"name"`          // Name of user
	MobileNum string `json:"mobile_num"`    // Phone number
	Email     string `json:"email"`         // Email address
	Address   string `json:"address"`       // Street address
	Pincode   string `json:"pincode"`       // Postal code
	Location  int    `json:"location"`      // City or locality
	State     int    `json:"state"`         // State
	RoleID    int    `json:"role_id"`       // Role: admin/retailer/etc.
	Password  string `json:"password"`      // Hashed password
	IsActive  int    `json:"active_status"` // Active status
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
