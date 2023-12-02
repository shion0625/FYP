package request

import (
	"strconv"

	"github.com/labstack/echo/v4"
)

type Pagination struct {
	PageNumber uint64 `json:"pageNumber" validate:"number"`
	Count      uint64 `json:"count"      validate:"number"`
}

const (
	defaultPageNumber = 1
	defaultPageCount  = 10
)

func GetPagination(ctx echo.Context) Pagination {
	pagination := Pagination{
		PageNumber: defaultPageNumber,
		Count:      defaultPageCount,
	}

	num, err := strconv.ParseUint(ctx.QueryParam("page_number"), 10, 64)
	if err == nil {
		pagination.PageNumber = num
	}

	num, err = strconv.ParseUint(ctx.QueryParam("count"), 10, 64)
	if err == nil {
		pagination.Count = num
	}

	return pagination
}
