package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/shion0625/FYP/backend/pkg/api/handler/interfaces"
	"github.com/shion0625/FYP/backend/pkg/api/handler/request"
	"github.com/shion0625/FYP/backend/pkg/api/handler/response"
	usecaseInterface "github.com/shion0625/FYP/backend/pkg/usecase/interfaces"
	"github.com/shion0625/FYP/backend/pkg/utils"
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
	userID, err := utils.GetUserIdFromContext(ctx)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusInternalServerError, "Unable to retrieve user id", err, nil)

		return nil
	}

	var body request.PayOrder
	if err := ctx.Bind(&body); err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, BindJsonFailMessage, err, nil)

		return nil
	}

	if err := ctx.Validate(body); err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, "Request data is invalid", err, nil)

		return nil
	}

	if err := o.orderUseCase.PayOrder(ctx, userID, body); err != nil {
		response.ErrorResponse(ctx, http.StatusInternalServerError, "Payment process failed", err, nil)

		return nil
	}

	response.SuccessResponse(ctx, http.StatusOK, "Order purchased successfully", nil)

	return nil
}

func (o *OrderHandler) GetOrderHistory(ctx echo.Context) error {
	pagination := request.GetPagination(ctx)

	userID, err := utils.GetUserIdFromContext(ctx)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusInternalServerError, "Unable to retrieve user id", err, nil)

		return nil
	}

	orderHistories, err := o.orderUseCase.GetAllShopOrders(ctx, userID, pagination)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to fetch order history", err, nil)

		return nil
	}

	if len(orderHistories) == 0 {
		response.SuccessResponse(ctx, http.StatusOK, "No orders found", nil)

		return nil
	}

	response.SuccessResponse(ctx, http.StatusOK, "Order history fetched successfully", orderHistories)

	return nil
}
