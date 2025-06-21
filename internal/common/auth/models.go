package auth

type IdentifierType int

const (
	IdentifierEmail IdentifierType = iota
	IdentifierPhone
)

type User struct {
	Email       string `json:"email"`
	PhoneNumber string `json:"mobile_num"`
	Password    string `json:"password"`
}
