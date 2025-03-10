package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Education struct {
	ID          uuid.UUID  `gorm:"type:uuid;primaryKey"`
	CvID        uuid.UUID  `gorm:"type:uuid;not null"`
	Institution string     `gorm:"not null"`
	StartDate   time.Time  `gorm:"type:date;not null"`
	EndDate     *time.Time `gorm:"type:date"`
	Location    string     `gorm:"not null"`
	Faculty     string     `gorm:"not null"`
	Degree      *string    `gorm:"type:varchar(255)"`
	Description *string    `gorm:"type:text"`
	gorm.Model
}

func (e *Education) BeforeCreate(tx *gorm.DB) (err error) {
	e.ID = uuid.New()
	return
}
