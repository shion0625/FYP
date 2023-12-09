package usecase

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/shion0625/FYP/backend/pkg/api/handler/request"
	"github.com/shion0625/FYP/backend/pkg/api/handler/response"
	repoInterfaces "github.com/shion0625/FYP/backend/pkg/repository/interfaces"
	"github.com/shion0625/FYP/backend/pkg/usecase/interfaces"
)

type orderUseCase struct {
	orderRepo repoInterfaces.OrderRepository
}

func NewOrderUseCase(
	orderRepo repoInterfaces.OrderRepository,
) interfaces.OrderUseCase {
	return &orderUseCase{
		orderRepo: orderRepo,
	}
}

func (o *orderUseCase) PayOrder(ctx echo.Context, userID string, payOrder request.PayOrder) error {
	if err := o.orderRepo.Transactions(ctx,
		func(repo repoInterfaces.OrderRepository) error {
			for _, itemInfo := range payOrder.ProductItemInfo {
				newStock, err := repo.UpdateProductItemStock(ctx, itemInfo.ProductItemID, itemInfo.Count)
				if err != nil {
					return fmt.Errorf("failed to update productItem stock to %d: %w", newStock, err)
				}
			}

			if err := repo.PayOrder(ctx, payOrder.PaymentMethodID); err != nil {
				return fmt.Errorf(": %w", err)
			}

			if err := repo.SaveOrder(ctx, userID, payOrder); err != nil {
				return fmt.Errorf(": %w", err)
			}

			return nil
		}); err != nil {
		return fmt.Errorf("failed to pay order: %w", err)
	}

	return nil
}

func (o *orderUseCase) GetAllShopOrders(ctx echo.Context, userID string, pagination request.Pagination) (orderHistory []response.Order, err error) {
	orderHistory, err = o.orderRepo.GetShopOrders(ctx, userID, pagination)
	if err != nil {
		return nil, fmt.Errorf("failed to find addresses: %w", err)
	}

	return orderHistory, nil
}
