package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/shion0625/FYP/backend/pkg/api/handler/interfaces"
	"github.com/shion0625/FYP/backend/pkg/api/handler/request"
	"github.com/shion0625/FYP/backend/pkg/api/handler/response"
	usecaseInterface "github.com/shion0625/FYP/backend/pkg/usecase/interfaces"
)

type OrderHandler struct {
	orderUseCase usecaseInterface.OrderUseCase
}

func NewOrderHandler(orderUsecase usecaseInterface.OrderUseCase) interfaces.OrderHandler {
	return &OrderHandler{
		orderUseCase: orderUsecase,
	}
}

func (o *OrderHandler) PayOrder(ctx echo.Context) error {
	var payOrder request.PayOrder
	if err := ctx.Bind(&payOrder); err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, BindJsonFailMessage, err, nil)

		return fmt.Errorf("Bind error: %w", err)
	}

	if err := o.orderUseCase.PayOrder(ctx, payOrder); err != nil {
		return fmt.Errorf("failed to PayOrder: %w", err)
	}

	return nil
}