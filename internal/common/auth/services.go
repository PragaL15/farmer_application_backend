package auth

import (
	"farmerapp/internal/common/utils"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type AuthServiceInterface interface {
	RegisterUser(identifier string, identifierType IdentifierType, password string, name string, address string, pincode string, location int, state int, roleID int, active int) error
<<<<<<< HEAD
	LoginUser(identifier string, identifierType IdentifierType, password string) (int, int, string, string, error)
	ValidateRefreshToken(token string) (int, error)
=======
	LoginUser(identifier string, identifierType IdentifierType, password string) (int, string, string, error)
	ValidateRefreshToken(token string) (string, error)
>>>>>>> 89c8837d6c691851780de9cd6d15506fa99fd796
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
<<<<<<< HEAD

	case IdentifierPhone:
		err = service.repository.RegisterUserUsingPhone(identifier, hashedPassword, name, address, pincode, location, state, roleID, active)

=======
	case IdentifierPhone:
		err = service.repository.RegisterUserUsingPhone(identifier, hashedPassword, name, address, pincode, location, state, roleID, active)
>>>>>>> 89c8837d6c691851780de9cd6d15506fa99fd796
	default:
		return fmt.Errorf("invalid identifier type: %d", identifierType)
	}
	if err != nil {
		return fmt.Errorf("service failed to register user: %w", err)
	}
	return nil

}

<<<<<<< HEAD
func (service *AuthService) LoginUser(identifier string, identifierType IdentifierType, password string) (int, int, string, string, error) {
=======
func (service *AuthService) LoginUser(identifier string, identifierType IdentifierType, password string) (int, string, string, error) {
>>>>>>> 89c8837d6c691851780de9cd6d15506fa99fd796
	var user User
	var err error

	switch identifierType {
	case IdentifierEmail:
		user, err = service.repository.FindUserUsingEmail(identifier)
	case IdentifierPhone:
		user, err = service.repository.FindUserUsingPhone(identifier)
	default:
<<<<<<< HEAD
		return -1, -1, "", "", fmt.Errorf("invalid identifier type")
	}

	if err != nil {
		return -1, -1, "", "", err
=======
		return -1, "", "", fmt.Errorf("invalid identifier type")
	}

	if err != nil {
		return -1, "", "", err
>>>>>>> 89c8837d6c691851780de9cd6d15506fa99fd796
	}

	// Check hashed password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
<<<<<<< HEAD
		return -1, -1, "", "", fmt.Errorf("invalid credentials")
=======
		return -1, "", "", fmt.Errorf("invalid credentials")
>>>>>>> 89c8837d6c691851780de9cd6d15506fa99fd796
	}

	// Generate JWT
	accessToken, err := utils.GenerateAccessToken(user.UserID)
	if err != nil {
<<<<<<< HEAD
		return -1, -1, "", "", fmt.Errorf("could not generate token: %w", err)
=======
		return -1, "", "", fmt.Errorf("could not generate token: %w", err)
>>>>>>> 89c8837d6c691851780de9cd6d15506fa99fd796
	}

	refreshToken, err := utils.GenerateRefreshToken()
	if err != nil {
<<<<<<< HEAD
		return -1, -1, "", "", err
=======
		return -1, "", "", err
>>>>>>> 89c8837d6c691851780de9cd6d15506fa99fd796
	}

	err = service.repository.StoreRefreshToken(user.UserID, refreshToken)
	if err != nil {
<<<<<<< HEAD
		return -1, -1, "", "", err
=======
		return -1, "", "", err
>>>>>>> 89c8837d6c691851780de9cd6d15506fa99fd796
	}

	// return accessToken, refreshToken, nil

<<<<<<< HEAD
	return user.UserID, user.RoleID, accessToken, refreshToken, nil
=======
	return user.RoleID, accessToken, refreshToken, nil
>>>>>>> 89c8837d6c691851780de9cd6d15506fa99fd796
}

func (service *AuthService) ValidateRefreshToken(token string) (int, error) {
	return service.repository.GetUserByRefreshToken(token)
}
