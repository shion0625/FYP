package interfaces

import (
	"github.com/labstack/echo/v4"
	"github.com/shion0625/FYP/backend/pkg/api/handler/request"
)

type OrderUseCase interface {
	PayOrder(ctx echo.Context, product request.PayOrder) error
}
