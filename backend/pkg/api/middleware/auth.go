package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/shion0625/FYP/backend/pkg/api/handler/response"
	"github.com/shion0625/FYP/backend/pkg/service/token"
)

const (
	authorizationHeaderKey string = "Authorization"
	authorizationType      string = "Bearer"
	minAuthFieldsLength    int    = 2
)

// Get User Auth middleware.
func (c *middleware) AuthenticateUser(next echo.HandlerFunc) echo.HandlerFunc {
	// return c.middlewareUsingCookie(token.User)
	return c.authorize(next, token.User)
}

// Get Admin Auth middleware.
func (c *middleware) AuthenticateAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	// return c.middlewareUsingCookie(token.Admin)
	return c.authorize(next, token.Admin)
}

// authorize request on request header using user type.
func (c *middleware) authorize(next echo.HandlerFunc, tokenUser token.UserType) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		authorizationValues := ctx.Request().Header.Get(authorizationHeaderKey)

		authFields := strings.Fields(authorizationValues)

		if len(authFields) < minAuthFieldsLength {
			err := errors.New("authorization token not provided properly with prefix of Bearer")

			response.ErrorResponse(ctx, http.StatusUnauthorized, "Failed to authorize request", err, nil)

			return nil
		}

		authType := authFields[0]
		accessToken := authFields[1]

		if !strings.EqualFold(authType, authorizationType) {
			err := errors.New("invalid authorization type")
			response.ErrorResponse(ctx, http.StatusUnauthorized, "Unauthorized user", err, nil)

			return nil
		}

		tokenVerifyReq := token.VerifyTokenRequest{
			TokenString: accessToken,
			UsedFor:     tokenUser,
		}

		verifyRes, err := c.tokenService.VerifyToken(tokenVerifyReq)
		if err != nil {
			response.ErrorResponse(ctx, http.StatusUnauthorized, "Unauthorized user", err, nil)

			return nil
		}

		ctx.Set("userId", verifyRes.UserID)

		return next(ctx)
	}
}
