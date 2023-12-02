// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/usecase/interfaces/auth.go
//
// Generated by this command:
//
//	mockgen -source=pkg/usecase/interfaces/auth.go -destination=pkg/usecase/mock/auth_mock.go
//
// Package mock_interfaces is a generated GoMock package.
package mock_interfaces

import (
	reflect "reflect"

	v4 "github.com/labstack/echo/v4"
	request "github.com/shion0625/FYP/backend/pkg/api/handler/request"
	domain "github.com/shion0625/FYP/backend/pkg/domain"
	interfaces "github.com/shion0625/FYP/backend/pkg/usecase/interfaces"
	gomock "go.uber.org/mock/gomock"
)

// MockAuthUseCase is a mock of AuthUseCase interface.
type MockAuthUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockAuthUseCaseMockRecorder
}

// MockAuthUseCaseMockRecorder is the mock recorder for MockAuthUseCase.
type MockAuthUseCaseMockRecorder struct {
	mock *MockAuthUseCase
}

// NewMockAuthUseCase creates a new mock instance.
func NewMockAuthUseCase(ctrl *gomock.Controller) *MockAuthUseCase {
	mock := &MockAuthUseCase{ctrl: ctrl}
	mock.recorder = &MockAuthUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthUseCase) EXPECT() *MockAuthUseCaseMockRecorder {
	return m.recorder
}

// GenerateAccessToken mocks base method.
func (m *MockAuthUseCase) GenerateAccessToken(ctx v4.Context, tokenParams interfaces.GenerateTokenParams) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateAccessToken", ctx, tokenParams)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateAccessToken indicates an expected call of GenerateAccessToken.
func (mr *MockAuthUseCaseMockRecorder) GenerateAccessToken(ctx, tokenParams any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateAccessToken", reflect.TypeOf((*MockAuthUseCase)(nil).GenerateAccessToken), ctx, tokenParams)
}

// GenerateRefreshToken mocks base method.
func (m *MockAuthUseCase) GenerateRefreshToken(ctx v4.Context, tokenParams interfaces.GenerateTokenParams) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateRefreshToken", ctx, tokenParams)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateRefreshToken indicates an expected call of GenerateRefreshToken.
func (mr *MockAuthUseCaseMockRecorder) GenerateRefreshToken(ctx, tokenParams any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateRefreshToken", reflect.TypeOf((*MockAuthUseCase)(nil).GenerateRefreshToken), ctx, tokenParams)
}

// UserLogin mocks base method.
func (m *MockAuthUseCase) UserLogin(ctx v4.Context, loginInfo request.Login) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserLogin", ctx, loginInfo)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserLogin indicates an expected call of UserLogin.
func (mr *MockAuthUseCaseMockRecorder) UserLogin(ctx, loginInfo any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserLogin", reflect.TypeOf((*MockAuthUseCase)(nil).UserLogin), ctx, loginInfo)
}

// UserSignUp mocks base method.
func (m *MockAuthUseCase) UserSignUp(ctx v4.Context, signUpDetails domain.User) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserSignUp", ctx, signUpDetails)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserSignUp indicates an expected call of UserSignUp.
func (mr *MockAuthUseCaseMockRecorder) UserSignUp(ctx, signUpDetails any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserSignUp", reflect.TypeOf((*MockAuthUseCase)(nil).UserSignUp), ctx, signUpDetails)
}
