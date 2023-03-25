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
	MakeID() string
	Validate(*models.User) error
	HashPassword(*models.User) error
	getHashedPassword(string) (string, error)
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
