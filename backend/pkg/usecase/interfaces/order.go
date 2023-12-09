package interfaces

import (
	"github.com/labstack/echo/v4"
	"github.com/shion0625/FYP/backend/pkg/api/handler/request"
	"github.com/shion0625/FYP/backend/pkg/api/handler/response"
)

type OrderUseCase interface {
	PayOrder(ctx echo.Context, product request.PayOrder) error
	GetAllShopOrders(ctx echo.Context, userId string, pagination request.Pagination) ([]response.Order, error)
}
