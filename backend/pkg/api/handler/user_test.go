package handler_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/shion0625/FYP/backend/pkg/api/handler"
	"github.com/shion0625/FYP/backend/pkg/api/handler/request"
	"github.com/shion0625/FYP/backend/pkg/api/handler/response"
	"github.com/shion0625/FYP/backend/pkg/domain"
	usecaseMock "github.com/shion0625/FYP/backend/pkg/usecase/mock"
	"github.com/stretchr/testify/assert"
	gomock "go.uber.org/mock/gomock"
)

func TestUserHandler_GetProfile(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserUseCase := usecaseMock.NewMockUserUseCase(ctrl)

	userHandler := handler.NewUserHandler(mockUserUseCase)

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
		prepareMockFn func(m *usecaseMock.MockUserUseCase)
		want          output
	}{
		"Normal Case: GetProfile": {
			input{userId: "testUserId"},
			func(m *usecaseMock.MockUserUseCase) {
				m.EXPECT().FindProfile(gomock.Any(), gomock.Any()).Return(domain.User{
					ID:          "fea",
					GoogleImage: "test_image.jpg",
					FirstName:   "Test",
					LastName:    "User",
					Age:         20,
					Email:       "test@example.com",
					UserName:    "testuser",
					Phone:       "1234567890",
					BlockStatus: false,
					CreatedAt:   time.Now(),
					UpdatedAt:   time.Now(),
				}, nil)
			},
			output{http.StatusOK, "Successfully retrieved user details", nil},
		},
		"Abnormal Case: GetProfile": {
			input{userId: "testUserId"},
			func(m *usecaseMock.MockUserUseCase) {
				m.EXPECT().FindProfile(gomock.Any(), gomock.Any()).Return(domain.User{}, errors.New("error"))
			},
			output{http.StatusInternalServerError, "Failed to retrieve user details", nil},
		},
	}

	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			tt.prepareMockFn(mockUserUseCase)

			req := httptest.NewRequest(echo.GET, "/getProfile", nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()

			e := echo.New()
			e.Validator = &CustomValidator{validator: validator.New(validator.WithRequiredStructEnabled())}
			ctx := e.NewContext(req, rec)
			ctx.Set("userId", tt.input.userId)

			err := userHandler.GetProfile(ctx)
			assert.Nil(t, err)
			assert.Equal(t, tt.want.wantCode, rec.Code)
			assert.Contains(t, rec.Body.String(), tt.want.wantStr)
		})
	}
}

func TestUserHandler_UpdateProfile(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserUseCase := usecaseMock.NewMockUserUseCase(ctrl)

	userHandler := handler.NewUserHandler(mockUserUseCase)

	type input struct {
		userId string
		body   request.EditUser
	}

	type output struct {
		wantCode int
		wantStr  string
		wantErr  error
	}

	tests := map[string]struct {
		input         input
		prepareMockFn func(m *usecaseMock.MockUserUseCase)
		want          output
	}{
		"Normal Case: UpdateProfile": {
			input{
				userId: "testUserId",
				body: request.EditUser{
					UserName:        "testUser",
					FirstName:       "Test",
					LastName:        "User",
					Age:             30,
					Email:           "test@example.com",
					Phone:           "+1234567890",
					Password:        "password123",
					ConfirmPassword: "password123",
				},
			},
			func(m *usecaseMock.MockUserUseCase) {
				m.EXPECT().UpdateProfile(gomock.Any(), gomock.Any()).Return(nil)
			},
			output{http.StatusOK, "Successfully profile updated", nil},
		},
		"Abnormal Case: UpdateProfile": {
			input{
				userId: "testUserId",
				body: request.EditUser{
					UserName:        "testUser",
					FirstName:       "Test",
					LastName:        "User",
					Age:             30,
					Email:           "test@example.com",
					Phone:           "+1234567890",
					Password:        "password123",
					ConfirmPassword: "password123",
				},
			},
			func(m *usecaseMock.MockUserUseCase) {
				m.EXPECT().UpdateProfile(gomock.Any(), gomock.Any()).Return(errors.New("error"))
			},
			output{http.StatusInternalServerError, "Failed to update profile", nil},
		},
	}

	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			tt.prepareMockFn(mockUserUseCase)

			body, _ := json.Marshal(tt.input.body)
			req := httptest.NewRequest(echo.PUT, "/updateProfile", bytes.NewBuffer(body))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()

			e := echo.New()
			e.Validator = &CustomValidator{validator: validator.New(validator.WithRequiredStructEnabled())}
			ctx := e.NewContext(req, rec)
			ctx.Set("userId", tt.input.userId)

			err := userHandler.UpdateProfile(ctx)
			assert.Nil(t, err)
			assert.Equal(t, tt.want.wantCode, rec.Code)
			assert.Contains(t, rec.Body.String(), tt.want.wantStr)
		})
	}
}

func TestUserHandler_SaveAddress(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserUseCase := usecaseMock.NewMockUserUseCase(ctrl)

	userHandler := handler.NewUserHandler(mockUserUseCase)

	type input struct {
		userId string
		body   request.Address
	}

	type output struct {
		wantCode int
		wantStr  string
		wantErr  error
	}

	tests := map[string]struct {
		input         input
		prepareMockFn func(m *usecaseMock.MockUserUseCase)
		want          output
	}{
		"Normal Case: SaveAddress": {
			input{
				userId: "testUserId",
				body: request.Address{
					Name:        "Test User",
					PhoneNumber: "+1234567890",
					House:       "123 Test St",
					Area:        "Test Area",
					LandMark:    "Test Landmark",
					City:        "Test City",
					Pincode:     "12345",
					CountryName: "Test Country",
					IsDefault:   new(bool),
				},
			},
			func(m *usecaseMock.MockUserUseCase) {
				m.EXPECT().SaveAddress(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
			},
			output{http.StatusCreated, "Successfully address saved", nil},
		},
		"Abnormal Case: SaveAddress": {
			input{
				userId: "testUserId",
				body: request.Address{
					Name:        "Test User",
					PhoneNumber: "+1234567890",
					House:       "123 Test St",
					Area:        "Test Area",
					LandMark:    "Test Landmark",
					City:        "Test City",
					Pincode:     "12345",
					CountryName: "Test Country",
					IsDefault:   new(bool),
				},
			},
			func(m *usecaseMock.MockUserUseCase) {
				m.EXPECT().SaveAddress(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(errors.New("error"))
			},
			output{http.StatusInternalServerError, "Failed to save address", nil},
		},
	}

	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			tt.prepareMockFn(mockUserUseCase)

			body, _ := json.Marshal(tt.input.body)
			req := httptest.NewRequest(echo.POST, "/saveAddress", bytes.NewBuffer(body))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()

			e := echo.New()
			e.Validator = &CustomValidator{validator: validator.New(validator.WithRequiredStructEnabled())}
			ctx := e.NewContext(req, rec)
			ctx.Set("userId", tt.input.userId)

			err := userHandler.SaveAddress(ctx)
			assert.Nil(t, err)
			assert.Equal(t, tt.want.wantCode, rec.Code)
			assert.Contains(t, rec.Body.String(), tt.want.wantStr)
		})
	}
}

func TestUserHandler_GetAllAddresses(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserUseCase := usecaseMock.NewMockUserUseCase(ctrl)

	userHandler := handler.NewUserHandler(mockUserUseCase)

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
		prepareMockFn func(m *usecaseMock.MockUserUseCase)
		want          output
	}{
		"Normal Case: GetAllAddresses": {
			input{userId: "testUserId"},
			func(m *usecaseMock.MockUserUseCase) {
				m.EXPECT().FindAddresses(gomock.Any(), gomock.Any()).Return([]response.Address{}, nil)
			},
			output{http.StatusOK, "Successfully retrieved all user addresses", nil},
		},
		"Abnormal Case: GetAllAddresses": {
			input{userId: "testUserId"},
			func(m *usecaseMock.MockUserUseCase) {
				m.EXPECT().FindAddresses(gomock.Any(), gomock.Any()).Return([]response.Address{}, errors.New("error"))
			},
			output{http.StatusInternalServerError, "Failed to get user addresses", nil},
		},
	}

	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			tt.prepareMockFn(mockUserUseCase)

			req := httptest.NewRequest(echo.GET, "/getAllAddresses", nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()

			e := echo.New()
			e.Validator = &CustomValidator{validator: validator.New(validator.WithRequiredStructEnabled())}
			ctx := e.NewContext(req, rec)
			ctx.Set("userId", tt.input.userId)

			err := userHandler.GetAllAddresses(ctx)
			assert.Nil(t, err)
			assert.Equal(t, tt.want.wantCode, rec.Code)
			assert.Contains(t, rec.Body.String(), tt.want.wantStr)
		})
	}
}

func TestUserHandler_GetAddressById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserUseCase := usecaseMock.NewMockUserUseCase(ctrl)

	userHandler := handler.NewUserHandler(mockUserUseCase)

	type input struct {
		userId    string
		addressId string
	}

	type output struct {
		wantCode int
		wantStr  string
		wantErr  error
	}

	tests := map[string]struct {
		input         input
		prepareMockFn func(m *usecaseMock.MockUserUseCase)
		want          output
	}{
		"Normal Case: GetAddressById": {
			input{userId: "testUserId", addressId: "1"},
			func(m *usecaseMock.MockUserUseCase) {
				m.EXPECT().FindAddress(gomock.Any(), gomock.Any(), gomock.Any()).Return(domain.Address{}, nil)
			},
			output{http.StatusOK, "Successfully retrieved all user addresses", nil},
		},
		"Abnormal Case: GetAddressById": {
			input{userId: "testUserId", addressId: "1"},
			func(m *usecaseMock.MockUserUseCase) {
				m.EXPECT().FindAddress(gomock.Any(), gomock.Any(), gomock.Any()).Return(domain.Address{}, errors.New("error"))
			},
			output{http.StatusInternalServerError, "Failed to get user addresses", nil},
		},
	}

	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			tt.prepareMockFn(mockUserUseCase)

			req := httptest.NewRequest(echo.GET, "/getAddressById/"+tt.input.addressId, nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()

			e := echo.New()
			e.Validator = &CustomValidator{validator: validator.New(validator.WithRequiredStructEnabled())}
			ctx := e.NewContext(req, rec)
			ctx.Set("userId", tt.input.userId)
			ctx.SetParamNames("address_id")
			ctx.SetParamValues(tt.input.addressId)

			err := userHandler.GetAddressById(ctx)
			assert.Nil(t, err)
			assert.Equal(t, tt.want.wantCode, rec.Code)
			assert.Contains(t, rec.Body.String(), tt.want.wantStr)
		})
	}
}

func TestUserHandler_UpdateAddress(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserUseCase := usecaseMock.NewMockUserUseCase(ctrl)

	userHandler := handler.NewUserHandler(mockUserUseCase)

	type input struct {
		userId string
		body   request.EditAddress
	}
	type output struct {
		wantCode int
		wantStr  string
		wantErr  error
	}

	tests := map[string]struct {
		input         input
		prepareMockFn func(m *usecaseMock.MockUserUseCase)
		want          output
	}{
		"Normal Case: UpdateAddress": {
			input{
				userId: "testUserId",
				body: request.EditAddress{
					ID:          1,
					Name:        "Updated User",
					PhoneNumber: "+0987654321",
					House:       "321 Updated St",
					Area:        "Updated Area",
					LandMark:    "Updated Landmark",
					City:        "Updated City",
					Pincode:     "54321",
					CountryName: "Updated Country",
					IsDefault:   new(bool),
				},
			},
			func(m *usecaseMock.MockUserUseCase) {
				m.EXPECT().UpdateAddress(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
			},
			output{http.StatusOK, "successfully addresses updated", nil},
		},
		"Abnormal Case: UpdateAddress": {
			input{
				userId: "testUserId",
				body: request.EditAddress{
					ID:          1,
					Name:        "Updated User",
					PhoneNumber: "+0987654321",
					House:       "321 Updated St",
					Area:        "Updated Area",
					LandMark:    "Updated Landmark",
					City:        "Updated City",
					Pincode:     "54321",
					CountryName: "Updated Country",
					IsDefault:   new(bool),
				},
			},
			func(m *usecaseMock.MockUserUseCase) {
				m.EXPECT().UpdateAddress(gomock.Any(), gomock.Any(), gomock.Any()).Return(errors.New("error"))
			},
			output{http.StatusInternalServerError, "Failed to update user address", nil},
		},
	}

	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			tt.prepareMockFn(mockUserUseCase)

			body, _ := json.Marshal(tt.input.body)
			req := httptest.NewRequest(echo.PUT, "/updateAddress", bytes.NewBuffer(body))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()

			e := echo.New()
			e.Validator = &CustomValidator{validator: validator.New(validator.WithRequiredStructEnabled())}
			ctx := e.NewContext(req, rec)
			ctx.Set("userId", tt.input.userId)

			err := userHandler.UpdateAddress(ctx)
			assert.Nil(t, err)
			assert.Equal(t, tt.want.wantCode, rec.Code)
			assert.Contains(t, rec.Body.String(), tt.want.wantStr)
		})
	}
}

func TestUserHandler_GetAllPaymentMethods(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserUseCase := usecaseMock.NewMockUserUseCase(ctrl)

	userHandler := handler.NewUserHandler(mockUserUseCase)

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
		prepareMockFn func(m *usecaseMock.MockUserUseCase)
		want          output
	}{
		"Normal Case: GetAllPaymentMethods": {
			input{userId: "testUserId"},
			func(m *usecaseMock.MockUserUseCase) {
				m.EXPECT().FindPaymentMethods(gomock.Any(), gomock.Any()).Return([]response.PaymentMethod{{
					ID:          1,
					Number:      "1234567812345678",
					CardCompany: "Test Card Company",
				}}, nil)
			},
			output{http.StatusOK, "Successfully retrieved all user addresses", nil},
		},
		"Abnormal Case: GetAllPaymentMethods": {
			input{userId: "testUserId"},
			func(m *usecaseMock.MockUserUseCase) {
				m.EXPECT().FindPaymentMethods(gomock.Any(), gomock.Any()).Return([]response.PaymentMethod{}, errors.New("error"))
			},
			output{http.StatusInternalServerError, "GetAllPaymentMethods", nil},
		},
	}

	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			tt.prepareMockFn(mockUserUseCase)

			req := httptest.NewRequest(echo.GET, "/getAllPaymentMethods", nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()

			e := echo.New()
			e.Validator = &CustomValidator{validator: validator.New(validator.WithRequiredStructEnabled())}
			ctx := e.NewContext(req, rec)
			ctx.Set("userId", tt.input.userId)

			err := userHandler.GetAllPaymentMethods(ctx)
			assert.Nil(t, err)
			assert.Equal(t, tt.want.wantCode, rec.Code)
			assert.Contains(t, rec.Body.String(), tt.want.wantStr)
		})
	}
}

func TestUserHandler_SavePaymentMethod(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserUseCase := usecaseMock.NewMockUserUseCase(ctrl)

	userHandler := handler.NewUserHandler(mockUserUseCase)

	type input struct {
		userId string
		body   request.PaymentMethod
	}
	type output struct {
		wantCode int
		wantStr  string
		wantErr  error
	}

	tests := map[string]struct {
		input         input
		prepareMockFn func(m *usecaseMock.MockUserUseCase)
		want          output
	}{
		"Normal Case: SavePaymentMethod": {
			input{
				userId: "testUserId",
				body: request.PaymentMethod{
					Number: "1234567812345678",
					Name:   "Test User",
					Expiry: "12/24",
					Cvc:    "123",
				},
			},
			func(m *usecaseMock.MockUserUseCase) {
				m.EXPECT().SavePaymentMethod(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
			},
			output{http.StatusOK, "successfully addresses updated", nil},
		},
		"Abnormal Case: SavePaymentMethod": {
			input{
				userId: "testUserId",
				body: request.PaymentMethod{
					Number: "1234567812345678",
					Name:   "Test User",
					Expiry: "12/24",
					Cvc:    "123",
				},
			},
			func(m *usecaseMock.MockUserUseCase) {
				m.EXPECT().SavePaymentMethod(gomock.Any(), gomock.Any(), gomock.Any()).Return(errors.New("error"))
			},
			output{http.StatusInternalServerError, "Failed to update user address", nil},
		},
	}

	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			tt.prepareMockFn(mockUserUseCase)

			body, _ := json.Marshal(tt.input.body)
			req := httptest.NewRequest(echo.POST, "/savePaymentMethod", bytes.NewBuffer(body))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()

			e := echo.New()
			e.Validator = &CustomValidator{validator: validator.New(validator.WithRequiredStructEnabled())}
			ctx := e.NewContext(req, rec)
			ctx.Set("userId", tt.input.userId)

			err := userHandler.SavePaymentMethod(ctx)
			assert.Nil(t, err)
			assert.Equal(t, tt.want.wantCode, rec.Code)
			assert.Contains(t, rec.Body.String(), tt.want.wantStr)
		})
	}
}
