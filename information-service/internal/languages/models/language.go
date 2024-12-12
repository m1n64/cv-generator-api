package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Language struct {
	ID    uuid.UUID `gorm:"type:uuid;primaryKey"`
	CvID  uuid.UUID `gorm:"type:uuid;not null"`
	Name  string    `gorm:"type:varchar(255);not null"`
	Level string    `gorm:"type:varchar(255);not null"`
	gorm.Model
}

func (u *Language) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}
