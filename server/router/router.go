package router

import (
	"github.com/GOsling-Inc/GOsling/handlers"
	"github.com/labstack/echo/v4"
)

func Init(server *echo.Echo, handler *handlers.Handler) {
	server.POST("/sign-in", handler.POST_SignIn)
	server.POST("/sign-up", handler.POST_SignUp)
	server.POST("/user", handler.POST_User)
	server.POST("/user/change/main", handler.POST_Change_Main)
	server.POST("/user/change/password", handler.POST_Change_Password)

	server.POST("/user/addaccount", handler.POST_Add_Account)
	server.POST("/user/accounts", handler.POST_User_Accounts)

	server.POST("user/transfer", handler.POST_Transfer) // (beta)

	server.POST("user/exchange", handler.POST_User_Exchange) // (beta)

	server.POST("/TEST", handler.TEST) // DONT TOUCH
}
