package auth

import (
	"context"
	"farmerapp/go_backend/db"
	"fmt"

	"github.com/jackc/pgconn"
)

type AuthRepositoryInterface interface {
	RegisterUserUsingPhone(phoneNumber string, password string) error
	RegisterUserUsingEmail(email string, password string) error
	FindUserUsingPhone(phoneNumber string) (User, error)
	FindUserUsingEmail(email string) (User, error)
	StoreRefreshToken(user string, token string) error
	GetUserByRefreshToken(token string) (string, error)
}

type AuthRepository struct {
}

func (repository *AuthRepository) RegisterUserUsingPhone(phoneNumber string, password string) error {

	//  Checks for user existence and creates a new user if not exists
	// TODO: Check if fields like mobile number, email, etc. are unique
	var query string
	query = `INSERT INTO admin.user_table (mobile_num,password) VALUES ($1,$2)`

	_, err := db.Pool.Exec(context.Background(), query, phoneNumber, password)

	if err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			if pgErr.Code == "23505" {
				return fmt.Errorf("User with this email already exists")
			}
		}
		return fmt.Errorf("Failed to register user: %w", err)
	}
	return nil
}

func (repository *AuthRepository) RegisterUserUsingEmail(email string, password string) error {

	//  Checks for user exitence and creates a new user if not exists
	// TODO: Check if fields like mobile number, email, etc. are unique
	var query string
	query = `INSERT INTO admin.user_table (email,password) VALUES ($1,$2)`

	_, err := db.Pool.Exec(context.Background(), query, email, password)

	if err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			if pgErr.Code == "23505" {
				return fmt.Errorf("User with this email already exists")
			}
		}
		return fmt.Errorf("Failed to register user: %w", err)
	}

	return nil
}

func (repository *AuthRepository) FindUserUsingPhone(phoneNumber string) (User, error) {
	var user User
	var query string
	query = `
		SELECT
			mobile_num, password
		FROM
			admin.user_table
		WHERE
			mobile_num = $1
		`
	err := db.Pool.QueryRow(context.Background(), query, phoneNumber).
		Scan(&user.PhoneNumber, &user.Password)
	if err != nil {
		return user, fmt.Errorf("Query failed: %w", err)
	}

	return user, nil

}

func (repository *AuthRepository) FindUserUsingEmail(email string) (User, error) {
	var user User
	var query string
	query = `
		SELECT
			email, password
		FROM
			admin.user_table
		WHERE
			email = $1
		`
	err := db.Pool.QueryRow(context.Background(), query, email).
		Scan(&user.Email, &user.Password)
	if err != nil {
		return user, fmt.Errorf("Query failed: %w", err)
	}

	return user, nil

}

func (repository *AuthRepository) StoreRefreshToken(user string, token string) error {
	query := `INSERT INTO admin.refresh_tokens (user, token, expires_at) VALUES ($1, $2, NOW() + interval '7 days')`
	_, err := db.Pool.Exec(context.Background(), query, user, token)
	return err
}

func (repository *AuthRepository) GetUserByRefreshToken(token string) (string, error) {
	var user string
	query := `SELECT user FROM admin.refresh_tokens WHERE token = $1 AND expires_at > NOW()`
	err := db.Pool.QueryRow(context.Background(), query, token).Scan(&user)
	return user, err
}
