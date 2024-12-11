package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Information struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey"`
	CvID        uuid.UUID `gorm:"type:uuid;not null"`
	FullName    string    `gorm:"type:varchar(255);not null"`
	PhotoFileID *string   `gorm:"type:varchar(255)"`
	Position    *string   `gorm:"type:varchar(255)"`
	Location    *string   `gorm:"type:varchar(255)"`
	Biography   *string   `gorm:"type:text"`
	gorm.Model
}

func (u *Information) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}
