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
	"github.com/shion0625/FYP/backend/pkg/config"
	"github.com/shion0625/FYP/backend/pkg/domain"
	"github.com/shion0625/FYP/backend/pkg/usecase"
	usecaseMock "github.com/shion0625/FYP/backend/pkg/usecase/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	gomock "go.uber.org/mock/gomock"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return nil
}

func TestAuthHandler_UserLogin(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAuthUseCase := usecaseMock.NewMockAuthUseCase(ctrl)

	cfg, _ := config.LoadConfig()

	authHandler := handler.NewAuthHandler(mockAuthUseCase, cfg)

	type input struct {
		loginRequest request.Login
	}

	type output struct {
		wantCode int
		wantStr  string
		wantErr  error
	}

	tests := map[string]struct {
		input         input
		prepareMockFn func(m *usecaseMock.MockAuthUseCase)
		want          output
	}{
		"Normal Case: UserLogin": {
			input{request.Login{Email: "testUser", Password: "testPassword"}},
			func(m *usecaseMock.MockAuthUseCase) {
				m.EXPECT().UserLogin(gomock.Any(), gomock.Any()).Return("testUserID", nil)
				m.EXPECT().GenerateAccessToken(gomock.Any(), gomock.Any()).Return("accessToken", nil)
				m.EXPECT().GenerateRefreshToken(gomock.Any(), gomock.Any()).Return("refreshToken", nil)
			},
			output{http.StatusOK, "Login successful", nil},
		},
		"Abnormal Case: UserLogin1": {
			input{request.Login{Email: "testUser", Password: "testPassword"}},
			func(m *usecaseMock.MockAuthUseCase) {
				m.EXPECT().UserLogin(gomock.Any(), gomock.Any()).Return("testUserID", nil)
				m.EXPECT().GenerateAccessToken(gomock.Any(), gomock.Any()).Return("accessToken", errors.New("error"))
			},
			output{http.StatusInternalServerError, "Failed to generate access token", nil},
		},
		"Normal Case: UserLogin2": {
			input{request.Login{Email: "testUser", Password: "testPassword"}},
			func(m *usecaseMock.MockAuthUseCase) {
				m.EXPECT().UserLogin(gomock.Any(), gomock.Any()).Return("testUserID", nil)
				m.EXPECT().GenerateAccessToken(gomock.Any(), gomock.Any()).Return("accessToken", nil)
				m.EXPECT().GenerateRefreshToken(gomock.Any(), gomock.Any()).Return("refreshToken", errors.New("error"))
			},
			output{http.StatusInternalServerError, "Failed to generate refresh token", nil},
		},
		"Abnormal Case: UserLogin": {
			input{request.Login{UserName: "testUser", Password: "testPassword"}},
			func(m *usecaseMock.MockAuthUseCase) {
				m.EXPECT().UserLogin(gomock.Any(), gomock.Any()).Return("", errors.New("error"))
			},
			output{http.StatusInternalServerError, "Login failed", nil},
		},
	}

	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			tt.prepareMockFn(mockAuthUseCase)

			body, _ := json.Marshal(tt.input.loginRequest)
			req := httptest.NewRequest(echo.POST, "/login", bytes.NewBuffer(body))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()

			e := echo.New()
			e.Validator = &CustomValidator{validator: validator.New(validator.WithRequiredStructEnabled())}
			ctx := e.NewContext(req, rec)

			err := authHandler.UserLogin(ctx)
			require.NoError(t, err)

			if tt.want.wantErr != nil {
				assert.Equal(t, tt.want.wantCode, rec.Code)
				assert.Contains(t, rec.Body.String(), tt.want.wantErr.Error())
			} else {
				assert.Equal(t, tt.want.wantCode, rec.Code)
				assert.Contains(t, rec.Body.String(), tt.want.wantStr)
			}
		})
	}
}

func TestAuthHandler_UserSignUp(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAuthUseCase := usecaseMock.NewMockAuthUseCase(ctrl)

	cfg, _ := config.LoadConfig()

	authHandler := handler.NewAuthHandler(mockAuthUseCase, cfg)

	type input struct {
		signUpRequest request.SignUp
	}

	type output struct {
		wantCode int
		wantStr  string
		wantErr  error
	}

	tests := map[string]struct {
		input         input
		prepareMockFn func(m *usecaseMock.MockAuthUseCase)
		want          output
	}{
		"Normal Case: UserSignUp": {
			input{request.SignUp{
				UserName:        "testUser",
				FirstName:       "Test",
				LastName:        "User",
				Age:             30,
				Email:           "testUser@example.com",
				Phone:           "+1234567890",
				Password:        "testPassword",
				ConfirmPassword: "testPassword",
			}},
			func(m *usecaseMock.MockAuthUseCase) {
				m.EXPECT().UserSignUp(gomock.Any(), gomock.Any()).Return("testUserID", nil)
			},
			output{http.StatusCreated, "Account created successfully", nil},
		},
		"Abnormal Case: UserSignUp": {
			input{request.SignUp{
				UserName:        "testUser",
				FirstName:       "Test",
				LastName:        "User",
				Age:             30,
				Email:           "testUser@example.com",
				Phone:           "+1234567890",
				Password:        "testPassword",
				ConfirmPassword: "testPassword",
			}},
			func(m *usecaseMock.MockAuthUseCase) {
				m.EXPECT().UserSignUp(gomock.Any(), gomock.Any()).Return("", errors.New("error"))
			},
			output{http.StatusInternalServerError, "Signup failed", nil},
		},
	}

	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			tt.prepareMockFn(mockAuthUseCase)

			body, _ := json.Marshal(tt.input.signUpRequest)
			req := httptest.NewRequest(echo.POST, "/signup", bytes.NewBuffer(body))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()

			e := echo.New()
			e.Validator = &CustomValidator{validator: validator.New(validator.WithRequiredStructEnabled())}
			ctx := e.NewContext(req, rec)

			err := authHandler.UserSignUp(ctx)
			require.NoError(t, err)

			if tt.want.wantErr != nil {
				assert.Equal(t, tt.want.wantCode, rec.Code)
				assert.Contains(t, rec.Body.String(), tt.want.wantErr.Error())
			} else {
				assert.Equal(t, tt.want.wantCode, rec.Code)
				assert.Contains(t, rec.Body.String(), tt.want.wantStr)
			}
		})
	}
}

func TestAuthHandler_UserRenewAccessToken(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAuthUseCase := usecaseMock.NewMockAuthUseCase(ctrl)

	cfg, _ := config.LoadConfig()

	authHandler := handler.NewAuthHandler(mockAuthUseCase, cfg)

	type input struct {
		refreshTokenRequest request.RefreshToken
	}

	type output struct {
		wantCode int
		wantStr  string
		wantErr  error
	}

	tests := map[string]struct {
		input         input
		prepareMockFn func(m *usecaseMock.MockAuthUseCase)
		want          output
	}{
		"Normal Case: UserRenewAccessToken": {
			input{request.RefreshToken{RefreshToken: "validRefreshToken"}},
			func(m *usecaseMock.MockAuthUseCase) {
				m.EXPECT().VerifyAndGetRefreshTokenSession(gomock.Any(), gomock.Any(), gomock.Any()).Return(domain.RefreshSession{UserID: "testUserID"}, nil)
				m.EXPECT().GenerateAccessToken(gomock.Any(), gomock.Any()).Return("newAccessToken", nil)
			},
			output{http.StatusOK, "Access token generated successfully using refresh token", nil},
		},
		"Abnormal Case: UserRenewAccessToken": {
			input{request.RefreshToken{RefreshToken: "invalidRefreshToken"}},
			func(m *usecaseMock.MockAuthUseCase) {
				m.EXPECT().VerifyAndGetRefreshTokenSession(gomock.Any(), gomock.Any(), gomock.Any()).Return(domain.RefreshSession{}, errors.New("error"))
			},
			output{http.StatusInternalServerError, "Failed to verify refresh token", nil},
		},
		"Abnormal Case: UserRenewAccessToken1": {
			input{request.RefreshToken{RefreshToken: "validRefreshToken"}},
			func(m *usecaseMock.MockAuthUseCase) {
				m.EXPECT().VerifyAndGetRefreshTokenSession(gomock.Any(), gomock.Any(), gomock.Any()).Return(domain.RefreshSession{UserID: "testUserID"}, nil)
				m.EXPECT().GenerateAccessToken(gomock.Any(), gomock.Any()).Return("newAccessToken", errors.New("error"))
			},
			output{http.StatusInternalServerError, "Failed to generate access token", nil},
		},
		"Abnormal Case: Invalid Refresh Token": {
			input{request.RefreshToken{RefreshToken: "invalidRefreshToken"}},
			func(m *usecaseMock.MockAuthUseCase) {
				m.EXPECT().VerifyAndGetRefreshTokenSession(gomock.Any(), gomock.Any(), gomock.Any()).Return(domain.RefreshSession{}, usecase.ErrInvalidRefreshToken)
			},
			output{http.StatusUnauthorized, "Failed to verify refresh token", nil},
		},
		"Abnormal Case: Refresh Session Not Exist": {
			input{request.RefreshToken{RefreshToken: "invalidRefreshToken"}},
			func(m *usecaseMock.MockAuthUseCase) {
				m.EXPECT().VerifyAndGetRefreshTokenSession(gomock.Any(), gomock.Any(), gomock.Any()).Return(domain.RefreshSession{}, usecase.ErrRefreshSessionNotExist)
			},
			output{http.StatusNotFound, "Failed to verify refresh token", nil},
		},
		"Abnormal Case: Refresh Session Expired": {
			input{request.RefreshToken{RefreshToken: "expiredRefreshToken"}},
			func(m *usecaseMock.MockAuthUseCase) {
				m.EXPECT().VerifyAndGetRefreshTokenSession(gomock.Any(), gomock.Any(), gomock.Any()).Return(domain.RefreshSession{}, usecase.ErrRefreshSessionExpired)
			},
			output{http.StatusGone, "Failed to verify refresh token", nil},
		},
		"Abnormal Case: Refresh Session Blocked": {
			input{request.RefreshToken{RefreshToken: "blockedRefreshToken"}},
			func(m *usecaseMock.MockAuthUseCase) {
				m.EXPECT().VerifyAndGetRefreshTokenSession(gomock.Any(), gomock.Any(), gomock.Any()).Return(domain.RefreshSession{}, usecase.ErrRefreshSessionBlocked)
			},
			output{http.StatusForbidden, "Failed to verify refresh token", nil},
		},
		"Abnormal Case: Other Error": {
			input{request.RefreshToken{RefreshToken: "otherError"}},
			func(m *usecaseMock.MockAuthUseCase) {
				m.EXPECT().VerifyAndGetRefreshTokenSession(gomock.Any(), gomock.Any(), gomock.Any()).Return(domain.RefreshSession{}, errors.New("other error"))
			},
			output{http.StatusInternalServerError, "Failed to verify refresh token", nil},
		},
	}

	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			tt.prepareMockFn(mockAuthUseCase)

			body, _ := json.Marshal(tt.input.refreshTokenRequest)
			req := httptest.NewRequest(echo.POST, "/renewAccessToken", bytes.NewBuffer(body))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()

			e := echo.New()
			e.Validator = &CustomValidator{validator: validator.New(validator.WithRequiredStructEnabled())}
			ctx := e.NewContext(req, rec)

			err := authHandler.UserRenewAccessToken()(ctx)
			require.NoError(t, err)

			if tt.want.wantErr != nil {
				assert.Equal(t, tt.want.wantCode, rec.Code)
				assert.Contains(t, rec.Body.String(), tt.want.wantErr.Error())
			} else {
				assert.Equal(t, tt.want.wantCode, rec.Code)
				assert.Contains(t, rec.Body.String(), tt.want.wantStr)
			}
		})
	}
}
