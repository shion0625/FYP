package usecase

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/shion0625/FYP/backend/pkg/api/handler/request"
	"github.com/shion0625/FYP/backend/pkg/domain"
	repoInterfaces "github.com/shion0625/FYP/backend/pkg/repository/interfaces"
	"github.com/shion0625/FYP/backend/pkg/usecase/interfaces"
	"github.com/shion0625/FYP/backend/pkg/utils"
)

const (
	countryCode       = "+91"
	otpExpireDuration = time.Minute * 2
)

type authUseCase struct {
	userRepo repoInterfaces.UserRepository
}

func NewAuthUseCase(
	userRepo repoInterfaces.UserRepository,
) interfaces.AuthUseCase {

	return &authUseCase{
		userRepo: userRepo,
	}
}

func (c *authUseCase) UserLogin(ctx echo.Context, loginInfo request.Login) (uint, error) {

	var (
		user domain.User
		err  error
	)
	switch {
	case loginInfo.Email != "":
		user, err = c.userRepo.FindUserByEmail(ctx, loginInfo.Email)
	case loginInfo.UserName != "":
		user, err = c.userRepo.FindUserByUserName(ctx, loginInfo.UserName)
	case loginInfo.Phone != "":
		user, err = c.userRepo.FindUserByPhoneNumber(ctx, loginInfo.Phone)
	default:
		return 0, ErrEmptyLoginCredentials
	}

	if err != nil {
		return 0, utils.PrependMessageToError(err, "failed to find user from database")
	}

	if user.ID == 0 {
		return 0, ErrUserNotExist
	}

	if !user.Verified {
		return 0, ErrUserNotVerified
	}

	if user.BlockStatus {
		return 0, ErrUserBlocked
	}

	err = utils.ComparePasswordWithHashedPassword(loginInfo.Password, user.Password)
	if err != nil {
		return 0, ErrWrongPassword
	}

	return user.ID, nil
}
