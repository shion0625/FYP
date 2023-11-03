package handler

import (
	"github.com/labstack/echo/v4"
)

const (
	authorizationHeaderKey = "Authorization"
	authorizationType      = "Bearer"
)

type AuthHandler struct {
	authUseCase usecaseInterface.AuthUseCase
	config      config.Config
}

func NewAuthHandler(authUsecase usecaseInterface.AuthUseCase, config config.Config) interfaces.AuthHandler {
	return &AuthHandler{
		authUseCase: authUsecase,
		config:      config,
	}
}

func (c *AuthHandler) UserLogin(ctx *echo.Context) {

	var body request.Login

	if err := ctx.ShouldBindJSON(&body); err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, BindJsonFailMessage, err, body)
		return
	}

	userID, err := c.authUseCase.UserLogin(ctx, body)

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
		return
	}

	print("login")
	// common functionality for admin and user
	// c.setupTokenAndResponse(ctx, token.User, userID)
}


// func (c *AuthHandler) setupTokenAndResponse(ctx *gin.Context, tokenUser token.UserType, userID uint) {

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
