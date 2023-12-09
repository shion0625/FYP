package interfaces

import (
	"github.com/labstack/echo/v4"
	"github.com/shion0625/FYP/backend/pkg/api/handler/request"
	"github.com/shion0625/FYP/backend/pkg/api/handler/response"
)

type OrderUseCase interface {
	GetAllShopOrders(ctx echo.Context, userID string, pagination request.Pagination) ([]response.Order, error)
	PayOrder(ctx echo.Context, userID string, product request.PayOrder) error
}
