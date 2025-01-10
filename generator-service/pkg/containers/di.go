package containers

import (
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"information-service/internal/generator/consumers"
	"information-service/internal/generator/repositories"
	"information-service/internal/generator/services"
	"information-service/pkg/utils"
	"log"
	"os"
)

const (
	maxPdfQueueWorkers = 10
)

type Dependencies struct {
	DB               *gorm.DB
	RedisClient      *redis.Client
	RedisAdapter     *utils.RedisAdapter
	Logger           *zap.Logger
	MinioClient      *utils.MinioClient
	RabbitMQ         *utils.RabbitMQConnection
	PdfGeneratorRepo repositories.PdfGeneratorRepository

	PdfGeneratorService *services.PdfGeneratorService
	GeneratePdfService  *services.GeneratePdfService
}

func InitializeDependencies() (*Dependencies, error) {
	utils.InitLogs()
	utils.LoadEnv()
	utils.CreateRedisConn()
	utils.InitDBConnection()
	utils.StartMigrations()
	utils.InitValidator()
	rabbitMQ := utils.ConnectRabbitMQ()
	utils.InitializeQueues()

	logger := utils.GetLogger()

	db := utils.GetDBConnection()

	_, redisClient := utils.GetRedisConn()

	redisAdapter := utils.NewRedisAdapter(redisClient)

	minioClient := utils.NewMinioClient(os.Getenv("MINIO_ENDPOINT"), os.Getenv("MINIO_ROOT_USER"), os.Getenv("MINIO_ROOT_PASSWORD"), "cv-pdf", os.Getenv("MINIO_SECURE") == "true")

	// Repositories
	pdfGeneratorRepo := repositories.NewPdfGeneratorGormRepository(db)

	// Services
	pdfGeneratorService := services.NewPdfGeneratorService(pdfGeneratorRepo, db)
	generatorPdfService := services.NewGeneratePdfService(pdfGeneratorService)

	// Dependencies
	return &Dependencies{
		DB:                  db,
		RedisClient:         redisClient,
		RedisAdapter:        redisAdapter,
		Logger:              logger,
		MinioClient:         minioClient,
		RabbitMQ:            rabbitMQ,
		PdfGeneratorRepo:    pdfGeneratorRepo,
		PdfGeneratorService: pdfGeneratorService,
		GeneratePdfService:  generatorPdfService,
	}, nil
}

func InitializeQueuesConsumer(dependencies *Dependencies) {
	generatorConsumer := consumers.NewGeneratorPdfConsumer(dependencies.GeneratePdfService, dependencies.Logger, maxPdfQueueWorkers)

	err := utils.ListenToQueue(utils.PDFGenerateQueue, generatorConsumer.HandleGenerateCvToPdf)
	if err != nil {
		log.Fatalf("Error starting listener for PDFGenerateQueue: %v", err)
	}
}
