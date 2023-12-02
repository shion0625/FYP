// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/usecase/interfaces/user.go
//
// Generated by this command:
//
//	mockgen -source=pkg/usecase/interfaces/user.go -destination=pkg/usecase/mock/user_mock.go
//
// Package mock_interfaces is a generated GoMock package.
package mock_interfaces

import (
	reflect "reflect"

	v4 "github.com/labstack/echo/v4"
	request "github.com/shion0625/FYP/backend/pkg/api/handler/request"
	response "github.com/shion0625/FYP/backend/pkg/api/handler/response"
	domain "github.com/shion0625/FYP/backend/pkg/domain"
	gomock "go.uber.org/mock/gomock"
)

// MockUserUseCase is a mock of UserUseCase interface.
type MockUserUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockUserUseCaseMockRecorder
}

// MockUserUseCaseMockRecorder is the mock recorder for MockUserUseCase.
type MockUserUseCaseMockRecorder struct {
	mock *MockUserUseCase
}

// NewMockUserUseCase creates a new mock instance.
func NewMockUserUseCase(ctrl *gomock.Controller) *MockUserUseCase {
	mock := &MockUserUseCase{ctrl: ctrl}
	mock.recorder = &MockUserUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserUseCase) EXPECT() *MockUserUseCaseMockRecorder {
	return m.recorder
}

// FindAddresses mocks base method.
func (m *MockUserUseCase) FindAddresses(ctx v4.Context, userID string) ([]response.Address, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAddresses", ctx, userID)
	ret0, _ := ret[0].([]response.Address)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAddresses indicates an expected call of FindAddresses.
func (mr *MockUserUseCaseMockRecorder) FindAddresses(ctx, userID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAddresses", reflect.TypeOf((*MockUserUseCase)(nil).FindAddresses), ctx, userID)
}

// FindProfile mocks base method.
func (m *MockUserUseCase) FindProfile(ctx v4.Context, userId string) (domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindProfile", ctx, userId)
	ret0, _ := ret[0].(domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindProfile indicates an expected call of FindProfile.
func (mr *MockUserUseCaseMockRecorder) FindProfile(ctx, userId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindProfile", reflect.TypeOf((*MockUserUseCase)(nil).FindProfile), ctx, userId)
}

// SaveAddress mocks base method.
func (m *MockUserUseCase) SaveAddress(ctx v4.Context, userID string, address domain.Address, isDefault bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveAddress", ctx, userID, address, isDefault)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveAddress indicates an expected call of SaveAddress.
func (mr *MockUserUseCaseMockRecorder) SaveAddress(ctx, userID, address, isDefault any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveAddress", reflect.TypeOf((*MockUserUseCase)(nil).SaveAddress), ctx, userID, address, isDefault)
}

// UpdateAddress mocks base method.
func (m *MockUserUseCase) UpdateAddress(ctx v4.Context, addressBody request.EditAddress, userID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAddress", ctx, addressBody, userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAddress indicates an expected call of UpdateAddress.
func (mr *MockUserUseCaseMockRecorder) UpdateAddress(ctx, addressBody, userID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAddress", reflect.TypeOf((*MockUserUseCase)(nil).UpdateAddress), ctx, addressBody, userID)
}

// UpdateProfile mocks base method.
func (m *MockUserUseCase) UpdateProfile(ctx v4.Context, user domain.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProfile", ctx, user)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateProfile indicates an expected call of UpdateProfile.
func (mr *MockUserUseCaseMockRecorder) UpdateProfile(ctx, user any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProfile", reflect.TypeOf((*MockUserUseCase)(nil).UpdateProfile), ctx, user)
}
