package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type WorkExperience struct {
	ID          uuid.UUID  `gorm:"type:uuid;primaryKey"`
	CvID        uuid.UUID  `gorm:"type:uuid;not null"`
	Company     string     `gorm:"not null"`
	Position    string     `gorm:"not null"`
	StartDate   time.Time  `gorm:"type:date;not null"`
	EndDate     *time.Time `gorm:"type:date"`
	Location    string     `gorm:"not null"`
	Description string     `gorm:"type:text;not null"`
	gorm.Model
}

func (u *WorkExperience) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}
