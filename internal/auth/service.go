package auth

import (
	"golang.org/x/crypto/bcrypt"
	"errors"
	"lowerkamacase/golang/internal/user"
)

type AuthService struct {
	UserRepository *user.UserRepository
}

func NewAuthService(userRepository *user.UserRepository) *AuthService {
	return &AuthService{UserRepository: userRepository}
}

func (service *AuthService) Register(email, password, name string) (string, error) {
	existingUser, _ := service.UserRepository.FindByEmail(email)
	if existingUser != nil {
		return "", errors.New("ErrUserExists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	user := &user.User{
		Email: email,
		Password: string(hashedPassword),		
		Name: name,
	}

	_, err = service.UserRepository.Create(user)
	if err != nil {
		return "", err
	}

	return user.Email, nil

}