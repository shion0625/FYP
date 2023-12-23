// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/service/cloud/cloud.go
//
// Generated by this command:
//
//	mockgen -source=pkg/service/cloud/cloud.go -destination=pkg/service/mock/cloud/cloud_mock.go
//
// Package mock_cloud is a generated GoMock package.
package mock_cloud

import (
	multipart "mime/multipart"
	reflect "reflect"

	echo "github.com/labstack/echo/v4"
	gomock "go.uber.org/mock/gomock"
)

// MockCloudService is a mock of CloudService interface.
type MockCloudService struct {
	ctrl     *gomock.Controller
	recorder *MockCloudServiceMockRecorder
}

// MockCloudServiceMockRecorder is the mock recorder for MockCloudService.
type MockCloudServiceMockRecorder struct {
	mock *MockCloudService
}

// NewMockCloudService creates a new mock instance.
func NewMockCloudService(ctrl *gomock.Controller) *MockCloudService {
	mock := &MockCloudService{ctrl: ctrl}
	mock.recorder = &MockCloudServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCloudService) EXPECT() *MockCloudServiceMockRecorder {
	return m.recorder
}

// GetFileUrl mocks base method.
func (m *MockCloudService) GetFileUrl(ctx echo.Context, uploadID string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFileUrl", ctx, uploadID)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFileUrl indicates an expected call of GetFileUrl.
func (mr *MockCloudServiceMockRecorder) GetFileUrl(ctx, uploadID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFileUrl", reflect.TypeOf((*MockCloudService)(nil).GetFileUrl), ctx, uploadID)
}

// SaveFile mocks base method.
func (m *MockCloudService) SaveFile(ctx echo.Context, fileHeader *multipart.FileHeader) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveFile", ctx, fileHeader)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SaveFile indicates an expected call of SaveFile.
func (mr *MockCloudServiceMockRecorder) SaveFile(ctx, fileHeader any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveFile", reflect.TypeOf((*MockCloudService)(nil).SaveFile), ctx, fileHeader)
}