package auth

import (
	"farmerapp/internal/common/utils"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type AuthServiceInterface interface {
	RegisterUser(identifier string, identifierType IdentifierType, password string, name string, address string, pincode string, location int, state int, roleID int, active int) error
	LoginUser(identifier string, identifierType IdentifierType, password string) (int, string, string, error)
	ValidateRefreshToken(token string) (string, error)
}

type AuthService struct {
	repository AuthRepositoryInterface
}

func NewAuthService(repository AuthRepositoryInterface) *AuthService {
	return &AuthService{repository: repository}
}

func (service *AuthService) RegisterUser(identifier string, identifierType IdentifierType, password string, name string, address string, pincode string, location int, state int, roleID int, active int) error {

	var err error

	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return fmt.Errorf("could not hash password: %w", err)
	}

	switch identifierType {
	case IdentifierEmail:
		err = service.repository.RegisterUserUsingEmail(identifier, hashedPassword, name, address, pincode, location, state, roleID, active)

	case IdentifierPhone:
		err = service.repository.RegisterUserUsingPhone(identifier, hashedPassword, name, address, pincode, location, state, roleID, active)

	default:
		return fmt.Errorf("invalid identifier type: %d", identifierType)
	}
	if err != nil {
		return fmt.Errorf("service failed to register user: %w", err)
	}
	return nil

}

func (service *AuthService) LoginUser(identifier string, identifierType IdentifierType, password string) (int, string, string, error) {
	var user User
	var err error

	switch identifierType {
	case IdentifierEmail:
		user, err = service.repository.FindUserUsingEmail(identifier)
	case IdentifierPhone:
		user, err = service.repository.FindUserUsingPhone(identifier)
	default:
		return -1, "", "", fmt.Errorf("invalid identifier type")
	}
    
	if err != nil {
		return -1, "", "", err
	}

	// Check hashed password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return -1, "", "", fmt.Errorf("invalid credentials")
	}

	// Generate JWT
	accessToken, err := utils.GenerateAccessToken(user.UserID)
	if err != nil {
		return -1, "", "", fmt.Errorf("could not generate token: %w", err)
	}

	refreshToken, err := utils.GenerateRefreshToken()
	if err != nil {
		return -1, "", "", err
	}

	err = service.repository.StoreRefreshToken(user.UserID, refreshToken)
	if err != nil {
		return -1, "", "", err
	}

	// return accessToken, refreshToken, nil

	return user.RoleID, accessToken, refreshToken, nil
}

func (service *AuthService) ValidateRefreshToken(token string) (string, error) {
	return service.repository.GetUserByRefreshToken(token)
}
