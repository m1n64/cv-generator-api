package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Template struct {
	ID             uuid.UUID `gorm:"type:uuid;primaryKey"`
	TemplateOrigin string    `gorm:"type:varchar(255);not null"`
	IsDefault      bool      `gorm:"type:boolean;not null;default:false"`
	IsPremium      bool      `gorm:"type:boolean;not null;default:false"`
}

func (u *Template) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()

	return
}
