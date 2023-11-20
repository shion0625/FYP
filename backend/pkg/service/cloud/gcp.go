package cloud

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"time"

	"cloud.google.com/go/storage"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/shion0625/FYP/backend/pkg/config"
	"google.golang.org/api/option"
)

type gcpService struct {
	service         *storage.Client
	bucketName      string
	credentialsFile string
}

const (
	filePreSignExpireDuration = time.Hour * 12
)

func NewGCPCloudService(cfg *config.Config) (CloudService, error) {
	ctx := context.Background()

	client, err := storage.NewClient(ctx, option.WithCredentialsFile(cfg.CredentialsFile))
	if err != nil {
		return nil, fmt.Errorf("failed to create client for gcp service : %w", err)
	}

	return &gcpService{
		service:         client,
		bucketName:      cfg.GcpBucketName,
		credentialsFile: cfg.CredentialsFile,
	}, nil
}

func (c *gcpService) SaveFile(ctx echo.Context, fileHeader *multipart.FileHeader) (string, error) {
	file, err := fileHeader.Open()
	if err != nil {
		return "", fmt.Errorf("failed to open file: %w", err)
	}

	uploadID := uuid.New().String()

	wc := c.service.Bucket(c.bucketName).Object(uploadID).NewWriter(ctx.Request().Context())
	if _, err = io.Copy(wc, file); err != nil {
		return "", fmt.Errorf("failed to upload file: %w", err)
	}

	if err := wc.Close(); err != nil {
		return "", fmt.Errorf("failed to close writer: %w", err)
	}

	return uploadID, nil
}

func (c *gcpService) GetFileUrl(ctx echo.Context, uploadID string) (string, error) {
	attrs, err := c.service.Bucket(c.bucketName).Object(uploadID).Attrs(ctx.Request().Context())
	if err != nil {
		return "", fmt.Errorf("failed to get attributes of uploaded file: %w", err)
	}

	return attrs.MediaLink, nil
}
