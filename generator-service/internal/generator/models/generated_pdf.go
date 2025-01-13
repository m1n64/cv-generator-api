package models

import (
	"cv-generator-service/internal/generator/enums"
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GeneratedPdf struct {
	ID         uuid.UUID        `gorm:"type:uuid;primaryKey"`
	CvID       uuid.UUID        `gorm:"type:uuid;not null"`
	UserID     uuid.UUID        `gorm:"type:uuid;not null"`
	Title      string           `gorm:"type:varchar(255);not null"`
	FileOrigin *string          `gorm:"type:varchar(255);nullable"`
	Status     enums.StatusType `gorm:"type:varchar(255);not null"`
	gorm.Model
}

func (u *GeneratedPdf) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()

	if !u.Status.IsValid() {
		return errors.New("invalid status value")
	}

	return
}

func (u *GeneratedPdf) BeforeUpdate(tx *gorm.DB) (err error) {
	if !u.Status.IsValid() {
		return errors.New("invalid status value")
	}

	return
}
