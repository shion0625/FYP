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
		response.ErrorResponse(ctx, http.StatusBadRequest, "Request data is invalid", err, nil)

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

		response.ErrorResponse(ctx, statusCode, "Login failed", err, nil)

		return nil
	}

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
		response.ErrorResponse(ctx, http.StatusBadRequest, "Request data is invalid", err, nil)

		return nil
	}

	var user domain.User
	if err := copier.Copy(&user, body); err != nil {
		response.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to copy user details", err, nil)

		return nil
	}

	_, err := c.authUseCase.UserSignUp(ctx, user)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if errors.Is(err, usecase.ErrUserAlreadyExit) {
			statusCode = http.StatusConflict
		}

		response.ErrorResponse(ctx, statusCode, "Signup failed", err, nil)

		return nil
	}

	response.SuccessResponse(ctx, http.StatusCreated,
		"Account created successfully", nil)

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
	ctx.Response().Header().Set("refresh_token", refreshToken)
	tokenRes := response.TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		UserID:       userID,
	}
	response.SuccessResponse(ctx, http.StatusOK, "Login successful", tokenRes)
}

func (a *AuthHandler) renewAccessToken(tokenUser token.UserType) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var body request.RefreshToken

		if err := ctx.Bind(&body); err != nil {
			response.ErrorResponse(ctx, http.StatusBadRequest, "Failed to bind JSON", err, body)

			return nil
		}

		if err := ctx.Validate(body); err != nil {
			response.ErrorResponse(ctx, http.StatusBadRequest, "Request data is invalid", err, nil)

			return nil
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

			response.ErrorResponse(ctx, statusCode, "Failed to verify refresh token", err, nil)

			return nil
		}

		accessTokenParams := usecaseInterface.GenerateTokenParams{
			UserID:   refreshSession.UserID,
			UserType: tokenUser,
		}

		accessToken, err := a.authUseCase.GenerateAccessToken(ctx, accessTokenParams)
		if err != nil {
			response.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to generate access token", err, nil)

			return nil
		}

		authorizationValue := authorizationType + " " + accessToken
		ctx.Response().Header().Set(authorizationHeaderKey, authorizationValue)
		ctx.Response().Header().Set("access_token", accessToken)

		response.SuccessResponse(ctx, http.StatusOK, "Access token generated successfully using refresh token", accessToken)

		return nil
	}
}
