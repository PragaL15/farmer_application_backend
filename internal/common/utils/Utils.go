package utils

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// IsPhoneNumber checks if the input looks like a phone number.
// It returns:
//   - true, nil       → if it's a valid Indian phone number
//   - true, error     → if it's a phone number but in incorrect format
//   - false, nil      → if it's not a phone number
func IsPhoneNumber(input string) (bool, error) {
	normalized := strings.ReplaceAll(input, " ", "")
	normalized = strings.TrimPrefix(normalized, "+91")
	normalized = strings.TrimPrefix(normalized, "91")

	digitsOnly := regexp.MustCompile(`^\d+$`).MatchString(normalized)
	if !digitsOnly {
		return false, nil
	}

	// Check for valid Indian mobile number format (must be 10 digits, starting with 6-9)
	if matched, _ := regexp.MatchString(`^[6-9]\d{9}$`, normalized); matched {
		return true, nil
	}

	return true, errors.New("Invalid phone number ")
}

func ValidateEmail(email string) error {
	const emailRegex = `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`

	re := regexp.MustCompile(emailRegex)
	if !re.MatchString(email) {
		return errors.New("Invalid email address")
	}
	return nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func GenerateAccessToken(userId string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userId,
		"exp":     time.Now().Add(5 * time.Minute).Unix(),
		"iat":     time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func ValidateJWT(tokenString string) (jwt.MapClaims, error) {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return nil, errors.New("JWT secret not set")
	}

	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		// Ensure token method is HMAC
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secret), nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("Invalid or expired token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid token claims")
	}

	return claims, nil
}

func GenerateRefreshToken() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}
