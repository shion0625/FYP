package router

import (
	"github.com/labstack/echo/v4"
	handlerInterfaces "github.com/shion0625/FYP/backend/pkg/api/handler/interfaces"
)

func UserRoutes(api *echo.Group, authHandler handlerInterfaces.AuthHandler) {
	auth := api.Group("/auth")

	// signup := auth.Group("/sign-up")
	// {
	// 	signup.GET("/", func(c echo.Context) error {
	// 		print("こんにちは")
	// 		return nil
	// 	})
	// 	signup.POST("/verify", authHandler.UserSignUpVerify)
	// }

	login := auth.Group("/sign-in")
	{
		login.POST("/", func(c echo.Context) error { authHandler.UserLogin(c); return nil })
		// login.POST("/otp/send", authHandler.UserLoginOtpSend)
		// login.POST("/otp/verify", authHandler.UserLoginOtpVerify)
	}

	// goath := auth.Group("/google-auth")
	// {
	// 	goath.GET("/", authHandler.UserGoogleAuthLoginPage)
	// 	goath.GET("/initialize", authHandler.UserGoogleAuthInitialize)
	// 	goath.GET("/callback", authHandler.UserGoogleAuthCallBack)
	// }

	// auth.POST("/renew-access-token", authHandler.UserRenewAccessToken())

	// api.POST("/logout")
}
