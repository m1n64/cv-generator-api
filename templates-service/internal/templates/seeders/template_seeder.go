package seeders

import (
	"context"
	"cv-templates-service/internal/templates/services"
	"cv-templates-service/pkg/utils"
	"fmt"
	"os"
)

type templateSeeder struct {
	defaultTemplateService *services.DefaultTemplateService
	minio                  *utils.MinioClient
}

func NewTemplateSeeder(defaultTemplateService *services.DefaultTemplateService, minio *utils.MinioClient) utils.Seeder {
	return &templateSeeder{
		defaultTemplateService: defaultTemplateService,
		minio:                  minio,
	}
}

func (s *templateSeeder) Seed() error {
	filePath := "./templates/main.html"

	_, err := os.Stat(filePath)
	if err != nil {
		return fmt.Errorf("file does not exist or cannot be accessed: %w", err)
	}

	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	objectName := fmt.Sprintf("template-%s.html", utils.RandStringBytesRmndr(16))
	err = s.minio.UploadFile(context.Background(), objectName, filePath, "text/html")
	if err != nil {
		return fmt.Errorf("failed to upload file to MinIO: %w", err)
	}

	_, err = s.defaultTemplateService.CreateDefaultTemplate(objectName)
	if err != nil {
		return err
	}

	return nil
}
