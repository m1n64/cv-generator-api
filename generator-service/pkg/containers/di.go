package containers

import (
	"cv-generator-service/internal/generator/consumers"
	"cv-generator-service/internal/generator/repositories"
	"cv-generator-service/internal/generator/services"
	services2 "cv-generator-service/internal/notifications/services"
	"cv-generator-service/pkg/utils"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"log"
	"os"
)

const (
	maxPdfQueueWorkers = 10
)

type Dependencies struct {
	DB                  *gorm.DB
	RedisClient         *redis.Client
	RedisAdapter        *utils.RedisAdapter
	Logger              *zap.Logger
	MinioClient         *utils.MinioClient
	RabbitMQ            *utils.RabbitMQConnection
	ChromeAllocator     *utils.ChromeAllocator
	PdfGeneratorRepo    repositories.PdfGeneratorRepository
	NotificationService *services2.NotificationService
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

	chromeAllocator := utils.NewChromeAllocator()
	chromeAllocator.Init()

	// Repositories
	pdfGeneratorRepo := repositories.NewPdfGeneratorGormRepository(db)

	// Services
	notificationService := services2.NewNotificationService(rabbitMQ, logger)
	pdfGeneratorService := services.NewPdfGeneratorService(pdfGeneratorRepo, db)
	generatorPdfService := services.NewGeneratePdfService(pdfGeneratorService, notificationService, minioClient, chromeAllocator)

	// Dependencies
	return &Dependencies{
		DB:                  db,
		RedisClient:         redisClient,
		RedisAdapter:        redisAdapter,
		Logger:              logger,
		MinioClient:         minioClient,
		RabbitMQ:            rabbitMQ,
		ChromeAllocator:     chromeAllocator,
		PdfGeneratorRepo:    pdfGeneratorRepo,
		NotificationService: notificationService,
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
