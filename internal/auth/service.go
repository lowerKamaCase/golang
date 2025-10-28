package auth

import (
	"errors"
	"fmt"
	"lowerkamacase/golang/internal/user"

	"golang.org/x/crypto/bcrypt"
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
		return "", errors.New(ErrUserExists)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	user := &user.User{
		Email:    email,
		Password: string(hashedPassword),
		Name:     name,
	}

	_, err = service.UserRepository.Create(user)
	if err != nil {
		return "", err
	}

	return user.Email, nil
}

func (service *AuthService) Login(email, password string) (string, error) {
	existingUser, _ := service.UserRepository.FindByEmail(email)
	if existingUser == nil {
		return "", errors.New(ErrWrongCredentials)
	}

	fmt.Println("existingUser.Password = ", existingUser.Password)

	err := bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(password))
	if err != nil {
		return "", err
	}

	return email, nil
}
