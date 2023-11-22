package cloud

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"time"

	"cloud.google.com/go/storage"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/shion0625/FYP/backend/pkg/config"
	"github.com/shion0625/FYP/backend/pkg/domain"
	"google.golang.org/api/option"
)

type gcpService struct {
	service         *storage.Client
	bucketName      string
	credentialsFile string
	googleAccessID  string
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
		googleAccessID:  cfg.GcpServiceAccount,
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
	// サービスアカウントキーファイルを読み込む
	jsonKey, err := os.ReadFile(c.credentialsFile)
	if err != nil {
		return "", fmt.Errorf("failed to read service account key file: %w", err)
	}

	// JSONを構造体にデコードする
	var sa domain.ServiceAccount
	if err := json.Unmarshal(jsonKey, &sa); err != nil {
		return "", fmt.Errorf("failed to decode service account key file: %w", err)
	}

	url, err := storage.SignedURL(c.bucketName, uploadID, &storage.SignedURLOptions{
		GoogleAccessID: c.googleAccessID,
		PrivateKey:     []byte(sa.PrivateKey),
		Method:         http.MethodGet,
		Expires:        time.Now().Add(filePreSignExpireDuration),
	})

	if err != nil {
		return "", fmt.Errorf("failed to sign url: %w", err)
	}

	return url, nil
}
