package interfaces

import (
	"context"

	"github.com/shion0625/backend/pkg/api/handler/request"
	"github.com/shion0625/backend/pkg/api/handler/response"
	"github.com/shion0625/backend/pkg/domain"
)

type OrderUseCase interface {

	//
	SaveOrder(ctx context.Context, userID, addressID uint) (shopOrderID uint, err error)

	// Find order and order items
	FindAllShopOrders(ctx context.Context, pagination request.Pagination) (shopOrders []response.ShopOrder, err error)
	FindUserShopOrder(ctx context.Context, userID uint, pagination request.Pagination) ([]response.ShopOrder, error)
	FindOrderItems(ctx context.Context, shopOrderID uint, pagination request.Pagination) ([]response.OrderItem, error)

	// cancel order and change order status
	FindAllOrderStatuses(ctx context.Context) (orderStatuses []domain.OrderStatus, err error)
	UpdateOrderStatus(ctx context.Context, shopOrderID, changeStatusID uint) error
	CancelOrder(ctx context.Context, shopOrderID uint) error

	// return and update
	SubmitReturnRequest(ctx context.Context, returnDetails request.Return) error
	FindAllPendingOrderReturns(ctx context.Context, pagination request.Pagination) ([]response.OrderReturn, error)
	FindAllOrderReturns(ctx context.Context, pagination request.Pagination) ([]response.OrderReturn, error)
	UpdateReturnDetails(ctx context.Context, updateDetails request.UpdateOrderReturn) error

	// wallet
	FindUserWallet(ctx context.Context, userID uint) (wallet domain.Wallet, err error)
	FindUserWalletTransactions(ctx context.Context, userID uint, pagination request.Pagination) (transactions []domain.Transaction, err error)
}
