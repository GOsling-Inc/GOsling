package router

import (
	"github.com/GOsling-Inc/GOsling/handlers"
	"github.com/labstack/echo/v4"
)

func Init(server *echo.Echo, handler *handlers.Handler) {
	server.POST("/sign-in", handler.POST_SignIn)
	server.POST("/sign-up", handler.POST_SignUp)
}
