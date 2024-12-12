package containers

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"information-service/internal/information/repositories"
	"information-service/internal/information/services"
	"information-service/pkg/utils"
)

type Dependencies struct {
	DB                   *gorm.DB
	RedisClient          *redis.Client
	RedisAdapter         *utils.RedisAdapter
	InformationRepo      repositories.InformationRepository
	CVInformationService *services.CVInformationService
}

func InitializeDependencies() (*Dependencies, error) {
	utils.InitLogs()
	utils.LoadEnv()
	utils.CreateRedisConn()
	utils.InitDBConnection()
	utils.StartMigrations()
	utils.InitValidator()

	db := utils.GetDBConnection()

	_, redisClient := utils.GetRedisConn()

	redisAdapter := utils.NewRedisAdapter(redisClient)

	// Repositories
	informationRepo := repositories.NewInformationRepository(db)

	// Services
	cvService := services.NewCVInformationService(informationRepo, db)

	// Dependencies
	return &Dependencies{
		DB:                   db,
		RedisClient:          redisClient,
		RedisAdapter:         redisAdapter,
		InformationRepo:      informationRepo,
		CVInformationService: cvService,
	}, nil
}
