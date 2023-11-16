package router

import (
	"github.com/labstack/echo/v4"
	handlerInterfaces "github.com/shion0625/FYP/backend/pkg/api/handler/interfaces"
	"github.com/shion0625/FYP/backend/pkg/api/middleware"
)

func UserRoutes(api *echo.Group, middleware middleware.Middleware, authHandler handlerInterfaces.AuthHandler,
	userHandler handlerInterfaces.UserHandler) {
	auth := api.Group("/auth")
	{
		signup := auth.Group("/sign-up")
		{
			signup.POST("/", authHandler.UserSignUp)
			// signup.POST("/verify", authHandler.UserSignUpVerify)
		}

		login := auth.Group("/login")
		{
			login.POST("/", authHandler.UserLogin)
			// login.POST("/otp/send", authHandler.UserLoginOtpSend)
			// login.POST("/otp/verify", authHandler.UserLoginOtpVerify)
		}
	}

	api.Use(middleware.AuthenticateUser())
	{
		// profile
		account := api.Group("/account")
		{
			account.GET("/", userHandler.GetProfile)
			account.PUT("/", userHandler.UpdateProfile)

			account.GET("/address", userHandler.GetAllAddresses) // to show all address and // show countries
			account.POST("/address", userHandler.SaveAddress)    // to add a new address
			account.PUT("/address", userHandler.UpdateAddress)   // to edit address
			// account.DELETE("/address", userHandler.DeleteAddress)
		}
	}
}
