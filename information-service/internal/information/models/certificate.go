package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Certificate struct {
	ID          uuid.UUID  `gorm:"type:uuid;primaryKey"`
	CvID        uuid.UUID  `gorm:"type:uuid;not null"`
	Title       string     `gorm:"not null"`
	Vendor      string     `gorm:"not null"`
	StartDate   time.Time  `gorm:"type:date;not null"`
	EndDate     *time.Time `gorm:"type:date"`
	Description *string    `gorm:"type:text"`
	gorm.Model
}

func (e *Certificate) BeforeCreate(tx *gorm.DB) (err error) {
	e.ID = uuid.New()
	return
}
