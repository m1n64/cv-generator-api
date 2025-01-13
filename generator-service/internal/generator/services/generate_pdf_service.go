package services

import (
	"context"
	"cv-generator-service/internal/generator/enums"
	"cv-generator-service/internal/generator/models"
	"cv-generator-service/pkg/utils"
	"fmt"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"os"
	"time"
)

type GeneratePdfService struct {
	pdfGeneratorService *PdfGeneratorService
	minio               *utils.MinioClient
}

func NewGeneratePdfService(service *PdfGeneratorService, minio *utils.MinioClient) *GeneratePdfService {
	return &GeneratePdfService{
		pdfGeneratorService: service,
		minio:               minio,
	}
}

func (s *GeneratePdfService) GeneratePDF(cvInfo models.CvInfo) error {
	cvName := fmt.Sprintf("%s - %s", cvInfo.CV.Title, time.Now().Format(time.DateOnly))
	generated, _ := s.pdfGeneratorService.CreateGeneratedPDF(cvInfo.CvID, cvInfo.UserID, cvName, nil, enums.StatusPending)

	ctx, cancel := chromedp.NewExecAllocator(context.Background(),
		append(chromedp.DefaultExecAllocatorOptions[:],
			chromedp.Flag("headless", true),
			chromedp.Flag("disable-gpu", true),
			chromedp.Flag("disable-extensions", true),
			chromedp.Flag("no-sandbox", true),
			chromedp.Flag("disable-setuid-sandbox", true),
		)...,
	)
	defer cancel()

	ctx, cancel = chromedp.NewContext(ctx)
	defer cancel()

	html := `<html>
	<body>
	<div>text</div>
	<img src="https://pkg.go.dev/static/shared/gopher/package-search-700x300.jpeg"/>
	<img src="https://go.dev/images/gophers/motorcycle.svg"/>
	<img src="https://go.dev/images/go_google_case_study_carousel.png" />
	</body>
	</html>`

	os.MkdirAll("tmp/", os.ModePerm)

	_ = s.pdfGeneratorService.UpdateStatus(generated.ID, enums.StatusInProgress)

	err := chromedp.Run(ctx, chromedp.Tasks{
		chromedp.Navigate("data:text/html," + html),
		chromedp.WaitReady("body"),
		chromedp.ActionFunc(func(ctx context.Context) error {
			buf, _, err := page.PrintToPDF().WithPrintBackground(false).Do(ctx)
			if err != nil {
				return fmt.Errorf("failed to generate PDF: %w", err)
			}

			tempFile, err := os.CreateTemp("tmp/", "cv-*.pdf")
			if err != nil {
				return fmt.Errorf("failed to create temp file: %w", err)
			}
			defer os.Remove(tempFile.Name())

			_, err = tempFile.Write(buf)
			if err != nil {
				return fmt.Errorf("failed to write to temp file: %w", err)
			}
			tempFile.Close()

			objectName := fmt.Sprintf("pdfs/%s/%s/%s.pdf", cvInfo.UserID.String(), cvInfo.CvID.String(), utils.RandStringBytesRmndr(16))

			err = s.minio.UploadFile(ctx, objectName, tempFile.Name(), "application/pdf")
			if err != nil {
				return fmt.Errorf("failed to upload file to MinIO: %w", err)
			}

			_, err = s.pdfGeneratorService.UpdateGeneratedPDF(generated.ID, generated.CvID, generated.UserID, generated.Title, &objectName, enums.StatusCompleted)
			if err != nil {
				return err
			}

			return nil
		}),
	})
	if err != nil {
		err := s.pdfGeneratorService.UpdateStatus(generated.ID, enums.StatusFailed)
		if err != nil {
			return err
		}
		return fmt.Errorf("chromedp execution failed: %w", err)
	}

	return nil
}
