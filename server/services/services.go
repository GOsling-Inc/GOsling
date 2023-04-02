package services

import (
	"github.com/GOsling-Inc/GOsling/database"
	"github.com/GOsling-Inc/GOsling/models"
)

type IAuthService interface {
	SignIn(*models.User) error
	SignUp(*models.User) error

	TEST() error // DONT TOUCH
}

type IUserService interface {
	GetUser(string) (*models.User, error)
	Change_Main_Info(models.User) error
	Change_Password(models.User) error
	MakeID() string
	CreateJWT(string) (string, error)
	ParseJWT(string) (string, error)
	Validate(*models.User) error
	Hash(string) (string, error)
	AddAccount(*models.User, *models.Account) error
	GetUserAccounts(*models.User) ([]models.Account, error)
}

type Service struct {
	IAuthService
	IUserService
}

func New(d *database.Database) *Service {
	return &Service{
		IAuthService: NewAuthService(d),
		IUserService: NewUserService(d),
	}
}
