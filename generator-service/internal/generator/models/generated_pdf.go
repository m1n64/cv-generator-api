package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GeneratedPDF struct {
	ID         uuid.UUID `gorm:"type:uuid;primaryKey"`
	CvID       uuid.UUID `gorm:"type:uuid;not null"`
	UserID     uuid.UUID `gorm:"type:uuid;not null"`
	Title      string    `gorm:"type:varchar(255);not null"`
	FileOrigin string    `gorm:"type:varchar(255);not null"`
	gorm.Model
}

func (u *GeneratedPDF) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}
