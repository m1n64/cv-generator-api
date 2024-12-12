package utils

import (
	models2 "information-service/internal/certificates/models"
	models3 "information-service/internal/contacts/models"
	models4 "information-service/internal/educations/models"
	models5 "information-service/internal/experiences/models"
	"information-service/internal/information/models"
	models6 "information-service/internal/languages/models"
	models7 "information-service/internal/skills/models"
)

func StartMigrations() {
	db := GetDBConnection()
	db.AutoMigrate(
		&models.Information{},
		&models6.Language{},
		&models7.Skill{},
		&models3.Contact{},
		&models5.WorkExperience{},
		&models4.Education{},
		&models2.Certificate{},
	)
}
