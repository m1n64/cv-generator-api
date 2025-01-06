package services

import (
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"information-service/internal/skills/models"
	repositories3 "information-service/internal/skills/repositories"
)

type SkillService struct {
	skillRepo             repositories3.SkillRepository
	skillAnalyticsService *SkillsAnalyticsService
	db                    *gorm.DB
}

func NewSkillService(skillRepo repositories3.SkillRepository, skillAnalyticsService *SkillsAnalyticsService, db *gorm.DB) *SkillService {
	return &SkillService{
		skillRepo:             skillRepo,
		skillAnalyticsService: skillAnalyticsService,
		db:                    db,
	}
}

func (s *SkillService) GetSkills(cvID uuid.UUID) ([]*models.Skill, error) {
	skills, err := s.skillRepo.GetSkillsByCvID(cvID)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return skills, nil
}

func (s *SkillService) GetSkill(ID uuid.UUID, cvID uuid.UUID) (*models.Skill, error) {
	skill, err := s.skillRepo.GetSkill(ID, cvID)
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return skill, nil
}

func (s *SkillService) CreateSkill(cvID uuid.UUID, name string) (*models.Skill, error) {
	skill := &models.Skill{
		CvID: cvID,
		Name: name,
	}

	err := s.db.Transaction(func(tx *gorm.DB) error {
		var err error
		skill, err = s.skillRepo.CreateSkill(skill)
		if err != nil {
			return err
		}

		s.skillAnalyticsService.SendCreateEvent(skill)

		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return skill, nil
}

func (s *SkillService) UpdateSkill(ID uuid.UUID, cvID uuid.UUID, name string) (*models.Skill, error) {
	skill := &models.Skill{
		CvID: cvID,
		Name: name,
	}

	err := s.db.Transaction(func(tx *gorm.DB) error {
		var err error
		skill, err = s.skillRepo.UpdateSkill(ID, skill)
		if err != nil {
			return err
		}

		s.skillAnalyticsService.SendUpdateEvent(skill)

		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return skill, nil
}

func (s *SkillService) DeleteSkill(ID uuid.UUID, cvID uuid.UUID) error {
	err := s.db.Transaction(func(tx *gorm.DB) error {
		return s.skillRepo.DeleteSkill(ID, cvID)
	})

	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	return nil
}

func (s *SkillService) DeleteSkillsByCvID(cvID uuid.UUID) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		return s.skillRepo.DeleteSkillsByCvID(cvID)
	})
}
