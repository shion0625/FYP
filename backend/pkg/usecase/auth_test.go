package usecase_test

import (
	"errors"
	"fmt"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/shion0625/FYP/backend/pkg/api/handler/request"
	"github.com/shion0625/FYP/backend/pkg/domain"
	authMock "github.com/shion0625/FYP/backend/pkg/repository/mock"
	tokenMock "github.com/shion0625/FYP/backend/pkg/service/mock/token"
	"github.com/shion0625/FYP/backend/pkg/service/token"
	"github.com/shion0625/FYP/backend/pkg/usecase"
	"github.com/shion0625/FYP/backend/pkg/usecase/interfaces"
	"github.com/stretchr/testify/assert"
	gomock "go.uber.org/mock/gomock"
)

func TestAuthUseCase_UserLogin(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAuthRepo := authMock.NewMockAuthRepository(ctrl)
	mockUserRepo := authMock.NewMockUserRepository(ctrl)
	mockToken := tokenMock.NewMockTokenService(ctrl)

	auth := usecase.NewAuthUseCase(mockUserRepo, mockToken, mockAuthRepo)

	req := httptest.NewRequest(echo.GET, "/", nil)
	rec := httptest.NewRecorder()
	ctx := echo.New().NewContext(req, rec)

	loginInfo := request.Login{Email: "test@example.com", Password: "xkaito0912@gmail.com"}

	type input struct {
		loginInfo request.Login
	}

	type output struct {
		wantErr error
	}

	tests := map[string]struct {
		input         input
		prepareMockFn func(mu *authMock.MockUserRepository)
		want          output
	}{
		"Normal Case: UserLogin with Email": {
			input{loginInfo},
			func(mu *authMock.MockUserRepository) {
				mu.EXPECT().FindUserByEmail(gomock.Any(), gomock.Any()).Return(domain.User{ID: "testUser", Password: "$2a$10$9rikUiKyu77gRc43i/qmxuSePlZ0zk/HwZKTJ11YSccISETYuxflW"}, nil)
			},
			output{nil},
		},
		"Normal Case: UserLogin with UserName": {
			input{request.Login{UserName: "testUser", Password: "xkaito0912@gmail.com"}},
			func(mu *authMock.MockUserRepository) {
				mu.EXPECT().FindUserByUserName(gomock.Any(), gomock.Any()).Return(domain.User{ID: "testUser", Password: "$2a$10$9rikUiKyu77gRc43i/qmxuSePlZ0zk/HwZKTJ11YSccISETYuxflW"}, nil)
			},
			output{nil},
		},
		"Normal Case: UserLogin with Phone": {
			input{request.Login{Phone: "1234567890", Password: "xkaito0912@gmail.com"}},
			func(mu *authMock.MockUserRepository) {
				mu.EXPECT().FindUserByPhoneNumber(gomock.Any(), gomock.Any()).Return(domain.User{ID: "testUser", Password: "$2a$10$9rikUiKyu77gRc43i/qmxuSePlZ0zk/HwZKTJ11YSccISETYuxflW"}, nil)
			},
			output{nil},
		},
		"Abnormal Case: UserLogin - Empty Login Credentials": {
			input{request.Login{}},
			func(mu *authMock.MockUserRepository) {
				// No mock expectations as the function should return early
			},
			output{usecase.ErrEmptyLoginCredentials},
		},
		"Abnormal Case: UserLogin - User not exist": {
			input{loginInfo},
			func(mu *authMock.MockUserRepository) {
				mu.EXPECT().FindUserByEmail(gomock.Any(), gomock.Any()).Return(domain.User{}, nil)
			},
			output{usecase.ErrUserNotExist},
		},
		"Abnormal Case: UserLogin - User blocked": {
			input{loginInfo},
			func(mu *authMock.MockUserRepository) {
				mu.EXPECT().FindUserByEmail(gomock.Any(), gomock.Any()).Return(domain.User{ID: "testUser", BlockStatus: true, Password: "hashedPassword"}, nil)
			},
			output{usecase.ErrUserBlocked},
		},
		"Abnormal Case: UserLogin with Email - User not found": {
			input{loginInfo},
			func(mu *authMock.MockUserRepository) {
				mu.EXPECT().FindUserByEmail(gomock.Any(), gomock.Any()).Return(domain.User{}, errors.New("error"))
			},
			output{fmt.Errorf("failed to find user from database: error")},
		},
		"Abnormal Case: UserLogin with Email - Wrong password": {
			input{loginInfo},
			func(mu *authMock.MockUserRepository) {
				mu.EXPECT().FindUserByEmail(gomock.Any(), gomock.Any()).Return(domain.User{ID: "testUser", Password: "wrongHashedPassword"}, nil)
			},
			output{usecase.ErrWrongPassword},
		},
	}

	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			tt.prepareMockFn(mockUserRepo)
			_, err := auth.UserLogin(ctx, tt.input.loginInfo)
			if err != nil {
				assert.Equal(t, tt.want.wantErr.Error(), err.Error())
			} else {
				assert.Equal(t, tt.want.wantErr, err)
			}
		})
	}
}

func TestAuthUseCase_UserSignUp(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAuthRepo := authMock.NewMockAuthRepository(ctrl)
	mockUserRepo := authMock.NewMockUserRepository(ctrl)
	mockToken := tokenMock.NewMockTokenService(ctrl)

	auth := usecase.NewAuthUseCase(mockUserRepo, mockToken, mockAuthRepo)

	req := httptest.NewRequest(echo.GET, "/", nil)
	rec := httptest.NewRecorder()
	ctx := echo.New().NewContext(req, rec)

	signUpDetails := domain.User{UserName: "testUser"}

	type input struct {
		signUpDetails domain.User
	}

	type output struct {
		wantErr error
	}

	tests := map[string]struct {
		input         input
		prepareMockFn func(mu *authMock.MockUserRepository)
		want          output
	}{
		"Normal Case: UserSignUp": {
			input{signUpDetails},
			func(mu *authMock.MockUserRepository) {
				mu.EXPECT().FindUserByUserNameEmailOrPhone(gomock.Any(), gomock.Any()).Return(domain.User{}, nil)
				mu.EXPECT().SaveUser(gomock.Any(), gomock.Any()).Return("success", nil)
			},
			output{nil},
		},
		"Abnormal Case: UserSignUp1": {
			input{signUpDetails},
			func(mu *authMock.MockUserRepository) {
				mu.EXPECT().FindUserByUserNameEmailOrPhone(gomock.Any(), gomock.Any()).Return(domain.User{}, errors.New("error"))
			},
			output{fmt.Errorf("failed to check user error")},
		},
		"Abnormal Case: UserSignUp": {
			input{signUpDetails},
			func(mu *authMock.MockUserRepository) {
				mu.EXPECT().FindUserByUserNameEmailOrPhone(gomock.Any(), gomock.Any()).Return(domain.User{}, nil)
				mu.EXPECT().SaveUser(gomock.Any(), gomock.Any()).Return("", errors.New("error"))
			},
			output{fmt.Errorf("failed to save user details: error")},
		},
		"Abnormal Case: UserSignUp - UserName already exists": {
			input{signUpDetails},
			func(mu *authMock.MockUserRepository) {
				existingUser := domain.User{UserName: "testUser"}
				mu.EXPECT().FindUserByUserNameEmailOrPhone(gomock.Any(), gomock.Any()).Return(existingUser, nil)
			},
			output{fmt.Errorf("failed to check user details already exist:\rUserName already exists\rEmail already exists\rPhone already exists.")},
		},
	}

	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			tt.prepareMockFn(mockUserRepo)
			_, err := auth.UserSignUp(ctx, tt.input.signUpDetails)
			if err != nil {
				assert.Equal(t, tt.want.wantErr.Error(), err.Error())
			} else {
				assert.Equal(t, tt.want.wantErr, err)
			}
		})
	}
}

func TestAuthUseCase_GenerateAccessToken(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAuthRepo := authMock.NewMockAuthRepository(ctrl)
	mockUserRepo := authMock.NewMockUserRepository(ctrl)
	mockToken := tokenMock.NewMockTokenService(ctrl)

	auth := usecase.NewAuthUseCase(mockUserRepo, mockToken, mockAuthRepo)

	req := httptest.NewRequest(echo.GET, "/", nil)
	rec := httptest.NewRecorder()
	ctx := echo.New().NewContext(req, rec)

	tokenParams := interfaces.GenerateTokenParams{UserID: "testUser", UserType: "user"}

	type input struct {
		tokenParams interfaces.GenerateTokenParams
	}

	type output struct {
		wantErr error
	}

	tests := map[string]struct {
		input         input
		prepareMockFn func(mt *tokenMock.MockTokenService)
		want          output
	}{
		"Normal Case: GenerateAccessToken": {
			input{tokenParams},
			func(mt *tokenMock.MockTokenService) {
				mt.EXPECT().GenerateToken(gomock.Any()).Return(token.GenerateTokenResponse{TokenString: "tokenString"}, nil)
			},
			output{nil},
		},
		"Abnormal Case: GenerateAccessToken": {
			input{tokenParams},
			func(mt *tokenMock.MockTokenService) {
				mt.EXPECT().GenerateToken(gomock.Any()).Return(token.GenerateTokenResponse{}, errors.New("error"))
			},
			output{fmt.Errorf("failed to generate access token: error")},
		},
	}

	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			tt.prepareMockFn(mockToken)
			_, err := auth.GenerateAccessToken(ctx, tt.input.tokenParams)
			if err != nil {
				assert.Equal(t, tt.want.wantErr.Error(), err.Error())
			} else {
				assert.Equal(t, tt.want.wantErr, err)
			}
		})
	}
}

func TestAuthUseCase_GenerateRefreshToken(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAuthRepo := authMock.NewMockAuthRepository(ctrl)
	mockUserRepo := authMock.NewMockUserRepository(ctrl)
	mockToken := tokenMock.NewMockTokenService(ctrl)

	auth := usecase.NewAuthUseCase(mockUserRepo, mockToken, mockAuthRepo)

	req := httptest.NewRequest(echo.GET, "/", nil)
	rec := httptest.NewRecorder()
	ctx := echo.New().NewContext(req, rec)

	tokenParams := interfaces.GenerateTokenParams{UserID: "testUser", UserType: "user"}

	type input struct {
		tokenParams interfaces.GenerateTokenParams
	}

	type output struct {
		wantErr error
	}

	tests := map[string]struct {
		input         input
		prepareMockFn func(mt *tokenMock.MockTokenService, ma *authMock.MockAuthRepository)
		want          output
	}{
		"Normal Case: GenerateRefreshToken": {
			input{tokenParams},
			func(mt *tokenMock.MockTokenService, ma *authMock.MockAuthRepository) {
				mt.EXPECT().GenerateToken(gomock.Any()).Return(token.GenerateTokenResponse{TokenString: "tokenString"}, nil)
				ma.EXPECT().SaveRefreshSession(gomock.Any(), gomock.Any()).Return(nil)
			},
			output{nil},
		},
		"Abnormal Case: GenerateRefreshToken - Failed to generate token": {
			input{tokenParams},
			func(mt *tokenMock.MockTokenService, ma *authMock.MockAuthRepository) {
				mt.EXPECT().GenerateToken(gomock.Any()).Return(token.GenerateTokenResponse{}, errors.New("error"))
			},
			output{fmt.Errorf("failed to generate refresh token: error")},
		},
		"Abnormal Case: GenerateRefreshToken - Failed to save refresh session": {
			input{tokenParams},
			func(mt *tokenMock.MockTokenService, ma *authMock.MockAuthRepository) {
				mt.EXPECT().GenerateToken(gomock.Any()).Return(token.GenerateTokenResponse{TokenString: "tokenString"}, nil)
				ma.EXPECT().SaveRefreshSession(gomock.Any(), gomock.Any()).Return(errors.New("error"))
			},
			output{fmt.Errorf("failed to save refresh session: error")},
		},
	}

	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			tt.prepareMockFn(mockToken, mockAuthRepo)
			_, err := auth.GenerateRefreshToken(ctx, tt.input.tokenParams)
			if err != nil {
				assert.Equal(t, tt.want.wantErr.Error(), err.Error())
			} else {
				assert.Equal(t, tt.want.wantErr, err)
			}
		})
	}
}

func TestAuthUseCase_VerifyAndGetRefreshTokenSession(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAuthRepo := authMock.NewMockAuthRepository(ctrl)
	mockUserRepo := authMock.NewMockUserRepository(ctrl)
	mockToken := tokenMock.NewMockTokenService(ctrl)

	auth := usecase.NewAuthUseCase(mockUserRepo, mockToken, mockAuthRepo)

	req := httptest.NewRequest(echo.GET, "/", nil)
	rec := httptest.NewRecorder()
	ctx := echo.New().NewContext(req, rec)

	refreshToken := "refreshToken"

	type input struct {
		refreshToken string
		usedFor      token.UserType
	}

	type output struct {
		wantErr error
	}

	inputOption := input{
		refreshToken: refreshToken,
		usedFor:      token.User,
	}
	tests := map[string]struct {
		input         input
		prepareMockFn func(mt *tokenMock.MockTokenService, ma *authMock.MockAuthRepository)
		want          output
	}{
		"Normal Case: VerifyAndGetRefreshTokenSession": {
			inputOption,
			func(mt *tokenMock.MockTokenService, ma *authMock.MockAuthRepository) {
				mt.EXPECT().VerifyToken(gomock.Any()).Return(token.VerifyTokenResponse{TokenID: "tokenID"}, nil)
				ma.EXPECT().FindRefreshSessionByTokenID(gomock.Any(), gomock.Any()).Return(domain.RefreshSession{TokenID: "tokenID", ExpireAt: time.Now().Add(time.Hour), IsBlocked: false}, nil)
			},
			output{nil},
		},
		"Abnormal Case: VerifyAndGetRefreshTokenSession - Failed to verify token": {
			inputOption,
			func(mt *tokenMock.MockTokenService, ma *authMock.MockAuthRepository) {
				mt.EXPECT().VerifyToken(gomock.Any()).Return(token.VerifyTokenResponse{}, errors.New("error"))
			},
			output{fmt.Errorf("failed to save refresh session: error")},
		},
		"Abnormal Case: VerifyAndGetRefreshTokenSession - Failed to find refresh session": {
			inputOption,
			func(mt *tokenMock.MockTokenService, ma *authMock.MockAuthRepository) {
				mt.EXPECT().VerifyToken(gomock.Any()).Return(token.VerifyTokenResponse{TokenID: "tokenID"}, nil)
				ma.EXPECT().FindRefreshSessionByTokenID(gomock.Any(), gomock.Any()).Return(domain.RefreshSession{}, errors.New("error"))
			},
			output{fmt.Errorf("failed to find refresh session by token ID: error")},
		},
		"Abnormal Case: VerifyAndGetRefreshTokenSession - Refresh session not exist": {
			inputOption,
			func(mt *tokenMock.MockTokenService, ma *authMock.MockAuthRepository) {
				mt.EXPECT().VerifyToken(gomock.Any()).Return(token.VerifyTokenResponse{TokenID: "tokenID"}, nil)
				ma.EXPECT().FindRefreshSessionByTokenID(gomock.Any(), gomock.Any()).Return(domain.RefreshSession{}, nil)
			},
			output{usecase.ErrRefreshSessionNotExist},
		},
		"Abnormal Case: VerifyAndGetRefreshTokenSession - Refresh session expired": {
			inputOption,
			func(mt *tokenMock.MockTokenService, ma *authMock.MockAuthRepository) {
				mt.EXPECT().VerifyToken(gomock.Any()).Return(token.VerifyTokenResponse{TokenID: "tokenID"}, nil)
				ma.EXPECT().FindRefreshSessionByTokenID(gomock.Any(), gomock.Any()).Return(domain.RefreshSession{TokenID: "tokenID", ExpireAt: time.Now().Add(-time.Hour)}, nil)
			},
			output{usecase.ErrRefreshSessionExpired},
		},
		"Abnormal Case: VerifyAndGetRefreshTokenSession - Refresh session blocked": {
			inputOption,
			func(mt *tokenMock.MockTokenService, ma *authMock.MockAuthRepository) {
				mt.EXPECT().VerifyToken(gomock.Any()).Return(token.VerifyTokenResponse{TokenID: "tokenID"}, nil)
				ma.EXPECT().FindRefreshSessionByTokenID(gomock.Any(), gomock.Any()).Return(domain.RefreshSession{TokenID: "tokenID", ExpireAt: time.Now().Add(time.Hour), IsBlocked: true}, nil)
			},
			output{usecase.ErrRefreshSessionBlocked},
		},
	}

	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			tt.prepareMockFn(mockToken, mockAuthRepo)
			_, err := auth.VerifyAndGetRefreshTokenSession(ctx, tt.input.refreshToken, tt.input.usedFor)
			if err != nil {
				assert.Equal(t, tt.want.wantErr.Error(), err.Error())
			} else {
				assert.Equal(t, tt.want.wantErr, err)
			}
		})
	}
}
