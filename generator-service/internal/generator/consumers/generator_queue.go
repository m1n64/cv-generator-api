package consumers

import (
	"cv-generator-service/internal/generator/models"
	"cv-generator-service/internal/generator/services"
	"github.com/goccy/go-json"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
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
		msg.Nack(false, false)
		return
	}

	h.logger.Info("starting to process message", zap.String("cv_id", cvInfo.CvID.String()))

	go func() {
		defer func() {
			if r := recover(); r != nil {
				h.logger.Error("panic in goroutine", zap.Any("recover", r))
			}
			<-h.semaphore
		}()

		h.semaphore <- struct{}{}
		defer func() { <-h.semaphore }()

		if err := h.generatePdfService.GeneratePDF(cvInfo); err != nil {
			h.logger.Error("error generating PDF", zap.Error(err))
			msg.Nack(false, true)
		}

		h.logger.Info("successfully processed message", zap.String("cv_id", cvInfo.CvID.String()))
		msg.Ack(false)
	}()
}
