package services

import (
	"github.com/GOsling-Inc/GOsling/database"
	"github.com/GOsling-Inc/GOsling/models"
)

type IAuthService interface {
	SignIn(*models.User) error
	SignUp(*models.User) error
}

type IUserService interface {
	GetUser(string) (*models.User, error)
	Change_Main_Info(models.User) error
	Change_Password(models.User) error
	MakeID() string
	CreateJWT(string) (string, error)
	ParseJWT(string) (string, error)
	Validate(user *models.User) error
	Hash(str string) (string, error)
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
