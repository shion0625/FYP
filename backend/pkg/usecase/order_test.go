package usecase_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/shion0625/FYP/backend/pkg/api/handler/request"
	"github.com/shion0625/FYP/backend/pkg/api/handler/response"
	repoInterfaces "github.com/shion0625/FYP/backend/pkg/repository/interfaces"
	orderMock "github.com/shion0625/FYP/backend/pkg/repository/mock"
	"github.com/shion0625/FYP/backend/pkg/usecase"
	"github.com/stretchr/testify/assert"
	gomock "go.uber.org/mock/gomock"
)

func TestOrderUseCase_PayOrder(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockOrderRepo := orderMock.NewMockOrderRepository(ctrl)

	order := usecase.NewOrderUseCase(mockOrderRepo)

	ctx := echo.New().NewContext(nil, nil)
	userID := "1"

	type input struct {
		userID   string
		payOrder request.PayOrder
	}

	type output struct {
		wantErr error
	}

	tests := map[string]struct {
		input         input
		prepareMockFn func(mr *orderMock.MockOrderRepository)
		want          output
	}{
		"Normal Case: PayOrder": {
			input{userID, request.PayOrder{
				PaymentMethodID: 1,
				ProductItemInfo: []request.ProductItemInfo{
					{
						ProductItemID: 1,
						Count:         1,
					},
				},
			}},
			func(mr *orderMock.MockOrderRepository) {
				mr.EXPECT().Transactions(gomock.Any(), gomock.Any()).Return(nil)
			},
			output{nil},
		},
		"Abnormal case: PayOrder": {
			input{userID, request.PayOrder{
				PaymentMethodID: 1,
				ProductItemInfo: []request.ProductItemInfo{
					{
						ProductItemID: 1,
						Count:         1,
					},
				},
			}},
			func(mr *orderMock.MockOrderRepository) {
				mr.EXPECT().Transactions(gomock.Any(), gomock.Any()).Return(errors.New("error"))
			},
			output{fmt.Errorf("failed to pay order: error")},
		},
	}

	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			tt.prepareMockFn(mockOrderRepo)
			err := order.PayOrder(ctx, tt.input.userID, tt.input.payOrder)
			if err != nil {
				assert.Equal(t, tt.want.wantErr.Error(), err.Error())
			} else {
				assert.Equal(t, tt.want.wantErr, err)
			}
		})
	}
}

func TestOrderUseCase_updateStockAndPayOrder(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockOrderRepo := orderMock.NewMockOrderRepository(ctrl)

	order := usecase.NewOrderUseCase(mockOrderRepo)

	ctx := echo.New().NewContext(nil, nil)
	userID := "1"

	type input struct {
		userID   string
		payOrder request.PayOrder
	}

	type output struct {
		wantErr error
	}

	tests := map[string]struct {
		input         input
		prepareMockFn func(mr *orderMock.MockOrderRepository)
		want          output
	}{
		"Normal Case: updateStockAndPayOrder": {
			input{userID, request.PayOrder{
				PaymentMethodID: 1,
				ProductItemInfo: []request.ProductItemInfo{
					{
						ProductItemID: 1,
						Count:         1,
					},
				},
			}},
			func(mr *orderMock.MockOrderRepository) {
				gomock.InOrder(
					mr.EXPECT().Transactions(gomock.Any(), gomock.Any()).DoAndReturn(
						func(_ echo.Context, f func(repo repoInterfaces.OrderRepository) error) error {
							gomock.InOrder(
								mr.EXPECT().UpdateProductItemStock(gomock.Any(), gomock.Any(), gomock.Any()).Return(uint(10), nil).Times(1),
								mr.EXPECT().PayOrder(gomock.Any(), gomock.Any()).Return(nil).Times(1),
								mr.EXPECT().SaveOrder(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil).Times(1),
							)

							return f(mr)
						}),
				)
			},
			output{nil},
		},
		"異常系: update error": {
			input{userID, request.PayOrder{
				PaymentMethodID: 1,
				ProductItemInfo: []request.ProductItemInfo{
					{
						ProductItemID: 1,
						Count:         1,
					},
				},
			}},
			func(mr *orderMock.MockOrderRepository) {
				gomock.InOrder(
					mr.EXPECT().Transactions(gomock.Any(), gomock.Any()).DoAndReturn(
						func(_ echo.Context, f func(repo repoInterfaces.OrderRepository) error) error {
							gomock.InOrder(
								mr.EXPECT().UpdateProductItemStock(gomock.Any(), gomock.Any(), gomock.Any()).Return(uint(0), errors.New("update error")).Times(1),
								mr.EXPECT().PayOrder(gomock.Any(), gomock.Any()).Return(errors.New("pay error")).Times(0),
								mr.EXPECT().SaveOrder(gomock.Any(), gomock.Any(), gomock.Any()).Return(errors.New("save error")).Times(0),
							)

							return f(mr)
						}),
				)
			},
			output{fmt.Errorf("failed to pay order: failed to update productItem stock to %d: %w", 0, errors.New("update error"))},
		},
		"異常系: pay error": {
			input{userID, request.PayOrder{
				PaymentMethodID: 1,
				ProductItemInfo: []request.ProductItemInfo{
					{
						ProductItemID: 1,
						Count:         1,
					},
				},
			}},
			func(mr *orderMock.MockOrderRepository) {
				gomock.InOrder(
					mr.EXPECT().Transactions(gomock.Any(), gomock.Any()).DoAndReturn(
						func(_ echo.Context, f func(repo repoInterfaces.OrderRepository) error) error {
							gomock.InOrder(
								mr.EXPECT().UpdateProductItemStock(gomock.Any(), gomock.Any(), gomock.Any()).Return(uint(10), nil).Times(1),
								mr.EXPECT().PayOrder(gomock.Any(), gomock.Any()).Return(errors.New("pay error")).Times(1),
								mr.EXPECT().SaveOrder(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil).Times(0),
							)

							return f(mr)
						}),
				)
			},
			output{fmt.Errorf("failed to pay order: %w", errors.New("pay error"))},
		},
		"異常系: save error": {
			input{userID, request.PayOrder{
				PaymentMethodID: 1,
				ProductItemInfo: []request.ProductItemInfo{
					{
						ProductItemID: 1,
						Count:         1,
					},
				},
			}},
			func(mr *orderMock.MockOrderRepository) {
				gomock.InOrder(
					mr.EXPECT().Transactions(gomock.Any(), gomock.Any()).DoAndReturn(
						func(_ echo.Context, f func(repo repoInterfaces.OrderRepository) error) error {
							gomock.InOrder(
								mr.EXPECT().UpdateProductItemStock(gomock.Any(), gomock.Any(), gomock.Any()).Return(uint(10), nil).Times(1),
								mr.EXPECT().PayOrder(gomock.Any(), gomock.Any()).Return(nil).Times(1),
								mr.EXPECT().SaveOrder(gomock.Any(), gomock.Any(), gomock.Any()).Return(errors.New("save error")).Times(1),
							)
							return f(mr)
						}),
				)
			},
			output{fmt.Errorf("failed to pay order: %w", errors.New("save error"))},
		},
	}

	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			tt.prepareMockFn(mockOrderRepo)
			err := order.PayOrder(ctx, tt.input.userID, tt.input.payOrder)
			if err != nil {
				assert.Equal(t, tt.want.wantErr.Error(), err.Error())
			} else {
				assert.Equal(t, tt.want.wantErr, err)
			}
		})
	}
}

func TestOrderUseCase_GetAllShopOrders(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockOrderRepo := orderMock.NewMockOrderRepository(ctrl)

	order := usecase.NewOrderUseCase(mockOrderRepo)

	ctx := echo.New().NewContext(nil, nil)
	userID := "1"
	pagination := request.Pagination{PageNumber: 1, Count: 10}

	type input struct {
		userID     string
		pagination request.Pagination
	}

	type output struct {
		wantOrderHistory []response.Order
		wantErr          error
	}

	tests := map[string]struct {
		input         input
		prepareMockFn func(mr *orderMock.MockOrderRepository)
		want          output
	}{
		"Normal Case: GetAllShopOrders": {
			input{userID, pagination},
			func(mr *orderMock.MockOrderRepository) {
				mr.EXPECT().GetShopOrders(gomock.Any(), userID, pagination).Return([]response.Order{}, nil)
			},
			output{[]response.Order{}, nil},
		},
		"Abnormal Case: GetAllShopOrders": {
			input{userID, pagination},
			func(mr *orderMock.MockOrderRepository) {
				mr.EXPECT().GetShopOrders(gomock.Any(), userID, pagination).Return(nil, errors.New("error"))
			},
			output{nil, fmt.Errorf("failed to find addresses: %w", errors.New("error"))},
		},
	}

	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			tt.prepareMockFn(mockOrderRepo)
			orderHistory, err := order.GetAllShopOrders(ctx, tt.input.userID, tt.input.pagination)
			if err != nil {
				assert.Equal(t, tt.want.wantErr.Error(), err.Error())
			} else {
				assert.Equal(t, tt.want.wantOrderHistory, orderHistory)
			}
		})
	}
}
