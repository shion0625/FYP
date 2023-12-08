package handler

import (
	"net/http"

	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
	"github.com/shion0625/FYP/backend/pkg/api/handler/interfaces"
	"github.com/shion0625/FYP/backend/pkg/api/handler/request"
	"github.com/shion0625/FYP/backend/pkg/api/handler/response"
	"github.com/shion0625/FYP/backend/pkg/domain"
	usecaseInterface "github.com/shion0625/FYP/backend/pkg/usecase/interfaces"
	"github.com/shion0625/FYP/backend/pkg/utils"
)

type UserHandler struct {
	userUseCase usecaseInterface.UserUseCase
}

func NewUserHandler(userUsecase usecaseInterface.UserUseCase) interfaces.UserHandler {
	return &UserHandler{
		userUseCase: userUsecase,
	}
}

func (u *UserHandler) GetProfile(ctx echo.Context) error {
	userID, err := utils.GetUserIdFromContext(ctx)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to retrieve user details", err, nil)

		return nil
	}

	user, err := u.userUseCase.FindProfile(ctx, userID)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to retrieve user details", err, nil)

		return nil
	}

	responseUser := response.User{
		ID:          user.ID,
		GoogleImage: user.GoogleImage,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Age:         user.Age,
		Email:       user.Email,
		UserName:    user.UserName,
		Phone:       user.Phone,
		BlockStatus: user.BlockStatus,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	response.SuccessResponse(ctx, http.StatusOK, "Successfully retrieved user details", responseUser)

	return nil
}

func (u *UserHandler) UpdateProfile(ctx echo.Context) error {
	userID, err := utils.GetUserIdFromContext(ctx)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to retrieve user details", err, nil)

		return nil
	}

	var body request.EditUser

	if err := ctx.Bind(&body); err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, BindJsonFailMessage, err, nil)

		return nil
	}

	if err := ctx.Validate(body); err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request data", err, nil)

		return nil
	}

	var user domain.User
	if err := copier.Copy(&user, &body); err != nil {
		response.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to copy user data", err, nil)

		return nil
	}

	user.ID = userID

	if err := u.userUseCase.UpdateProfile(ctx, user); err != nil {
		response.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to update profile", err, nil)

		return nil
	}

	response.SuccessResponse(ctx, http.StatusOK, "Successfully profile updated", nil)

	return nil
}

func (u *UserHandler) SaveAddress(ctx echo.Context) error {
	var body request.Address
	if err := ctx.Bind(&body); err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, BindJsonFailMessage, err, nil)

		return nil
	}

	if err := ctx.Validate(body); err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request data", err, nil)

		return nil
	}

	userID, err := utils.GetUserIdFromContext(ctx)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to retrieve user details", err, nil)

		return nil
	}

	var address domain.Address

	if err := copier.Copy(&address, &body); err != nil {
		response.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to copy address data", err, nil)

		return nil
	}

	// check is default is null
	if body.IsDefault == nil {
		body.IsDefault = new(bool)
	}

	if err := u.userUseCase.SaveAddress(ctx, userID, address, *body.IsDefault); err != nil {
		response.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to save address", err, nil)

		return nil
	}

	response.SuccessResponse(ctx, http.StatusCreated, "Successfully address saved", nil)

	return nil
}

func (u *UserHandler) GetAllAddresses(ctx echo.Context) error {
	userID, err := utils.GetUserIdFromContext(ctx)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to retrieve user details", err, nil)

		return nil
	}

	addresses, err := u.userUseCase.FindAddresses(ctx, userID)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to get user addresses", err, nil)

		return nil
	}

	if addresses == nil {
		response.SuccessResponse(ctx, http.StatusOK, "No addresses found", nil)

		return nil
	}

	response.SuccessResponse(ctx, http.StatusOK, "Successfully retrieved all user addresses", addresses)

	return nil
}

func (u *UserHandler) UpdateAddress(ctx echo.Context) error {
	userID, err := utils.GetUserIdFromContext(ctx)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to retrieve user details", err, nil)

		return nil
	}

	var body request.EditAddress

	if err := ctx.Bind(&body); err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, BindJsonFailMessage, err, nil)

		return nil
	}

	if err := ctx.Validate(body); err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request data", err, nil)

		return nil
	}

	// address is_default reference pointer need to change in future
	if body.IsDefault == nil {
		body.IsDefault = new(bool)
	}

	if err := u.userUseCase.UpdateAddress(ctx, body, userID); err != nil {
		response.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to update user address", err, nil)

		return nil
	}

	response.SuccessResponse(ctx, http.StatusOK, "successfully addresses updated", body)

	return nil
}
