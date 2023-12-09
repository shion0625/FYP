package interfaces

import (
	"github.com/labstack/echo/v4"
	"github.com/shion0625/FYP/backend/pkg/api/handler/request"
	"github.com/shion0625/FYP/backend/pkg/api/handler/response"
)

type OrderRepository interface {
	Transactions(ctx echo.Context, trxFn func(repo OrderRepository) error) error
	UpdateProductItemStock(ctx echo.Context, productItemID uint, purchaseQuantity uint) (uint, error)
	SaveOrder(ctx echo.Context, userID string, payOrder request.PayOrder) error
	PayOrder(ctx echo.Context, paymentMethodID uint) error
	GetShopOrders(ctx echo.Context, userID string, pagination request.Pagination) ([]response.Order, error)
}
