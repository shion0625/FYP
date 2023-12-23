package handler_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/shion0625/FYP/backend/pkg/api/handler"
	"github.com/shion0625/FYP/backend/pkg/api/handler/request"
	"github.com/shion0625/FYP/backend/pkg/api/handler/response"
	usecaseMock "github.com/shion0625/FYP/backend/pkg/usecase/mock"
	"github.com/stretchr/testify/assert"
	gomock "go.uber.org/mock/gomock"
)

func TestOrderHandler_PayOrder(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockOrderUseCase := usecaseMock.NewMockOrderUseCase(ctrl)

	orderHandler := handler.NewOrderHandler(mockOrderUseCase)

	type input struct {
		payOrderRequest request.PayOrder
	}

	type output struct {
		wantCode int
		wantStr  string
		wantErr  error
	}

	tests := map[string]struct {
		input         input
		prepareMockFn func(m *usecaseMock.MockOrderUseCase)
		want          output
	}{
		"Normal Case: PayOrder": {
			input{
				request.PayOrder{
					UserID:    "testUserID",
					AddressID: 1,
					ProductItemInfo: []request.ProductItemInfo{
						{
							ProductItemID: 1,
							Count:         1,
						},
					},
					TotalFee:        1000,
					PaymentMethodID: 1,
				},
			},
			func(m *usecaseMock.MockOrderUseCase) {
				m.EXPECT().PayOrder(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
			},
			output{http.StatusOK, "Successfully purchase order", nil},
		},
		"Abnormal Case: PayOrder": {
			input{
				request.PayOrder{
					UserID:    "testUserID",
					AddressID: 1,
					ProductItemInfo: []request.ProductItemInfo{
						{
							ProductItemID: 1,
							Count:         1,
						},
					},
					TotalFee:        1000,
					PaymentMethodID: 1,
				},
			},
			func(m *usecaseMock.MockOrderUseCase) {
				m.EXPECT().PayOrder(gomock.Any(), gomock.Any(), gomock.Any()).Return(errors.New("error"))
			},
			output{http.StatusInternalServerError, "Failed to PayOrder", nil},
		},
	}

	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			tt.prepareMockFn(mockOrderUseCase)

			body, _ := json.Marshal(tt.input.payOrderRequest)
			req := httptest.NewRequest(echo.POST, "/payOrder", bytes.NewBuffer(body))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()

			e := echo.New()
			e.Validator = &CustomValidator{validator: validator.New(validator.WithRequiredStructEnabled())}
			ctx := e.NewContext(req, rec)
			ctx.Set("userId", tt.input.payOrderRequest.UserID)

			err := orderHandler.PayOrder(ctx)
			assert.Nil(t, err)
			assert.Equal(t, tt.want.wantCode, rec.Code)
			assert.Contains(t, rec.Body.String(), tt.want.wantStr)
		})
	}
}

func TestOrderHandler_GetOrderHistory(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockOrderUseCase := usecaseMock.NewMockOrderUseCase(ctrl)

	orderHandler := handler.NewOrderHandler(mockOrderUseCase)

	type input struct {
		userId string
	}

	type output struct {
		wantCode int
		wantStr  string
		wantErr  error
	}

	tests := map[string]struct {
		input         input
		prepareMockFn func(m *usecaseMock.MockOrderUseCase)
		want          output
	}{
		"Normal Case: GetOrderHistory": {
			input{userId: "testUserId"},
			func(m *usecaseMock.MockOrderUseCase) {
				m.EXPECT().GetAllShopOrders(gomock.Any(), gomock.Any(), gomock.Any()).Return([]response.Order{}, nil)
			},
			output{http.StatusOK, "No shopOrders found", nil},
		},
		"Normal Case: GetOrderHistory with Orders": {
			input{userId: "testUserId"},

			func(m *usecaseMock.MockOrderUseCase) {
				m.EXPECT().GetAllShopOrders(gomock.Any(), gomock.Any(), gomock.Any()).Return([]response.Order{
					{
						UserID:      "testUserID",
						ShopOrderId: "testShopOrderId",
						ProductItemInfo: []response.ProductItemInfo{
							{},
						},
						Address:       response.Address{},
						TotalFee:      1000,
						PaymentMethod: response.PaymentMethod{},
					},
				}, nil)
			},
			output{http.StatusOK, "successfully addresses updated", nil},
		},
		"Abnormal Case: GetOrderHistory": {
			input{userId: "testUserId"},
			func(m *usecaseMock.MockOrderUseCase) {
				m.EXPECT().GetAllShopOrders(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, errors.New("error"))
			},
			output{http.StatusInternalServerError, "Failed to update user address", nil},
		},
	}

	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			tt.prepareMockFn(mockOrderUseCase)

			req := httptest.NewRequest(echo.GET, "/getOrderHistory", nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()

			e := echo.New()
			e.Validator = &CustomValidator{validator: validator.New(validator.WithRequiredStructEnabled())}
			ctx := e.NewContext(req, rec)
			ctx.Set("userId", tt.input.userId)

			err := orderHandler.GetOrderHistory(ctx)
			assert.Nil(t, err)
			assert.Equal(t, tt.want.wantCode, rec.Code)
			assert.Contains(t, rec.Body.String(), tt.want.wantStr)
		})
	}
}
