package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/shion0625/FYP/backend/pkg/service/token"
)

type Middleware interface {
	AuthenticateUser(next echo.HandlerFunc) echo.HandlerFunc
	// AuthenticateAdmin() echo.HandlerFunc
	// TrimSpaces() echo.HandlerFunc
	Context(next echo.HandlerFunc) echo.HandlerFunc
	AccessControlExposeHeaders(next echo.HandlerFunc) echo.HandlerFunc
}

type middleware struct {
	tokenService token.TokenService
}

func NewMiddleware(tokenService token.TokenService) Middleware {
	return &middleware{
		tokenService: tokenService,
	}
}
