package handlers

import (
	"github.com/GOsling-Inc/GOsling/services"
	"github.com/labstack/echo/v4"
)

type IAuthHandler interface {
	POST_SignUp(echo.Context) error
	POST_SignIn(echo.Context) error

	TEST(echo.Context) error // DONT TOUCH
}

type IUserHandler interface {
	POST_User(echo.Context) error
	POST_Change_Main(c echo.Context) error
	POST_Change_Password(c echo.Context) error
}

type Handler struct {
	IAuthHandler
	IUserHandler
}

func New(s *services.Service) *Handler {
	return &Handler{
		IAuthHandler: NewAuthHandler(s),
		IUserHandler: NewUserHandler(s),
	}
}
