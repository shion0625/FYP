package utils

import (
	"github.com/labstack/echo/v4"
)

// To append the message to the error.
func GetParamID(ctx echo.Context, key string) (*uint, error) {
	IDStr := ctx.QueryParam(key)

	if IDStr == "" {
		return nil, nil
	}

	ID, err := ParseStringToUint32(IDStr)
	if err != nil {
		return nil, err
	}

	return &ID, nil
}
