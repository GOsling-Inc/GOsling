package router

import (
	"github.com/GOsling-Inc/GOsling/handlers"
	"github.com/labstack/echo/v4"
)

func Init(server *echo.Echo, handler *handlers.Handler) {
	server.POST("/sign-in", handler.POST_SignIn)
	server.POST("/sign-up", handler.POST_SignUp)
	server.GET("/exchanges", handler.GET_Exchanges)

	user := server.Group("/user")
	{
		user.POST("", handler.POST_User)
		user.POST("/change/main", handler.POST_Change_Main)
		user.POST("/change/password", handler.POST_Change_Password)
		user.POST("/new-account", handler.POST_Add_Account)
		user.POST("/accounts", handler.POST_User_Accounts)
		user.POST("/transfer", handler.POST_Transfer)
		user.POST("/exchange", handler.POST_User_Exchange) 
		user.POST("/loans" , handler.GET_User_Loans)
		user.POST("/new-loan", handler.POST_Loan)
	}
	server.POST("/TEST", handler.TEST) // DONT TOUCH
}
