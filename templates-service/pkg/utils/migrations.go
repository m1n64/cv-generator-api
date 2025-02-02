package utils

import (
	"cv-templates-service/internal/templates/models"
	models2 "cv-templates-service/pkg/infrastructure/models"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Seeder interface {
	Seed() error
}

func StartMigrations() {
	db := GetDBConnection()
	db.AutoMigrate(&models.Template{}, &models2.Migration{})

	addDefaultMigration(db)
}

func addDefaultMigration(db *gorm.DB) {
	var count int64
	db.Model(&models2.Migration{}).Where("name = ?", "template").Count(&count)

	if count == 0 {
		fmt.Println("Adding default migrations...")
		err := db.Create(&models2.Migration{
			Name: "template",
		}).Error
		if err != nil {
			fmt.Println("Failed to add migrations:", err)
		}
		fmt.Println("Default migrations added.")
	}
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
