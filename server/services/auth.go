package services

import (
	"errors"
	"log"

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
	user.Id = tempUser.Id
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

func (s *AuthService) TEST() error { // DONT TOUCH (PROTO?)
	acc := new(models.Account)
	acc.Id = "x4erBWbf08RkW3R41"
	acc.Name = "Main"
	acc.Type = "BASIC"
	acc.Unit = "BYN"
	acc.UserId = "x4erBWb"
	if err := s.database.AddAccount(acc); err != nil {
		log.Println(err)
		return err
	}
	accounts, err := s.database.GetUserAccounts("x4erBWb")
	log.Println(accounts, err)
	return err
}
