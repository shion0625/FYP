package handler

import (
	"errors"
	"net/http"

	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
	"github.com/shion0625/FYP/backend/pkg/api/handler/interfaces"
	"github.com/shion0625/FYP/backend/pkg/api/handler/request"
	"github.com/shion0625/FYP/backend/pkg/api/handler/response"
	"github.com/shion0625/FYP/backend/pkg/config"
	"github.com/shion0625/FYP/backend/pkg/domain"
	"github.com/shion0625/FYP/backend/pkg/service/token"
	"github.com/shion0625/FYP/backend/pkg/usecase"
	usecaseInterface "github.com/shion0625/FYP/backend/pkg/usecase/interfaces"
)

const (
	authorizationHeaderKey = "Authorization"
	authorizationType      = "Bearer"
)

type AuthHandler struct {
	authUseCase usecaseInterface.AuthUseCase
	config      *config.Config
}

func NewAuthHandler(authUsecase usecaseInterface.AuthUseCase, config *config.Config) interfaces.AuthHandler {
	return &AuthHandler{
		authUseCase: authUsecase,
		config:      config,
	}
}

func (a *AuthHandler) UserLogin(ctx echo.Context) error {
	var body request.Login

	if err := ctx.Bind(&body); err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, BindJsonFailMessage, err, body)

		return nil
	}

	if err := ctx.Validate(body); err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request data", err, nil)

		return nil
	}

	userID, err := a.authUseCase.UserLogin(ctx, body)
	if err != nil {
		var statusCode int

		switch {
		case errors.Is(err, usecase.ErrEmptyLoginCredentials):
			statusCode = http.StatusBadRequest
		case errors.Is(err, usecase.ErrUserNotExist):
			statusCode = http.StatusNotFound
		case errors.Is(err, usecase.ErrUserBlocked):
			statusCode = http.StatusForbidden
		// case errors.Is(err, usecase.ErrUserNotVerified):
		// 	statusCode = http.StatusUnauthorized
		case errors.Is(err, usecase.ErrWrongPassword):
			statusCode = http.StatusUnauthorized
		default:
			statusCode = http.StatusInternalServerError
		}

		response.ErrorResponse(ctx, statusCode, "Failed to login", err, nil)

		return nil
	}

	print("login")

	// common functionality for admin and user
	a.setupTokenAndResponse(ctx, token.User, userID)

	return nil
}

func (c *AuthHandler) UserSignUp(ctx echo.Context) error {
	var body request.SignUp

	if err := ctx.Bind(&body); err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, BindJsonFailMessage, err, body)

		return nil
	}

	if err := ctx.Validate(body); err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request data", err, nil)

		return nil
	}

	var user domain.User
	if err := copier.Copy(&user, body); err != nil {
		response.ErrorResponse(ctx, http.StatusInternalServerError, "failed to copy details", err, nil)

		return nil
	}

	_, err := c.authUseCase.UserSignUp(ctx, user)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if errors.Is(err, usecase.ErrUserAlreadyExit) {
			statusCode = http.StatusConflict
		}

		response.ErrorResponse(ctx, statusCode, "Failed to signup", err, nil)

		return nil
	}

	response.SuccessResponse(ctx, http.StatusCreated,
		"Successfully account created", nil)

	return nil
}

func (a *AuthHandler) UserRenewAccessToken() func(ctx echo.Context) error {
	return a.renewAccessToken(token.User)
}

func (a *AuthHandler) setupTokenAndResponse(ctx echo.Context, tokenUser token.UserType, userID string) {
	tokenParams := usecaseInterface.GenerateTokenParams{
		UserID:   userID,
		UserType: tokenUser,
	}

	accessToken, err := a.authUseCase.GenerateAccessToken(ctx, tokenParams)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to generate access token", err, nil)

		return
	}

	refreshToken, err := a.authUseCase.GenerateRefreshToken(ctx, usecaseInterface.GenerateTokenParams{
		UserID:   userID,
		UserType: tokenUser,
	})
	if err != nil {
		response.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to generate refresh token", err, nil)

		return
	}

	authorizationValue := authorizationType + " " + accessToken
	ctx.Response().Header().Set(authorizationHeaderKey, authorizationValue)
	ctx.Response().Header().Set("access_token", accessToken)
	ctx.Response().Header().Set("Access-Control-Expose-Headers", "access_token")

	// リフレッシュトークンをHTTP Only Cookieに設定
	cookie := new(http.Cookie)
	cookie.Name = "refresh_token"
	cookie.Value = refreshToken
	cookie.HttpOnly = true
	ctx.SetCookie(cookie)

	tokenRes := response.TokenResponse{
		AccessToken: accessToken,
	}

	response.SuccessResponse(ctx, http.StatusOK, "Successfully logged in", tokenRes)
}

func (a *AuthHandler) renewAccessToken(tokenUser token.UserType) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		cookie, err := ctx.Cookie("refresh_token")
		if err != nil {
			response.ErrorResponse(ctx, http.StatusBadRequest, BindJsonFailMessage, err, nil)

			return nil
		}

		body := request.RefreshToken{
			RefreshToken: cookie.Value,
		}

		refreshSession, err := a.authUseCase.VerifyAndGetRefreshTokenSession(ctx, body.RefreshToken, tokenUser)
		if err != nil {
			var statusCode int

			switch {
			case errors.Is(err, usecase.ErrInvalidRefreshToken):
				statusCode = http.StatusUnauthorized
			case errors.Is(err, usecase.ErrRefreshSessionNotExist):
				statusCode = http.StatusNotFound
			case errors.Is(err, usecase.ErrRefreshSessionExpired):
				statusCode = http.StatusGone
			case errors.Is(err, usecase.ErrRefreshSessionBlocked):
				statusCode = http.StatusForbidden
			default:
				statusCode = http.StatusInternalServerError
			}

			response.ErrorResponse(ctx, statusCode, "Failed verify refresh token", err, nil)

			return nil
		}

		accessTokenParams := usecaseInterface.GenerateTokenParams{
			UserID:   refreshSession.UserID,
			UserType: tokenUser,
		}

		accessToken, err := a.authUseCase.GenerateAccessToken(ctx, accessTokenParams)
		if err != nil {
			response.ErrorResponse(ctx, http.StatusInternalServerError, "Failed generate access token", err, nil)

			return nil
		}

		authorizationValue := authorizationType + " " + accessToken
		ctx.Response().Header().Set(authorizationHeaderKey, authorizationValue)
		ctx.Response().Header().Set("access_token", accessToken)

		accessTokenRes := response.TokenResponse{
			AccessToken: accessToken,
		}
		response.SuccessResponse(ctx, http.StatusOK, "Successfully generated access token using refresh token", accessTokenRes)

		return nil
	}
}
