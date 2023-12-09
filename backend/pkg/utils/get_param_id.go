package utils

import (
	"errors"

	"github.com/labstack/echo/v4"
)

var ErrInvalidParam = errors.New("invalid parameter")

// To append the message to the error.
func GetParamID(ctx echo.Context, key string) (*uint, error) {
	IDStr := ctx.QueryParam(key)

	if IDStr == "" {
		//nolint:nilnil
		return nil, nil
	}

	ID, err := ParseStringToUint32(IDStr)
	if err != nil {
		return nil, err
	}

	return &ID, nil
}
