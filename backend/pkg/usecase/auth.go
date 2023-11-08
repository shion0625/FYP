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

func (c *authUseCase) UserSignUp(ctx echo.Context, signUpDetails domain.User) (string, error) {
	existUser, err := c.userRepo.FindUserByUserNameEmailOrPhoneNotID(ctx, signUpDetails)
	if err != nil {
		return "", utils.PrependMessageToError(err, "failed to check user details already exist")
	}

	// if user credentials already exist and  verified then return it as errors
	if existUser.ID != 0 && existUser.Verified {
		err = utils.CompareUserExistingDetails(existUser, signUpDetails)
		err = utils.AppendMessageToError(ErrUserAlreadyExit, err.Error())
		return "", err
	}

	userID := existUser.ID

	if userID == 0 { // if user not exist then save user on database
		hashPass, err := utils.GenerateHashFromPassword(signUpDetails.Password)
		if err != nil {
			return "", utils.PrependMessageToError(err, "failed to hash the password")
		}

		signUpDetails.Password = string(hashPass)
		userID, err = c.userRepo.SaveUser(ctx, signUpDetails)

		if err != nil {
			return "", utils.PrependMessageToError(err, "failed to save user details")
		}
	}

	return "", nil
}
