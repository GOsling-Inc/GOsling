package services

import (
	"github.com/GOsling-Inc/GOsling/database"
	"github.com/GOsling-Inc/GOsling/utils"
)

type Service struct {
	IAuthService
	IUserService
	IAccountService
	*utils.Utils
}

func New(d *database.Database) *Service {
	u := utils.NewUtils(d);
	return &Service{
		IAuthService: NewAuthService(d, u),
		IUserService: NewUserService(d, u),
		IAccountService: NewAccountService(d, u),
		Utils: u,
	}
}
