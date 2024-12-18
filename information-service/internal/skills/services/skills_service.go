package services

import (
	"gorm.io/gorm"
	repositories3 "information-service/internal/skills/repositories"
)

type SkillService struct {
	skillRepo repositories3.SkillRepository
	db        *gorm.DB
}

func NewSkillService(skillRepo repositories3.SkillRepository, db *gorm.DB) *SkillService {
	return &SkillService{
		skillRepo: skillRepo,
		db:        db,
	}
}
