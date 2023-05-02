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
		user.POST("/delete-account", handler.POST_Delete_Account)
		user.GET("/accounts", handler.GET_User_Accounts)

		user.POST("/transfer", handler.POST_Transfer)

		user.POST("/exchange", handler.POST_User_Exchange)

		user.GET("/loans", handler.GET_User_Loans)
		user.POST("/new-loan", handler.POST_Loan)

		user.GET("/deposits", handler.GET_User_Deposits)
		user.POST("/new-deposit", handler.POST_NewDeposit)

		user.GET("/insurances", handler.GET_User_Insurances)
		user.POST("/new-insurance", handler.POST_NewInsurance)

		user.POST("/stocks/new-order", handler.POST_User_Stocks_NewOrder)
		user.POST("/stocks/buy", handler.POST_User_Stocks_Buy)
		user.POST("/stocks/sell", handler.POST_User_Stocks_Sell)
	}

	manage := server.Group("/manage")
	{
		//manage.GET("/confirms", handler.GetConfirms)
		manage.POST("/confirmation", handler.Confirm)

		//manage.GET("/accounts", handler.GetAccounts)
		//manage.POST("/freeze-account", handler.FreezeAccount)
		//manage.POST("/block-account", handler.BlockAccount)

		//manage.GET("/transactions", handler.GetTransactions)
		//manage.POST("/cancel-transaction", handler.CancelTransaction)

	}

	server.POST("/TEST", handler.DBTEST) // DONT TOUCH
}
