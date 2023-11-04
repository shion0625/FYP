package interfaces

import "github.com/labstack/echo/v4"

type AuthHandler interface {
	//userSide
	UserLogin(ctx echo.Context) echo.HandlerFunc
	// UserSignUp(ctx *echo.Context)
	// UserSignUpVerify(ctx *echo.Context)

	// UserGoogleAuthInitialize(ctx *echo.Context)
	// UserGoogleAuthLoginPage(ctx *echo.Context)
	// UserGoogleAuthCallBack(ctx *echo.Context)

	// UserLoginOtpVerify(ctx *echo.Context)
	// UserLoginOtpSend(ctx *echo.Context)

	// UserRenewAccessToken() echo.HandlerFunc

	// //admin side
	// AdminLogin(ctx *echo.Context)
	// AdminRenewAccessToken() echo.HandlerFunc
}
