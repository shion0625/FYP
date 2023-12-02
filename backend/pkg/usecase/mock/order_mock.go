// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/usecase/interfaces/order.go
//
// Generated by this command:
//
//	mockgen -source=pkg/usecase/interfaces/order.go -destination=pkg/usecase/mock/order_mock.go
//
// Package mock_interfaces is a generated GoMock package.
package mock_interfaces

import (
	reflect "reflect"

	v4 "github.com/labstack/echo/v4"
	request "github.com/shion0625/FYP/backend/pkg/api/handler/request"
	gomock "go.uber.org/mock/gomock"
)

// MockOrderUseCase is a mock of OrderUseCase interface.
type MockOrderUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockOrderUseCaseMockRecorder
}

// MockOrderUseCaseMockRecorder is the mock recorder for MockOrderUseCase.
type MockOrderUseCaseMockRecorder struct {
	mock *MockOrderUseCase
}

// NewMockOrderUseCase creates a new mock instance.
func NewMockOrderUseCase(ctrl *gomock.Controller) *MockOrderUseCase {
	mock := &MockOrderUseCase{ctrl: ctrl}
	mock.recorder = &MockOrderUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrderUseCase) EXPECT() *MockOrderUseCaseMockRecorder {
	return m.recorder
}

// PayOrder mocks base method.
func (m *MockOrderUseCase) PayOrder(ctx v4.Context, product request.PayOrder) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PayOrder", ctx, product)
	ret0, _ := ret[0].(error)
	return ret0
}

// PayOrder indicates an expected call of PayOrder.
func (mr *MockOrderUseCaseMockRecorder) PayOrder(ctx, product any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PayOrder", reflect.TypeOf((*MockOrderUseCase)(nil).PayOrder), ctx, product)
}
