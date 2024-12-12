package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Contact struct {
	ID    uuid.UUID `gorm:"type:uuid;primaryKey"`
	CvID  uuid.UUID `gorm:"type:uuid;not null"`
	Title string    `gorm:"type:varchar(255);not null"`
	Link  *string   `gorm:"type:varchar(255)"`
	gorm.Model
}

func (u *Contact) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}
