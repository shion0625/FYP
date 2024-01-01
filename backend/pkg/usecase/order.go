package usecase

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/shion0625/FYP/backend/pkg/api/handler/request"
	"github.com/shion0625/FYP/backend/pkg/api/handler/response"
	"github.com/shion0625/FYP/backend/pkg/config"
	repoInterfaces "github.com/shion0625/FYP/backend/pkg/repository/interfaces"
	"github.com/shion0625/FYP/backend/pkg/usecase/interfaces"
	"github.com/shion0625/FYP/backend/pkg/utils"
)

type orderUseCase struct {
	orderRepo     repoInterfaces.OrderRepository
	creditCardKey string
}

func NewOrderUseCase(
	cfg *config.Config,
	orderRepo repoInterfaces.OrderRepository,
) interfaces.OrderUseCase {
	return &orderUseCase{
		orderRepo:     orderRepo,
		creditCardKey: cfg.CreditCardKey,
	}
}

func (o *orderUseCase) PayOrder(ctx echo.Context, userID string, payOrder request.PayOrder) error {
	if err := o.orderRepo.Transactions(ctx, o.updateStockAndPayOrder(ctx, userID, payOrder)); err != nil {
		return fmt.Errorf("order payment failed: %w", err)
	}

	return nil
}

func (o *orderUseCase) updateStockAndPayOrder(ctx echo.Context, userID string, payOrder request.PayOrder) func(repo repoInterfaces.OrderRepository) error {
	return func(repo repoInterfaces.OrderRepository) error {
		for _, itemInfo := range payOrder.ProductItemInfo {
			newStock, err := repo.UpdateProductItemStock(ctx, itemInfo.ProductItemID, itemInfo.Count)
			if err != nil {
				return fmt.Errorf("stock update for productItem failed, new stock: %d: %w", newStock, err)
			}
		}

		if err := repo.PayOrder(ctx, payOrder.PaymentMethodID); err != nil {
			return fmt.Errorf("%w", err)
		}

		if err := repo.SaveOrder(ctx, userID, payOrder); err != nil {
			return fmt.Errorf("%w", err)
		}

		return nil
	}
}

func (o *orderUseCase) GetAllShopOrders(ctx echo.Context, userID string, pagination request.Pagination) (orderHistory []response.Order, err error) {
	orderHistory, err = o.orderRepo.GetShopOrders(ctx, userID, pagination)

	for i, order := range orderHistory {
		creditNumberDecrypted := utils.Decrypt(order.PaymentMethod.Number, userID+o.creditCardKey)
		if len(creditNumberDecrypted) >= MinCreditNumberLength {
			orderHistory[i].PaymentMethod.Number = creditNumberDecrypted[len(creditNumberDecrypted)-4:]
		}
	}

	if err != nil {
		return nil, fmt.Errorf("address retrieval failed: %w", err)
	}

	return orderHistory, nil
}
