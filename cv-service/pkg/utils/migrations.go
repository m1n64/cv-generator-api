package utils

import "cv-service/internal/cv/models"

func StartMigrations() {
	db := GetDBConnection()
	db.AutoMigrate(&models.CV{})
}
