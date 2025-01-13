package utils

import "cv-generator-service/internal/generator/models"

func StartMigrations() {
	db := GetDBConnection()
	db.AutoMigrate(&models.GeneratedPdf{})
}
