package consumers

import (
	"github.com/goccy/go-json"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
	"information-service/internal/generator/models"
	"information-service/internal/generator/services"
)

type GeneratorPdfConsumer struct {
	generatePdfService *services.GeneratePdfService
	logger             *zap.Logger
	semaphore          chan struct{}
}

func NewGeneratorPdfConsumer(generatePdfService *services.GeneratePdfService, logger *zap.Logger, maxWorkers int) *GeneratorPdfConsumer {
	return &GeneratorPdfConsumer{
		generatePdfService: generatePdfService,
		logger:             logger,
		semaphore:          make(chan struct{}, maxWorkers),
	}
}

func (h *GeneratorPdfConsumer) HandleGenerateCvToPdf(msg amqp.Delivery) {
	var cvInfo models.CvInfo
	if err := json.Unmarshal(msg.Body, &cvInfo); err != nil {
		h.logger.Error("error unmarshalling cv info", zap.Error(err))
		return
	}

	go func() {
		h.semaphore <- struct{}{}
		defer func() { <-h.semaphore }()

		if err := h.generatePdfService.GeneratePDF(cvInfo); err != nil {
			h.logger.Error("error generating PDF", zap.Error(err))
			msg.Nack(false, true)
			return
		}

		msg.Ack(false)
	}()
}
