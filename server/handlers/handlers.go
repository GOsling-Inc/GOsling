package handlers

import (
	"github.com/GOsling-Inc/GOsling/services"
	"github.com/labstack/echo/v4"
)

type IAuthHandler interface {
	POST_SignUp(echo.Context) error
	POST_SignIn(echo.Context) error
	CreateJWT(string) (string, error)
}

type Handler struct {
	IAuthHandler
}

func New(s *services.Service) *Handler {
	return &Handler{
		IAuthHandler: NewAuthHandler(s),
	}
}
