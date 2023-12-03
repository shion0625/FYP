package middleware

import (
	"github.com/labstack/echo/v4"
)

func (m *middleware) AccessControlExposeHeaders(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("Access-Control-Expose-Headers", "access_token")

		return next(c)
	}
}
