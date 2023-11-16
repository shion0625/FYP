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

func (u *UserHandler) GetProfile(ctx echo.Context) {
	userID, err := utils.GetUserIdFromContext(ctx)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to retrieve user details", err, nil)

		return
	}

	user, err := u.userUseCase.FindProfile(ctx, userID)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to retrieve user details", err, nil)

		return
	}

	response.SuccessResponse(ctx, http.StatusOK, "Successfully retrieved user details", user)
}

func (u *UserHandler) UpdateProfile(ctx echo.Context) {
	userID, err := utils.GetUserIdFromContext(ctx)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to retrieve user details", err, nil)

		return
	}

	var body request.EditUser

	if err := ctx.Bind(&body); err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, BindJsonFailMessage, err, nil)

		return
	}

	var user domain.User
	if err := copier.Copy(&user, &body); err != nil {
		response.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to copy user data", err, nil)

		return
	}

	user.ID = userID

	if err := u.userUseCase.UpdateProfile(ctx, user); err != nil {
		response.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to update profile", err, nil)

		return
	}

	response.SuccessResponse(ctx, http.StatusOK, "Successfully profile updated", nil)
}

func (u *UserHandler) SaveAddress(ctx echo.Context) {
	var body request.Address
	if err := ctx.Bind(&body); err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, BindJsonFailMessage, err, nil)

		return
	}

	userID, err := utils.GetUserIdFromContext(ctx)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to retrieve user details", err, nil)

		return
	}

	var address domain.Address

	if err := copier.Copy(&address, &body); err != nil {
		response.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to copy address data", err, nil)

		return
	}

	// check is default is null
	if body.IsDefault == nil {
		body.IsDefault = new(bool)
	}

	if err := u.userUseCase.SaveAddress(ctx, userID, address, *body.IsDefault); err != nil {
		response.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to save address", err, nil)

		return
	}

	response.SuccessResponse(ctx, http.StatusCreated, "Successfully address saved")
}

func (u *UserHandler) GetAllAddresses(ctx echo.Context) {
	userID, err := utils.GetUserIdFromContext(ctx)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to retrieve user details", err, nil)

		return
	}

	addresses, err := u.userUseCase.FindAddresses(ctx, userID)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to get user addresses", err, nil)

		return
	}

	if addresses == nil {
		response.SuccessResponse(ctx, http.StatusOK, "No addresses found")

		return
	}

	response.SuccessResponse(ctx, http.StatusOK, "Successfully retrieved all user addresses", addresses)
}

func (u *UserHandler) UpdateAddress(ctx echo.Context) {
	userID, err := utils.GetUserIdFromContext(ctx)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to retrieve user details", err, nil)

		return
	}

	var body request.EditAddress

	if err := ctx.Bind(&body); err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, BindJsonFailMessage, err, nil)

		return
	}

	// address is_default reference pointer need to change in future
	if body.IsDefault == nil {
		body.IsDefault = new(bool)
	}

	if err := u.userUseCase.UpdateAddress(ctx, body, userID); err != nil {
		response.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to update user address", err, nil)

		return
	}

	response.SuccessResponse(ctx, http.StatusOK, "successfully addresses updated", body)
}
