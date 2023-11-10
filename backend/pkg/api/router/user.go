package router

import (
	"github.com/labstack/echo/v4"
	handlerInterfaces "github.com/shion0625/FYP/backend/pkg/api/handler/interfaces"
)

func UserRoutes(api *echo.Group, authHandler handlerInterfaces.AuthHandler) {
	auth := api.Group("/auth")

	signup := auth.Group("/sign-up")
	{
		signup.POST("/", func(c echo.Context) error {
			authHandler.UserSignUp(c)

			return nil
		})
		// signup.POST("/verify", authHandler.UserSignUpVerify)
	}

	login := auth.Group("/login")
	{
		login.POST("/", func(c echo.Context) error {
			authHandler.UserLogin(c)

			return nil
		})
		// login.POST("/otp/send", authHandler.UserLoginOtpSend)
		// login.POST("/otp/verify", authHandler.UserLoginOtpVerify)
	}
}
