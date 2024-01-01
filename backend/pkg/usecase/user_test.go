package usecase_test

import (
	"errors"
	"fmt"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/shion0625/FYP/backend/pkg/api/handler/request"
	"github.com/shion0625/FYP/backend/pkg/api/handler/response"
	"github.com/shion0625/FYP/backend/pkg/config"
	"github.com/shion0625/FYP/backend/pkg/domain"
	userMock "github.com/shion0625/FYP/backend/pkg/repository/mock"
	"github.com/shion0625/FYP/backend/pkg/usecase"
	"github.com/stretchr/testify/assert"
	gomock "go.uber.org/mock/gomock"
)

func TestUserUseCase_FindProfile(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := userMock.NewMockUserRepository(ctrl)

	cfg, _ := config.LoadConfig()

	user := usecase.NewUserUseCase(cfg, mockUserRepo)

	req := httptest.NewRequest(echo.GET, "/", nil)
	rec := httptest.NewRecorder()
	ctx := echo.New().NewContext(req, rec)

	const userID = "testUserID"

	userDomain := domain.User{ID: "testUserID"}

	type input struct {
		userID string
	}

	type output struct {
		wantUser domain.User
		wantErr  error
	}

	tests := map[string]struct {
		input         input
		prepareMockFn func(mr *userMock.MockUserRepository)
		want          output
	}{
		"Normal Case: FindProfile": {
			input{userID},
			func(mr *userMock.MockUserRepository) {
				mr.EXPECT().FindUserByUserID(gomock.Any(), userID).Return(userDomain, nil)
			},
			output{userDomain, nil},
		},
		"Abnormal Case: FindProfile": {
			input{userID},
			func(mr *userMock.MockUserRepository) {
				mr.EXPECT().FindUserByUserID(gomock.Any(), userID).Return(domain.User{}, errors.New("error"))
			},
			output{domain.User{}, fmt.Errorf("unable to retrieve user details: error")},
		},
	}

	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			tt.prepareMockFn(mockUserRepo)
			user, err := user.FindProfile(ctx, tt.input.userID)
			if err != nil {
				assert.Equal(t, tt.want.wantErr.Error(), err.Error())
			} else {
				assert.Equal(t, tt.want.wantUser, user)
			}
		})
	}
}

func TestUserUseCase_UpdateProfile(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := userMock.NewMockUserRepository(ctrl)

	cfg, _ := config.LoadConfig()

	user := usecase.NewUserUseCase(cfg, mockUserRepo)

	req := httptest.NewRequest(echo.GET, "/", nil)
	rec := httptest.NewRecorder()
	ctx := echo.New().NewContext(req, rec)
	userDomain := domain.User{
		ID:          "testUserID",
		Age:         30,
		GoogleImage: "https://example.com/image.jpg",
		FirstName:   "Test",
		LastName:    "User",
		UserName:    "testuser",
		Email:       "testuser@example.com",
		Phone:       "1234567890",
		Password:    "password",
		Verified:    true,
		BlockStatus: false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	type input struct {
		user domain.User
	}

	type output struct {
		wantErr error
	}

	tests := map[string]struct {
		input         input
		prepareMockFn func(mr *userMock.MockUserRepository)
		want          output
	}{
		"Normal Case: UpdateProfile": {
			input{userDomain},
			func(mr *userMock.MockUserRepository) {
				mr.EXPECT().FindUserByUserNameEmailOrPhone(gomock.Any(), userDomain).Return(userDomain, nil)
				mr.EXPECT().UpdateUser(gomock.Any(), gomock.Any()).Return(nil)
			},
			output{nil},
		},
		"Abnormal Case: UpdateProfile": {
			input{userDomain},
			func(mr *userMock.MockUserRepository) {
				mr.EXPECT().FindUserByUserNameEmailOrPhone(gomock.Any(), userDomain).Return(domain.User{}, errors.New("error"))
			},
			output{fmt.Errorf("unable to find user: error")},
		},
		"Abnormal Case: UpdateProfile1": {
			input{userDomain},
			func(mr *userMock.MockUserRepository) {
				mr.EXPECT().FindUserByUserNameEmailOrPhone(gomock.Any(), userDomain).Return(userDomain, nil)
				mr.EXPECT().UpdateUser(gomock.Any(), gomock.Any()).Return(errors.New("error"))
			},
			output{fmt.Errorf("unable to update user: error")},
		},
	}

	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			tt.prepareMockFn(mockUserRepo)
			err := user.UpdateProfile(ctx, tt.input.user)
			if err != nil {
				assert.Equal(t, tt.want.wantErr.Error(), err.Error())
			} else {
				assert.Equal(t, tt.want.wantErr, err)
			}
		})
	}
}

func TestUserUseCase_SaveAddress(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := userMock.NewMockUserRepository(ctrl)

	cfg, _ := config.LoadConfig()

	user := usecase.NewUserUseCase(cfg, mockUserRepo)

	req := httptest.NewRequest(echo.GET, "/", nil)
	rec := httptest.NewRecorder()
	ctx := echo.New().NewContext(req, rec)

	const userID = "testUserID"

	address := domain.Address{ID: 1, Name: "Test Street"}

	type input struct {
		userID    string
		address   domain.Address
		isDefault bool
	}

	type output struct {
		wantErr error
	}

	tests := map[string]struct {
		input         input
		prepareMockFn func(mr *userMock.MockUserRepository)
		want          output
	}{
		"Normal Case: SaveAddress": {
			input{userID, address, true},
			func(mr *userMock.MockUserRepository) {
				mr.EXPECT().IsAddressAlreadyExistForUser(gomock.Any(), address, userID).Return(false, nil)
				mr.EXPECT().SaveAddress(gomock.Any(), gomock.Any()).Return(uint(1), nil)
				mr.EXPECT().SaveUserAddress(gomock.Any(), gomock.Any()).Return(nil)
			},
			output{nil},
		},
		"Abnormal Case: SaveAddress": {
			input{userID, address, true},
			func(mr *userMock.MockUserRepository) {
				mr.EXPECT().IsAddressAlreadyExistForUser(gomock.Any(), gomock.Any(), userID).Return(true, nil)
			},
			output{fmt.Errorf("the provided address already exists for this user")},
		},
		"Abnormal Case: SaveAddress1": {
			input{userID, address, true},
			func(mr *userMock.MockUserRepository) {
				mr.EXPECT().IsAddressAlreadyExistForUser(gomock.Any(), address, userID).Return(false, errors.New("error"))
			},
			output{fmt.Errorf("unable to check if address already exists \nerror:error")},
		},
		"Abnormal Case: SaveAddress2": {
			input{userID, address, true},
			func(mr *userMock.MockUserRepository) {
				mr.EXPECT().IsAddressAlreadyExistForUser(gomock.Any(), address, userID).Return(false, nil)
				mr.EXPECT().SaveAddress(gomock.Any(), gomock.Any()).Return(uint(1), errors.New("error"))
			},
			output{fmt.Errorf("unable to save address: error")},
		},
		"Abnormal Case: SaveAddress3": {
			input{userID, address, true},
			func(mr *userMock.MockUserRepository) {
				mr.EXPECT().IsAddressAlreadyExistForUser(gomock.Any(), address, userID).Return(false, nil)
				mr.EXPECT().SaveAddress(gomock.Any(), gomock.Any()).Return(uint(1), nil)
				mr.EXPECT().SaveUserAddress(gomock.Any(), gomock.Any()).Return(errors.New("error"))
			},
			output{fmt.Errorf("unable to save user address: error")},
		},
	}

	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			tt.prepareMockFn(mockUserRepo)
			err := user.SaveAddress(ctx, tt.input.userID, tt.input.address, tt.input.isDefault)
			if err != nil {
				assert.Equal(t, tt.want.wantErr.Error(), err.Error())
			} else {
				assert.Equal(t, tt.want.wantErr, err)
			}
		})
	}
}

func TestUserUseCase_UpdateAddress(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := userMock.NewMockUserRepository(ctrl)

	cfg, _ := config.LoadConfig()

	user := usecase.NewUserUseCase(cfg, mockUserRepo)

	req := httptest.NewRequest(echo.GET, "/", nil)
	rec := httptest.NewRecorder()
	ctx := echo.New().NewContext(req, rec)

	const userID = "testUserID"

	addressBody := request.EditAddress{ID: 1, Name: "Change Street", IsDefault: new(bool)}

	type input struct {
		addressBody request.EditAddress
		userID      string
	}

	type output struct {
		wantErr error
	}

	tests := map[string]struct {
		input         input
		prepareMockFn func(mr *userMock.MockUserRepository)
		want          output
	}{
		"Normal Case: UpdateAddress": {
			input{addressBody, userID},
			func(mr *userMock.MockUserRepository) {
				mr.EXPECT().IsAddressIDExist(gomock.Any(), gomock.Any()).Return(true, nil)
				mr.EXPECT().UpdateAddress(gomock.Any(), gomock.Any()).Return(nil)
			},
			output{nil},
		},
		"Abnormal Case: UpdateAddress": {
			input{addressBody, userID},
			func(mr *userMock.MockUserRepository) {
				mr.EXPECT().IsAddressIDExist(gomock.Any(), gomock.Any()).Return(false, nil)
			},
			output{errors.New("invalid address id")},
		},
		"Normal Case: UpdateAddress with Default Address": {
			input{addressBody, userID},
			func(mr *userMock.MockUserRepository) {
				*addressBody.IsDefault = true
				mr.EXPECT().IsAddressIDExist(gomock.Any(), gomock.Any()).Return(true, nil)
				mr.EXPECT().UpdateAddress(gomock.Any(), gomock.Any()).Return(nil)
				mr.EXPECT().UpdateUserAddress(gomock.Any(), gomock.Any()).Return(nil).AnyTimes()
			},
			output{nil},
		},
		"Abnormal Case: UpdateAddress2": {
			input{addressBody, userID},
			func(mr *userMock.MockUserRepository) {
				mr.EXPECT().IsAddressIDExist(gomock.Any(), gomock.Any()).Return(false, errors.New("error"))
			},
			output{fmt.Errorf("unable to check address ID existence: error")},
		},
		"Abnormal Case: UpdateAddress1": {
			input{addressBody, userID},
			func(mr *userMock.MockUserRepository) {
				mr.EXPECT().IsAddressIDExist(gomock.Any(), gomock.Any()).Return(true, nil)
				mr.EXPECT().UpdateAddress(gomock.Any(), gomock.Any()).Return(errors.New("error"))
			},
			output{fmt.Errorf("unable to update address: error")},
		},
	}

	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			tt.prepareMockFn(mockUserRepo)
			err := user.UpdateAddress(ctx, tt.input.addressBody, tt.input.userID)
			if err != nil {
				assert.Equal(t, tt.want.wantErr.Error(), err.Error())
			} else {
				assert.Equal(t, tt.want.wantErr, err)
			}
		})
	}
}

func TestUserUseCase_FindAddresses(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := userMock.NewMockUserRepository(ctrl)

	cfg, _ := config.LoadConfig()

	user := usecase.NewUserUseCase(cfg, mockUserRepo)

	req := httptest.NewRequest(echo.GET, "/", nil)
	rec := httptest.NewRecorder()
	ctx := echo.New().NewContext(req, rec)

	const userID = "testUserID"

	addresses := []response.Address{{ID: 1, Name: "Test Street"}}

	type input struct {
		userID string
	}

	type output struct {
		wantAddresses []response.Address
		wantErr       error
	}

	tests := map[string]struct {
		input         input
		prepareMockFn func(mr *userMock.MockUserRepository)
		want          output
	}{
		"Normal Case: FindAddresses": {
			input{userID},
			func(mr *userMock.MockUserRepository) {
				mr.EXPECT().FindAllAddressByUserID(gomock.Any(), gomock.Any()).Return(addresses, nil)
			},
			output{addresses, nil},
		},
		"Abnormal Case: FindAddresses": {
			input{userID},
			func(mr *userMock.MockUserRepository) {
				mr.EXPECT().FindAllAddressByUserID(gomock.Any(), gomock.Any()).Return(nil, errors.New("error"))
			},
			output{nil, fmt.Errorf("unable to find addresses: error")},
		},
	}

	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			tt.prepareMockFn(mockUserRepo)
			addresses, err := user.FindAddresses(ctx, tt.input.userID)
			if err != nil {
				assert.Equal(t, tt.want.wantErr.Error(), err.Error())
			} else {
				assert.Equal(t, tt.want.wantAddresses, addresses)
			}
		})
	}
}

func TestUserUseCase_FindAddress(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := userMock.NewMockUserRepository(ctrl)

	cfg, _ := config.LoadConfig()

	user := usecase.NewUserUseCase(cfg, mockUserRepo)

	req := httptest.NewRequest(echo.GET, "/", nil)
	rec := httptest.NewRecorder()
	ctx := echo.New().NewContext(req, rec)
	userID := "testUserID"
	addressID := uint(1)
	address := domain.Address{ID: addressID, Name: "Test Street"}

	type input struct {
		userID    string
		addressID uint
	}

	type output struct {
		wantAddress domain.Address
		wantErr     error
	}

	tests := map[string]struct {
		input         input
		prepareMockFn func(mr *userMock.MockUserRepository)
		want          output
	}{
		"Normal Case: FindAddress": {
			input{userID, addressID},
			func(mr *userMock.MockUserRepository) {
				mr.EXPECT().FindAddressByUserIDAndAddressID(gomock.Any(), gomock.Any(), gomock.Any()).Return(address, nil)
			},
			output{address, nil},
		},
		"Abnormal Case: FindAddress": {
			input{userID, addressID},
			func(mr *userMock.MockUserRepository) {
				mr.EXPECT().FindAddressByUserIDAndAddressID(gomock.Any(), gomock.Any(), gomock.Any()).Return(domain.Address{}, errors.New("error"))
			},
			output{domain.Address{}, fmt.Errorf("unable to find user details: error")},
		},
	}

	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			tt.prepareMockFn(mockUserRepo)
			address, err := user.FindAddress(ctx, tt.input.userID, tt.input.addressID)
			if err != nil {
				assert.Equal(t, tt.want.wantErr.Error(), err.Error())
			} else {
				assert.Equal(t, tt.want.wantAddress, address)
			}
		})
	}
}

func TestUserUseCase_FindPaymentMethods(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := userMock.NewMockUserRepository(ctrl)

	cfg, _ := config.LoadConfig()

	user := usecase.NewUserUseCase(cfg, mockUserRepo)

	req := httptest.NewRequest(echo.GET, "/", nil)
	rec := httptest.NewRecorder()
	ctx := echo.New().NewContext(req, rec)
	userID := "7cbad53a-97f0-407f-84f5-40f2a1fcc053"
	paymentMethods := []response.PaymentMethod{{Number: "be4fb7c3dbc0c814f554baa2621e8f525d87b0dc50e3a039a60667876cd793c8c18077336cd37081"}}

	type input struct {
		userID string
	}

	type output struct {
		wantPaymentMethods []response.PaymentMethod
		wantErr            error
	}

	tests := map[string]struct {
		input         input
		prepareMockFn func(mr *userMock.MockUserRepository)
		want          output
	}{
		"Normal Case: FindPaymentMethods": {
			input{userID},
			func(mr *userMock.MockUserRepository) {
				mr.EXPECT().FindAllPaymentMethodsByUserID(gomock.Any(), userID).Return(paymentMethods, nil)
			},
			output{paymentMethods, nil},
		},
		"Abnormal Case: FindPaymentMethods": {
			input{userID},
			func(mr *userMock.MockUserRepository) {
				mr.EXPECT().FindAllPaymentMethodsByUserID(gomock.Any(), userID).Return(nil, errors.New("error"))
			},
			output{nil, fmt.Errorf("unable to find payment method: error")},
		},
	}

	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			tt.prepareMockFn(mockUserRepo)
			paymentMethods, err := user.FindPaymentMethods(ctx, tt.input.userID)
			if err != nil {
				assert.Equal(t, tt.want.wantErr.Error(), err.Error())
			} else {
				assert.Equal(t, tt.want.wantPaymentMethods, paymentMethods)
			}
		})
	}
}

func TestUserUseCase_SavePaymentMethod(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := userMock.NewMockUserRepository(ctrl)

	cfg, _ := config.LoadConfig()

	user := usecase.NewUserUseCase(cfg, mockUserRepo)

	req := httptest.NewRequest(echo.GET, "/", nil)
	rec := httptest.NewRecorder()
	ctx := echo.New().NewContext(req, rec)
	userID := "testUserID"
	paymentMethod := request.PaymentMethod{Number: "1234567890123456", Expiry: "12/24", Cvc: "123"}

	type input struct {
		userID        string
		paymentMethod request.PaymentMethod
	}

	type output struct {
		wantErr error
	}

	tests := map[string]struct {
		input         input
		prepareMockFn func(mr *userMock.MockUserRepository)
		want          output
	}{
		"Normal Case: SavePaymentMethod": {
			input{userID, paymentMethod},
			func(mr *userMock.MockUserRepository) {
				mr.EXPECT().SavePaymentMethod(gomock.Any(), gomock.Any()).Return(uint(1), nil)
			},
			output{nil},
		},
		"Abnormal Case: SavePaymentMethod": {
			input{userID, paymentMethod},
			func(mr *userMock.MockUserRepository) {
				mr.EXPECT().SavePaymentMethod(gomock.Any(), gomock.Any()).Return(uint(0), errors.New("error"))
			},
			output{fmt.Errorf("unable to save product: error")},
		},
	}

	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			tt.prepareMockFn(mockUserRepo)
			err := user.SavePaymentMethod(ctx, tt.input.userID, tt.input.paymentMethod)
			if err != nil {
				assert.Equal(t, tt.want.wantErr.Error(), err.Error())
			} else {
				assert.Equal(t, tt.want.wantErr, err)
			}
		})
	}
}
