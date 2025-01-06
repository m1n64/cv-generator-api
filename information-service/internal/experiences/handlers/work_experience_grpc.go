package handlers

import (
	"context"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	experiences "information-service/internal/experiences/grpc"
	"information-service/internal/experiences/models"
	services6 "information-service/internal/experiences/services"
	"information-service/pkg/utils"
)

type WorkExperienceServiceServer struct {
	experiences.UnimplementedExperiencesServiceServer
	workExperienceService *services6.WorkExperienceService
	logger                *zap.Logger
}

func NewWorkExperienceServiceServer(workExperienceService *services6.WorkExperienceService, logger *zap.Logger) *WorkExperienceServiceServer {
	return &WorkExperienceServiceServer{
		workExperienceService: workExperienceService,
		logger:                logger,
	}
}

func (s *WorkExperienceServiceServer) GetExperiences(ctx context.Context, req *experiences.GetExperiencesRequest) (*experiences.AllExperiencesResponse, error) {
	if uuid.Validate(req.CvId) != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid uuid")
	}

	expList, err := s.workExperienceService.GetWorkExperiencesByCvID(uuid.MustParse(req.CvId))
	if err != nil {
		s.logger.Error("error getting experiences", zap.Error(err))
		return nil, err
	}

	var experienceResponse []*experiences.ExperienceResponse
	for _, experience := range expList {
		experienceResponse = append(experienceResponse, s.getWorkExperienceResponse(experience))
	}

	return &experiences.AllExperiencesResponse{
		WorkExperiences: experienceResponse,
	}, nil
}

func (s *WorkExperienceServiceServer) GetExperienceByID(ctx context.Context, req *experiences.GetExperienceByIDRequest) (*experiences.ExperienceResponse, error) {
	if uuid.Validate(req.Id) != nil || uuid.Validate(req.CvId) != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid id and cv_id")
	}

	experience, err := s.workExperienceService.GetWorkExperience(uuid.MustParse(req.Id), uuid.MustParse(req.CvId))
	if err != nil {
		s.logger.Error("error getting experience", zap.Error(err))
		return nil, err
	}

	return s.getWorkExperienceResponse(experience), nil
}

func (s *WorkExperienceServiceServer) CreateExperience(ctx context.Context, req *experiences.CreateExperienceRequest) (*experiences.ExperienceResponse, error) {
	if uuid.Validate(req.CvId) != nil || req.Company == "" || req.Position == "" || req.StartDate == "" || req.Location == "" || req.Description == "" {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	startDate, endDate, err := utils.ParseStartAndEndDate(req.StartDate, req.EndDate)
	if err != nil {
		return nil, err
	}

	experience, err := s.workExperienceService.CreateWorkExperience(uuid.MustParse(req.CvId), req.Company, req.Position, startDate, endDate, req.Location, req.Description)
	if err != nil {
		s.logger.Error("error creating experience", zap.Error(err))
		return nil, err
	}

	return s.getWorkExperienceResponse(experience), nil
}

func (s *WorkExperienceServiceServer) UpdateExperienceByID(ctx context.Context, req *experiences.UpdateExperienceByIDRequest) (*experiences.ExperienceResponse, error) {
	if uuid.Validate(req.Id) != nil || uuid.Validate(req.CvId) != nil || req.Company == "" || req.Position == "" || req.StartDate == "" || req.Location == "" || req.Description == "" {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	startDate, endDate, err := utils.ParseStartAndEndDate(req.StartDate, req.EndDate)
	if err != nil {
		return nil, err
	}

	experience, err := s.workExperienceService.UpdateWorkExperience(uuid.MustParse(req.Id), uuid.MustParse(req.CvId), req.Company, req.Position, startDate, endDate, req.Location, req.Description)
	if err != nil {
		s.logger.Error("error updating experience", zap.Error(err))
		return nil, err
	}

	return s.getWorkExperienceResponse(experience), nil
}

func (s *WorkExperienceServiceServer) DeleteExperienceByID(ctx context.Context, req *experiences.DeleteExperienceByIDRequest) (*experiences.DeleteExperienceByIDResponse, error) {
	if uuid.Validate(req.Id) != nil || uuid.Validate(req.CvId) != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid id and cv_id")
	}

	err := s.workExperienceService.DeleteWorkExperience(uuid.MustParse(req.Id), uuid.MustParse(req.CvId))
	if err != nil {
		s.logger.Error("error deleting experience", zap.Error(err))
		return nil, err
	}

	return &experiences.DeleteExperienceByIDResponse{Success: true}, nil
}

func (s *WorkExperienceServiceServer) getWorkExperienceResponse(experience *models.WorkExperience) *experiences.ExperienceResponse {
	return &experiences.ExperienceResponse{
		Id:          experience.ID.String(),
		CvId:        experience.CvID.String(),
		Company:     experience.Company,
		Position:    experience.Position,
		StartDate:   *utils.PtrToTimeString(&experience.StartDate),
		EndDate:     utils.PtrToTimeString(experience.EndDate),
		Location:    experience.Location,
		Description: experience.Description,
		CreatedAt:   experience.CreatedAt.String(),
		UpdatedAt:   experience.UpdatedAt.String(),
	}
}
