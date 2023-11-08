package interfaces

import (
	"github.com/labstack/echo/v4"
	"github.com/shion0625/FYP/backend/pkg/api/handler/request"
	"github.com/shion0625/FYP/backend/pkg/domain"
)

type AuthUseCase interface {
	UserSignUp(ctx echo.Context, signUpDetails domain.User) (otpID string, err error)
	// SingUpOtpVerify(ctx echo.Context, otpVerifyDetails request.OTPVerify) (userID uint, err error)
	// GoogleLogin(ctx echo.Context, user domain.User) (userID uint, err error)
	UserLogin(ctx echo.Context, loginInfo request.Login) (userID uint, err error)
	// UserLoginOtpSend(ctx echo.Context, loginInfo request.OTPLogin) (otpID string, err error)
	// LoginOtpVerify(ctx echo.Context, otpVerifyDetails request.OTPVerify) (userID uint, err error)

	// // admin
	// AdminLogin(ctx echo.Context, loginInfo request.Login) (adminID uint, err error)
	// // token
	// GenerateAccessToken(ctx echo.Context, tokenParams GenerateTokenParams) (tokenString string, err error)
	// GenerateRefreshToken(ctx echo.Context, tokenParams GenerateTokenParams) (tokenString string, err error)
	// VerifyAndGetRefreshTokenSession(ctx echo.Context, refreshToken string, usedFor token.UserType) (domain.RefreshSession, error)
}

// type GenerateTokenParams struct {
// 	UserID   uint
// 	UserType token.UserType
// }
