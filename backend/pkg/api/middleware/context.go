package middleware

import (
	"github.com/labstack/echo/v4"
)

type CustomContext struct {
	echo.Context
}

func (c *middleware) Context(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
        cc := &CustomContext{c}
        return next(cc)
    }
}
