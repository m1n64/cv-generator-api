package services

import (
	"bytes"
	"context"
	"cv-generator-service/internal/generator/entities"
	"cv-generator-service/internal/generator/enums"
	"cv-generator-service/internal/notifications/services"
	"cv-generator-service/pkg/utils"
	"encoding/base64"
	"fmt"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"github.com/gomarkdown/markdown"
	"html/template"
	"net/url"
	"os"
	"time"
)

type GeneratePdfService struct {
	pdfGeneratorService *PdfGeneratorService
	notificationService *services.NotificationService
	minio               *utils.MinioClient
	chromeAllocator     *utils.ChromeAllocator
}

func NewGeneratePdfService(service *PdfGeneratorService, notificationService *services.NotificationService, minio *utils.MinioClient, chromeAllocator *utils.ChromeAllocator) *GeneratePdfService {
	return &GeneratePdfService{
		pdfGeneratorService: service,
		notificationService: notificationService,
		minio:               minio,
		chromeAllocator:     chromeAllocator,
	}
}

func (s *GeneratePdfService) GeneratePDF(cvInfo entities.CvInfo) error {
	cvName := fmt.Sprintf("%s - %s (%s)", cvInfo.CV.Title, time.Now().Format(time.DateOnly), *cvInfo.Color)
	generated, _ := s.pdfGeneratorService.CreateGeneratedPDF(cvInfo.CvID, cvInfo.UserID, cvName, nil, enums.StatusPending)

	ctx, cancel := chromedp.NewContext(s.chromeAllocator.AllocatorCtx)
	defer cancel()

	var buf bytes.Buffer
	t := template.Must(template.New("resume").Funcs(template.FuncMap{
		"toBase64": func(data *[]byte) string {
			if data == nil {
				return ""
			}

			return base64.StdEncoding.EncodeToString(*data)
		},
		"formatDate": func(dateStr string) string {
			layout := "2006-01-02 15:04:05 -0700 MST"
			t, err := time.Parse(layout, dateStr)
			if err != nil {
				return dateStr
			}
			return t.Format("January 2, 2006")
		},
		"renderMd": func(data string) template.HTML {
			htmlContent := markdown.ToHTML([]byte(data), nil, nil)

			return template.HTML(htmlContent)
		},
	}).Parse(cvInfo.Template))
	err := t.Execute(&buf, map[string]interface{}{
		"AccentColor":     cvInfo.Color,
		"Information":     cvInfo.Information,
		"Contacts":        cvInfo.Contacts,
		"Skills":          cvInfo.Skills,
		"Languages":       cvInfo.Languages,
		"WorkExperiences": cvInfo.WorkExperiences,
		"Educations":      cvInfo.Educations,
		"Certificates":    cvInfo.Certificates,
	})
	if err != nil {
		return err
	}

	escapedHTML := "data:text/html," + url.PathEscape(buf.String())

	os.MkdirAll("tmp/", os.ModePerm)

	_ = s.pdfGeneratorService.UpdateStatus(generated.ID, enums.StatusInProgress)

	var pageHeight float64

	err = chromedp.Run(ctx, chromedp.Tasks{
		chromedp.Navigate(escapedHTML),
		chromedp.WaitReady("body"),
		chromedp.Evaluate(`document.body.scrollHeight / 96`, &pageHeight),
		chromedp.ActionFunc(func(ctx context.Context) error {
			buf, _, err := page.PrintToPDF().
				WithPrintBackground(true).
				WithPreferCSSPageSize(false).
				WithPaperWidth(8.5).
				WithPaperHeight(pageHeight).
				WithMarginTop(0).
				WithMarginBottom(0).
				WithMarginLeft(0).
				WithMarginRight(0).
				Do(ctx)

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

			s.notificationService.SendSuccess(cvInfo.UserID, cvInfo.CvID, generated.ID, "")

			return nil
		}),
	})
	if err != nil {
		err := s.pdfGeneratorService.UpdateStatus(generated.ID, enums.StatusFailed)
		if err != nil {
			return err
		}

		s.notificationService.SendError(cvInfo.UserID, cvInfo.CvID, generated.ID, err)

		return fmt.Errorf("chromedp execution failed: %w", err)
	}

	return nil
}
