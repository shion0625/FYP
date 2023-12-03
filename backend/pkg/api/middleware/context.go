package middleware

import (
	"github.com/labstack/echo/v4"
)

type CustomContext struct {
	echo.Context
}

func (m *middleware) Context(next echo.HandlerFunc) echo.HandlerFunc {
	return func(m echo.Context) error {
		cc := &CustomContext{m}

		return next(cc)
	}
}
