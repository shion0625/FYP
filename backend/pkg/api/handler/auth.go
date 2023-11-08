package handler

import (
	"errors"
	"net/http"

	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
	handlerInterface "github.com/shion0625/FYP/backend/pkg/api/handler/interfaces"
	"github.com/shion0625/FYP/backend/pkg/api/handler/request"
	"github.com/shion0625/FYP/backend/pkg/api/handler/response"
	"github.com/shion0625/FYP/backend/pkg/config"
	"github.com/shion0625/FYP/backend/pkg/domain"
	"github.com/shion0625/FYP/backend/pkg/usecase"
	usecaseInterface "github.com/shion0625/FYP/backend/pkg/usecase/interfaces"
	"net/http"
)

const (
	authorizationHeaderKey = "Authorization"
	authorizationType      = "Bearer"
)

type AuthHandler struct {
	authUseCase usecaseInterface.AuthUseCase
	config      *config.Config
}

func NewAuthHandler(authUsecase usecaseInterface.AuthUseCase, config *config.Config) handlerInterface.AuthHandler {
	return &AuthHandler{
		authUseCase: authUsecase,
		config:      config,
	}
}

func (a *AuthHandler) UserLogin(ctx echo.Context) echo.HandlerFunc {
	var body request.Login

	if err := ctx.Bind(&body); err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, BindJsonFailMessage, err, body)
		return nil
	}

	_, err := a.authUseCase.UserLogin(ctx, body)
	if err != nil {
		var statusCode int

		switch {
		case errors.Is(err, usecase.ErrEmptyLoginCredentials):
			statusCode = http.StatusBadRequest
		case errors.Is(err, usecase.ErrUserNotExist):
			statusCode = http.StatusNotFound
		case errors.Is(err, usecase.ErrUserBlocked):
			statusCode = http.StatusForbidden
		case errors.Is(err, usecase.ErrUserNotVerified):
			statusCode = http.StatusUnauthorized
		case errors.Is(err, usecase.ErrWrongPassword):
			statusCode = http.StatusUnauthorized
		default:
			statusCode = http.StatusInternalServerError
		}

		response.ErrorResponse(ctx, statusCode, "Failed to login", err, nil)
		return nil
	}
	print("login")

	return nil
	// common functionality for admin and user
	// c.setupTokenAndResponse(ctx, token.User, userID)
}

func (c *AuthHandler) UserSignUp(ctx echo.Context) echo.HandlerFunc {
	var body request.SignUp

	if err := ctx.Bind(&body); err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, BindJsonFailMessage, err, body)
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
		"Successfully account created")
	return nil
}

// func (c *AuthHandler) setupTokenAndResponse(ctx *echo.Context, tokenUser token.UserType, userID uint) {

// 	tokenParams := usecaseInterface.GenerateTokenParams{
// 		UserID:   userID,
// 		UserType: tokenUser,
// 	}

// 	accessToken, err := c.authUseCase.GenerateAccessToken(ctx, tokenParams)

// 	if err != nil {
// 		response.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to generate access token", err, nil)
// 		return
// 	}

// 	refreshToken, err := c.authUseCase.GenerateRefreshToken(ctx, usecaseInterface.GenerateTokenParams{
// 		UserID:   userID,
// 		UserType: tokenUser,
// 	})
// 	if err != nil {
// 		response.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to generate refresh token", err, nil)
// 		return
// 	}

// 	authorizationValue := authorizationType + " " + accessToken
// 	ctx.Header(authorizationHeaderKey, authorizationValue)

// 	ctx.Header("access_token", accessToken)
// 	ctx.Header("refresh_token", refreshToken)

// 	tokenRes := response.TokenResponse{
// 		AccessToken:  accessToken,
// 		RefreshToken: refreshToken,
// 	}

// 	response.SuccessResponse(ctx, http.StatusOK, "Successfully logged in", tokenRes)
// }
