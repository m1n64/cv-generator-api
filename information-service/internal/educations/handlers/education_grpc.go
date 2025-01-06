package handlers

import (
	"context"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	educations "information-service/internal/educations/grpc"
	"information-service/internal/educations/models"
	"information-service/internal/educations/services"
	"information-service/pkg/utils"
)

type EducationServiceServer struct {
	educations.UnimplementedEducationServiceServer
	educationService *services.EducationService
	logger           *zap.Logger
}

func NewEducationServiceServer(educationService *services.EducationService, logger *zap.Logger) *EducationServiceServer {
	return &EducationServiceServer{
		educationService: educationService,
		logger:           logger,
	}
}

func (s *EducationServiceServer) GetEducations(ctx context.Context, req *educations.GetEducationsRequest) (*educations.AllEducationsResponse, error) {
	if uuid.Validate(req.CvId) != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid cv_id")
	}

	eduList, err := s.educationService.GetEducationsByCvID(uuid.MustParse(req.CvId))
	if err != nil {
		s.logger.Error("error getting educations", zap.Error(err))
		return nil, err
	}

	var educationResponse []*educations.EducationResponse
	for _, education := range eduList {
		educationResponse = append(educationResponse, s.getEducationResponse(education))
	}

	return &educations.AllEducationsResponse{
		Educations: educationResponse,
	}, nil
}

func (s *EducationServiceServer) GetEducationByID(ctx context.Context, req *educations.GetEducationByIDRequest) (*educations.EducationResponse, error) {
	if uuid.Validate(req.Id) != nil || uuid.Validate(req.CvId) != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid id and cv_id")
	}

	education, err := s.educationService.GetEducation(uuid.MustParse(req.Id), uuid.MustParse(req.CvId))
	if err != nil {
		s.logger.Error("error getting education", zap.Error(err))
		return nil, err
	}

	return s.getEducationResponse(education), nil
}

func (s *EducationServiceServer) CreateEducation(ctx context.Context, req *educations.CreateEducationRequest) (*educations.EducationResponse, error) {
	if uuid.Validate(req.CvId) != nil || req.Institution == "" || req.Location == "" || req.Faculty == "" || req.StartDate == "" {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	startDate, endDate, err := utils.ParseStartAndEndDate(req.StartDate, req.EndDate)
	if err != nil {
		return nil, err
	}

	education, err := s.educationService.CreateEducation(uuid.MustParse(req.CvId), req.Institution, req.Location, req.Faculty, startDate, endDate, req.Degree, req.Description)
	if err != nil {
		s.logger.Error("error creating education", zap.Error(err))
		return nil, err
	}

	return s.getEducationResponse(education), nil
}

func (s *EducationServiceServer) UpdateEducationByID(ctx context.Context, req *educations.UpdateEducationByIDRequest) (*educations.EducationResponse, error) {
	if uuid.Validate(req.Id) != nil || uuid.Validate(req.CvId) != nil || req.Institution == "" || req.Location == "" || req.Faculty == "" || req.StartDate == "" {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	startDate, endDate, err := utils.ParseStartAndEndDate(req.StartDate, req.EndDate)
	if err != nil {
		return nil, err
	}

	education, err := s.educationService.UpdateEducation(uuid.MustParse(req.Id), uuid.MustParse(req.CvId), req.Institution, req.Location, req.Faculty, startDate, endDate, req.Degree, req.Description)
	if err != nil {
		s.logger.Error("error updating education", zap.Error(err))
		return nil, err
	}

	return s.getEducationResponse(education), nil
}

func (s *EducationServiceServer) DeleteEducationByID(ctx context.Context, req *educations.DeleteEducationByIDRequest) (*educations.DeleteEducationByIDResponse, error) {
	if uuid.Validate(req.Id) != nil || uuid.Validate(req.CvId) != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	err := s.educationService.DeleteEducation(uuid.MustParse(req.Id), uuid.MustParse(req.CvId))
	if err != nil {
		s.logger.Error("error deleting education", zap.Error(err))
		return nil, err
	}

	return &educations.DeleteEducationByIDResponse{Success: true}, nil
}

func (s *EducationServiceServer) getEducationResponse(education *models.Education) *educations.EducationResponse {
	return &educations.EducationResponse{
		Id:          education.ID.String(),
		CvId:        education.CvID.String(),
		Institution: education.Institution,
		Location:    education.Location,
		Faculty:     education.Faculty,
		Degree:      education.Degree,
		StartDate:   *utils.PtrToTimeString(&education.StartDate),
		EndDate:     utils.PtrToTimeString(education.EndDate),
		Description: education.Description,
	}
}
