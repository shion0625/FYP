package usecase

const (
	countryCode       = "+91"
	otpExpireDuration = time.Minute * 2
)

type authUseCase struct {

	userRepo     interfaces.UserRepository
	// adminRepo    interfaces.AdminRepository
	// tokenService token.TokenService
	// authRepo interfaces.AuthRepository
	// optAuth      otp.OtpAuth
}

func NewAuthUseCase(
	userRepo interfaces.UserRepository,
	// adminRepo interfaces.AdminRepository,
	// tokenService token.TokenService,
	// authRepo interfaces.AuthRepository,
	// optAuth otp.OtpAut
	) service.AuthUseCase {

	return &authUseCase{
		userRepo:     userRepo,
		// adminRepo:    adminRepo,
		// tokenService: tokenService,
		// authRepo:     authRepo,
		// optAuth:      optAuth,
	}
}


func (c *authUseCase) UserLogin(ctx context.Context, loginInfo request.Login) (uint, error) {

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
