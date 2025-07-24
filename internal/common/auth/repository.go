package auth

import (
	"context"
	"farmerapp/go_backend/db"
	"fmt"

	"github.com/jackc/pgconn"
)

type AuthRepositoryInterface interface {
	RegisterUserUsingPhone(phoneNumber string, password string, name string, address string, pincode string, location int, state int, roleID int, active int) error
	RegisterUserUsingEmail(email string, password string, name string, address string, pincode string, location int, state int, roleID int, active int) error
	FindUserUsingPhone(phoneNumber string) (User, error)
	FindUserUsingEmail(email string) (User, error)
	StoreRefreshToken(user int, token string) error
	GetUserByRefreshToken(token string) (int, error)
}

type AuthRepository struct {
}

func (repository *AuthRepository) RegisterUserUsingPhone(phoneNumber string, password string, name string, address string, pincode string, location int, state int, roleID int, active int) error {

	//  Checks for user existence and creates a new user if not exists
	// TODO: Check if fields like mobile number, email, etc. are unique
	var query string = ""
	query = `INSERT INTO admin_schema.user_table (mobile_num,password,name,address,pincode,location,state,role_id,active_status) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)`

	_, err := db.Pool.Exec(context.Background(), query, phoneNumber, password, name, address, pincode, location, state, roleID, active)

	if err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			if pgErr.Code == "23505" {
				return fmt.Errorf("User with this phone number already exists")
			}
		}
		return fmt.Errorf("Failed to register user: %w", err)
	}
	return nil
}

func (repository *AuthRepository) RegisterUserUsingEmail(email string, password string, name string, address string, pincode string, location int, state int, roleID int, active int) error {

	//  Checks for user exitence and creates a new user if not exists
	// TODO: Check if fields like mobile number, email, etc. are unique
	var query string
	query = `INSERT INTO admin_schema.user_table (email,password,name,address,pincode,location,state,role_id,active_status) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)`

	_, err := db.Pool.Exec(context.Background(), query, email, password, name, address, pincode, location, state, roleID, active)

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
			user_id,role_id, password
		FROM
			admin_schema.user_table
		WHERE
			mobile_num = $1
		`
	err := db.Pool.QueryRow(context.Background(), query, phoneNumber).
		Scan(&user.UserID, &user.RoleID, &user.Password)
	if err != nil {
		return user, fmt.Errorf("Query failed: %w", err)
	}

	return user, nil

}

func (repository *AuthRepository) FindUserUsingEmail(email string) (User, error) {
	var user User

	query := `
		SELECT
			user_id ,role_id , password
		FROM
			admin_schema.user_table
		WHERE
			email = $1
		`
	err := db.Pool.QueryRow(context.Background(), query, email).
		Scan(&user.UserID, &user.RoleID, &user.Password)
	if err != nil {
		return user, fmt.Errorf("Query failed: %w", err)
	}

	return user, nil

}

func (repository *AuthRepository) StoreRefreshToken(userId int, token string) error {
	query := `INSERT INTO admin_schema.refresh_tokens (user_id, token,issued_at, expires_at) VALUES ($1, $2,NOW(),NOW() + interval '7 days')`
	_, err := db.Pool.Exec(context.Background(), query, userId, token)
	return err
}

func (repository *AuthRepository) GetUserByRefreshToken(token string) (int, error) {
	var userId int
	query := `SELECT user_id FROM admin_schema.refresh_tokens WHERE token = $1 AND expires_at > NOW()`
	err := db.Pool.QueryRow(context.Background(), query, token).Scan(&userId)
	return userId, err
}
