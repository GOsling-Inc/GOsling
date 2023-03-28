package router

import (
	"github.com/GOsling-Inc/GOsling/handlers"
	"github.com/labstack/echo/v4"
)

func Init(server *echo.Echo, handler *handlers.Handler) {
	server.POST("/sign-in", handler.POST_SignIn)
	server.POST("/sign-up", handler.POST_SignUp)
	server.POST("/user", handler.POST_User)
	server.POST("user/change/main", handler.POST_Change_Main)
	server.POST("/user/change/password", handler.POST_Change_Password)

	/*
		TODO:
		server.POST("/user", handler.POST_User)    -   get information about user
		server.POST("/user/change", handler.POST_Change)  -  change user field (name, surname, password, email)
		user sends html form with name, surname, email and password
		firstly, parse jwt from header to get user's mail
	*/
}
