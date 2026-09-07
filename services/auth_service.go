package services

import (
	"errors"
	"weather-api/models"
	"weather-api/repositories"
	"weather-api/utils"
)

var ErrInvalidCredentials = errors.New("invalid email or password")

type AuthService struct {
	userRepo *repositories.UserRepository
}

func NewAuthService(userRepo *repositories.UserRepository) *AuthService {
	return &AuthService{userRepo: userRepo}
}

func (s *AuthService) Signup(user *models.User) error {
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword

	return s.userRepo.Create(user)
}

func (s *AuthService) Login(email, password string) (*models.User, string, error) {
	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		return nil, "", ErrInvalidCredentials
	}

	if !utils.CheckPasswordHash(password, user.Password) {
		return nil, "", ErrInvalidCredentials
	}

	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		return nil, "", err
	}

	user.Password = ""
	return user, token, nil
}
