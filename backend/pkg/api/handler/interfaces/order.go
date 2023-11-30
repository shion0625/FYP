package interfaces

import "github.com/labstack/echo/v4"

type OrderHandler interface {
	PayOrder(ctx echo.Context) error
}
