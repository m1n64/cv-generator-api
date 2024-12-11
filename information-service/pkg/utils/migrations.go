package utils

import "information-service/internal/information/models"

func StartMigrations() {
	db := GetDBConnection()
	db.AutoMigrate(
		&models.Information{},
		&models.Language{},
		&models.Skill{},
		&models.Contact{},
		&models.WorkExperience{},
		&models.Education{},
		&models.Certificate{},
	)
}
