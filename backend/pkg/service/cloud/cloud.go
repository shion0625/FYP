package cloud

import (
	"mime/multipart"

	"github.com/labstack/echo/v4"
)

type CloudService interface {
	SaveFile(ctx echo.Context, fileHeader *multipart.FileHeader) (uploadId string, err error)
	GetFileUrl(ctx echo.Context, uploadID string) (url string, err error)
}
