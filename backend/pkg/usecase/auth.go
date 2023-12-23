package usecase

import (
	"fmt"
	"log"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/shion0625/FYP/backend/pkg/api/handler/request"
	"github.com/shion0625/FYP/backend/pkg/domain"
	repoInterfaces "github.com/shion0625/FYP/backend/pkg/repository/interfaces"
	"github.com/shion0625/FYP/backend/pkg/service/token"
	"github.com/shion0625/FYP/backend/pkg/usecase/interfaces"
	"github.com/shion0625/FYP/backend/pkg/utils"
)

const (
	countryCode       = "+91"
	otpExpireDuration = time.Minute * 2
)

const (
	AccessTokenDuration  = time.Minute * 25   // 25Minutes
	RefreshTokenDuration = time.Hour*24*7 + 1 //  7 days+ 1hour
)

type authUseCase struct {
	userRepo     repoInterfaces.UserRepository
	tokenService token.TokenService
	authRepo     repoInterfaces.AuthRepository
}

func NewAuthUseCase(
	userRepo repoInterfaces.UserRepository,
	tokenService token.TokenService,
	authRepo repoInterfaces.AuthRepository,
) interfaces.AuthUseCase {
	return &authUseCase{
		userRepo:     userRepo,
		tokenService: tokenService,
		authRepo:     authRepo,
	}
}

func (a *authUseCase) UserLogin(ctx echo.Context, loginInfo request.Login) (string, error) {
	var (
		user domain.User
		err  error
	)

	switch {
	case loginInfo.Email != "":
		user, err = a.userRepo.FindUserByEmail(ctx, loginInfo.Email)
	case loginInfo.UserName != "":
		user, err = a.userRepo.FindUserByUserName(ctx, loginInfo.UserName)
	case loginInfo.Phone != "":
		user, err = a.userRepo.FindUserByPhoneNumber(ctx, loginInfo.Phone)
	default:
		return "", ErrEmptyLoginCredentials
	}

	if err != nil {
		return "", fmt.Errorf("failed to find user from database: %w", err)
	}

	if user.ID == "" {
		return "", ErrUserNotExist
	}

	if user.BlockStatus {
		return "", ErrUserBlocked
	}

	if err := utils.ComparePasswordWithHashedPassword(loginInfo.Password, user.Password); err != nil {
		return "", ErrWrongPassword
	}

	return user.ID, nil
}

func (a *authUseCase) UserSignUp(ctx echo.Context, signUpDetails domain.User) (string, error) {
	existUser, err := a.userRepo.FindUserByUserNameEmailOrPhone(ctx, signUpDetails)
	if err != nil {
		return "", fmt.Errorf("failed to check user %w", err)
	}

	if existUser != (domain.User{}) {
		// 一致しているプロパティをエラー内容として返す
		errorMsg := "failed to check user details already exist:"
		if signUpDetails.UserName == existUser.UserName {
			errorMsg += "\rUserName already exists"
		}

		if signUpDetails.Email == existUser.Email {
			errorMsg += "\rEmail already exists"
		}

		if signUpDetails.Phone == existUser.Phone {
			errorMsg += "\rPhone already exists"
		}

		return "", fmt.Errorf("%s.", errorMsg)
	}

	// // if user credentials already exist and  verified then return it as errors
	// if existUser.ID != "" && existUser.Verified {
	// 	err = utils.CompareUserExistingDetails(existUser, signUpDetails)

	// 	return "", fmt.Errorf("failed to user is not otp verified: %w", err)
	// }

	userID := existUser.ID

	if userID == "" { // if user not exist then save user on database
		hashPass, err := utils.GenerateHashFromPassword(signUpDetails.Password)
		if err != nil {
			return "", fmt.Errorf("failed to hash the password: %w", err)
		}

		signUpDetails.Password = hashPass
		_, err = a.userRepo.SaveUser(ctx, signUpDetails)

		if err != nil {
			return "", fmt.Errorf("failed to save user details: %w", err)
		}
	}

	return "success", nil
}

func (a *authUseCase) GenerateAccessToken(ctx echo.Context, tokenParams interfaces.GenerateTokenParams) (string, error) {
	tokenReq := token.GenerateTokenRequest{
		UserID:   tokenParams.UserID,
		UsedFor:  tokenParams.UserType,
		ExpireAt: time.Now().Add(AccessTokenDuration),
	}

	tokenRes, err := a.tokenService.GenerateToken(tokenReq)
	if err != nil {
		return "", fmt.Errorf("failed to generate access token: %w", err)
	}

	return tokenRes.TokenString, nil
}

func (a *authUseCase) GenerateRefreshToken(ctx echo.Context, tokenParams interfaces.GenerateTokenParams) (string, error) {
	expireAt := time.Now().Add(RefreshTokenDuration)
	tokenReq := token.GenerateTokenRequest{
		UserID:   tokenParams.UserID,
		UsedFor:  tokenParams.UserType,
		ExpireAt: expireAt,
	}

	tokenRes, err := a.tokenService.GenerateToken(tokenReq)
	if err != nil {
		return "", fmt.Errorf("failed to generate refresh token: %w", err)
	}

	err = a.authRepo.SaveRefreshSession(ctx, domain.RefreshSession{
		UserID:       tokenParams.UserID,
		TokenID:      tokenRes.TokenID,
		RefreshToken: tokenRes.TokenString,
		ExpireAt:     expireAt,
	})

	if err != nil {
		return "", fmt.Errorf("failed to save refresh session: %w", err)
	}

	log.Printf("successfully refresh token created and refresh session stored in database")

	return tokenRes.TokenString, nil
}

func (c *authUseCase) VerifyAndGetRefreshTokenSession(ctx echo.Context, refreshToken string, usedFor token.UserType) (domain.RefreshSession, error) {
	verifyReq := token.VerifyTokenRequest{
		TokenString: refreshToken,
		UsedFor:     usedFor,
	}

	verifyRes, err := c.tokenService.VerifyToken(verifyReq)
	if err != nil {
		return domain.RefreshSession{}, fmt.Errorf("failed to save refresh session: %w", err)
	}

	refreshSession, err := c.authRepo.FindRefreshSessionByTokenID(ctx, verifyRes.TokenID)
	if err != nil {
		return refreshSession, fmt.Errorf("failed to find refresh session by token ID: %w", err)
	}

	if refreshSession.TokenID == "" {
		return refreshSession, ErrRefreshSessionNotExist
	}

	if time.Since(refreshSession.ExpireAt) > 0 {
		return domain.RefreshSession{}, ErrRefreshSessionExpired
	}

	if refreshSession.IsBlocked {
		return domain.RefreshSession{}, ErrRefreshSessionBlocked
	}

	return refreshSession, nil
}
