package interfaces

import "github.com/labstack/echo/v4"

type AuthHandler interface {
	UserLogin(ctx echo.Context) error
	UserSignUp(ctx echo.Context) error
	// UserSignUpVerify(ctx *echo.Context)

	// UserGoogleAuthInitialize(ctx *echo.Context)
	// UserGoogleAuthLoginPage(ctx *echo.Context)
	// UserGoogleAuthCallBack(ctx *echo.Context)

	// UserLoginOtpVerify(ctx *echo.Context)
	// UserLoginOtpSend(ctx *echo.Context)

	UserRenewAccessToken() func(ctx echo.Context) error

	// //admin side
	// AdminLogin(ctx *echo.Context)
	// AdminRenewAccessToken() echo.HandlerFunc
}
