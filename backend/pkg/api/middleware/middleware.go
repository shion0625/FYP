package middleware

import (
	"github.com/labstack/echo/v4"
)

type Middleware interface {
	// AuthenticateUser() echo.HandlerFunc
	// AuthenticateAdmin() echo.HandlerFunc
	// TrimSpaces() echo.HandlerFunc
	Context(echo.HandlerFunc) echo.HandlerFunc
}

type middleware struct{}

func NewMiddleware() Middleware {
	return &middleware{}
}
