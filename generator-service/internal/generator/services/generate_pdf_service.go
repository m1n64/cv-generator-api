package services

import (
	"fmt"
	"information-service/internal/generator/models"
)

type GeneratePdfService struct {
	pdfGeneratorService *PdfGeneratorService
}

func NewGeneratePdfService(service *PdfGeneratorService) *GeneratePdfService {
	return &GeneratePdfService{
		pdfGeneratorService: service,
	}
}

func (s *GeneratePdfService) GeneratePDF(cvInfo models.CvInfo) error {
	fmt.Println(cvInfo)

	return nil
}
