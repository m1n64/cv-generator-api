package handlers

import (
	"context"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	skills "information-service/internal/skills/grpc"
	"information-service/internal/skills/models"
	"information-service/internal/skills/services"
)

type SkillServiceServer struct {
	skills.UnimplementedSkillsServiceServer
	skillService *services.SkillService
	logger       *zap.Logger
}

func NewSkillServiceServer(skillService *services.SkillService, logger *zap.Logger) *SkillServiceServer {
	return &SkillServiceServer{
		skillService: skillService,
		logger:       logger,
	}
}

func (s *SkillServiceServer) GetSkills(ctx context.Context, request *skills.GetSkillsRequest) (*skills.AllSkillsResponse, error) {
	if uuid.Validate(request.CvId) != nil {
		return nil, status.Error(codes.InvalidArgument, "cv_id is required and must be a valid uuid")
	}

	skillsList, err := s.skillService.GetSkills(uuid.MustParse(request.CvId))
	if err != nil {
		s.logger.Error("error getting skills", zap.Error(err))
		return nil, err
	}

	var skillsResp []*skills.SkillResponse
	for _, skill := range skillsList {
		skillsResp = append(skillsResp, s.getSkillResponse(skill))
	}
	return &skills.AllSkillsResponse{Skills: skillsResp}, nil
}

func (s *SkillServiceServer) GetSkillByID(ctx context.Context, request *skills.GetSkillByIDRequest) (*skills.SkillResponse, error) {
	if uuid.Validate(request.Id) != nil || uuid.Validate(request.CvId) != nil {
		return nil, status.Error(codes.InvalidArgument, "cv_id and id is required and must be a valid uuid")
	}

	skill, err := s.skillService.GetSkill(uuid.MustParse(request.Id), uuid.MustParse(request.CvId))
	if err != nil {
		s.logger.Error("error getting skill", zap.Error(err))
		return nil, err
	}

	return s.getSkillResponse(skill), nil
}

func (s *SkillServiceServer) CreateSkill(ctx context.Context, request *skills.CreateSkillRequest) (*skills.SkillResponse, error) {
	if request.Name == "" || uuid.Validate(request.CvId) != nil {
		return nil, status.Error(codes.InvalidArgument, "cv_id and name is required and must be a valid uuid")
	}

	skill, err := s.skillService.CreateSkill(uuid.MustParse(request.CvId), request.Name)
	if err != nil {
		s.logger.Error("error creating skill", zap.Error(err))
		return nil, err
	}

	return s.getSkillResponse(skill), nil
}

func (s *SkillServiceServer) UpdateSkillByID(ctx context.Context, request *skills.UpdateSkillByIDRequest) (*skills.SkillResponse, error) {
	if uuid.Validate(request.Id) != nil || uuid.Validate(request.CvId) != nil || request.Name == "" {
		return nil, status.Error(codes.InvalidArgument, "cv_id, id and name is required and must be a valid uuid")
	}

	skill, err := s.skillService.UpdateSkill(uuid.MustParse(request.Id), uuid.MustParse(request.CvId), request.Name)
	if err != nil {
		s.logger.Error("error updating skill", zap.Error(err))
		return nil, err
	}

	return s.getSkillResponse(skill), nil
}

func (s *SkillServiceServer) DeleteSkillByID(ctx context.Context, request *skills.DeleteSkillByIDRequest) (*skills.DeleteSkillByIDResponse, error) {
	if uuid.Validate(request.Id) != nil || uuid.Validate(request.CvId) != nil {
		return nil, status.Error(codes.InvalidArgument, "cv_id and id is required and must be a valid uuid")
	}

	err := s.skillService.DeleteSkill(uuid.MustParse(request.Id), uuid.MustParse(request.CvId))
	if err != nil {
		s.logger.Error("error deleting skill", zap.Error(err))
		return nil, err
	}

	return &skills.DeleteSkillByIDResponse{Success: true}, nil
}

func (s *SkillServiceServer) getSkillResponse(skill *models.Skill) *skills.SkillResponse {
	return &skills.SkillResponse{
		Id:        skill.ID.String(),
		CvId:      skill.CvID.String(),
		Name:      skill.Name,
		CreatedAt: skill.CreatedAt.String(),
		UpdatedAt: skill.UpdatedAt.String(),
	}
}
