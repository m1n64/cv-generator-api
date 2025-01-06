package services

import (
	"bytes"
	"context"
	"fmt"
	"github.com/chai2010/webp"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"information-service/pkg/utils"
	"os"
)

type FileService struct {
	minioClient *utils.MinioClient
	logger      *zap.Logger
}

func NewFileService(minioClient *utils.MinioClient, logger *zap.Logger) *FileService {
	return &FileService{
		minioClient: minioClient,
		logger:      logger,
	}
}

func (fs *FileService) SaveFile(ctx context.Context, cvID uuid.UUID, file []byte) (*string, error) {
	img, _, err := image.Decode(bytes.NewReader(file))
	if err != nil {
		return nil, fmt.Errorf("failed to decode image: %v", err)
	}

	tempFile, err := os.CreateTemp("tmp/", "photo-*.webp")
	if err != nil {
		return nil, fmt.Errorf("failed to create temp file: %v", err)
	}
	defer os.Remove(tempFile.Name())

	var options webp.Options
	options.Lossless = false
	options.Quality = 80

	if err := webp.Encode(tempFile, img, &options); err != nil {
		return nil, fmt.Errorf("failed to encode image to webp: %v", err)
	}
	if err := tempFile.Close(); err != nil {
		return nil, fmt.Errorf("failed to close temp file: %v", err)
	}

	objectName, err := fs.UploadUserPhoto(ctx, cvID, tempFile.Name())
	if err != nil {
		return nil, fmt.Errorf("failed to upload photo: %v", err)
	}

	return &objectName, nil
}

func (fs *FileService) UploadUserPhoto(ctx context.Context, cvID uuid.UUID, filePath string) (string, error) {
	objectName := fmt.Sprintf("users/%s/photo-%s.webp", cvID.String(), uuid.New().String())
	err := fs.minioClient.UploadFile(ctx, objectName, filePath, "image/webp")
	if err != nil {
		return "", err
	}
	return objectName, nil
}

func (fs *FileService) GetFileURL(ctx context.Context, objectName string) (string, error) {
	return fs.minioClient.GetFileURL(ctx, objectName)
}

func (fs *FileService) GetFileAsBytes(ctx context.Context, objectName string) ([]byte, error) {
	return fs.minioClient.GetFileAsBytes(ctx, objectName)
}
