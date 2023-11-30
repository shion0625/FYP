package interfaces

import (
	"github.com/labstack/echo/v4"
	"github.com/shion0625/FYP/backend/pkg/api/handler/request"
)

type OrderRepository interface {
	Transactions(ctx echo.Context, trxFn func(repo OrderRepository) error) error
	UpdateProductItemStock(ctx echo.Context, productItemID uint, purchaseQuantity uint) (uint, error)
	SaveOrder(ctx echo.Context, payOrder request.PayOrder) error
	PayOrder(ctx echo.Context, payOrder request.PayOrder) error
}
