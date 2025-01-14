package utils

import (
	"cv-templates-service/internal/templates/models"
	"go.uber.org/zap"
)

type Seeder interface {
	Seed() error
}

func StartMigrations() {
	db := GetDBConnection()
	db.AutoMigrate(&models.Template{})
}

func SeedDB(seeders []Seeder) {
	logger.Info("Seeding database...")

	for _, seeder := range seeders {
		if err := seeder.Seed(); err != nil {
			logger.Error("Failed to seed database", zap.Error(err))
		}
	}

	logger.Info("Database seeded successfully")
}
