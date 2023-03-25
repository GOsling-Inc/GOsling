package services

import (
	"errors"

	"github.com/GOsling-Inc/GOsling/database"
	"github.com/GOsling-Inc/GOsling/models"
)

type AuthService struct {
	database *database.Database
}

func NewAuthService(d *database.Database) *AuthService {
	return &AuthService{
		database: d,
	}
}

func (s *AuthService) SignIn(user *models.User) error {
	tempUser, err := s.database.GetUserByMail(user.Email)
	if err != nil {
		return errors.New("incorrect email or password")
	}
	if user.Password != tempUser.Password {
		return errors.New("incorrect email or password")
	}
	return nil
}

func (s *AuthService) SignUp(user *models.User) error {
	_, err := s.database.GetUserByMail(user.Email)
	if err == nil {
		return errors.New("user with this email already registered")
	}

	err = s.database.AddUser(user)
	return err
}
