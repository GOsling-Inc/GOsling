package services

import (
	"errors"
	"log"

	"github.com/GOsling-Inc/GOsling/database"
	"github.com/GOsling-Inc/GOsling/models"
	"github.com/GOsling-Inc/GOsling/utils"
)

type IAuthService interface {
	SignIn(*models.User) error
	SignUp(*models.User) error

	TEST() error // DONT TOUCH
}

type AuthService struct {
	database *database.Database
	Utils    *utils.Utils
}

func NewAuthService(d *database.Database, u *utils.Utils) *AuthService {
	return &AuthService{
		database: d,
		Utils:    u,
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

func (s *AuthService) TEST() error { // DONT TOUCH
	l := models.Loan{
		AccountId: "ASXKukwHSunCRdBYN",
		UserId: "HSunCRd",
		Amount: 7000000,
		Remaining: 7000000,
		Part: 80000,
		Percent: 10,
		Period: "2030-04-06",
		Deadline: "2023-04-07",
	}
	err := s.database.AddLoan(l)
	log.Println(err)
	log.Println(s.database.GetUserLoans(l.UserId))
	return nil
}
