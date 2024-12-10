package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CV struct {
	ID         uuid.UUID `gorm:"type:uuid;primaryKey"`
	ExternalID uuid.UUID `gorm:"type:uuid;unique;not null"`
	UserID     uuid.UUID `gorm:"type:uuid;not null"`
	Title      string    `gorm:"not null"`
	gorm.Model
}

func (u *CV) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	u.ExternalID = uuid.New()
	return
}
