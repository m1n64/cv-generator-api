package handlers

import (
	"context"
	"cv-service/internal/cv/grpc/cv"
	"cv-service/internal/cv/service"
	"cv-service/pkg/utils"
	"fmt"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

type CVServiceServer struct {
	cv.UnimplementedCVServiceServer
	cvService *service.CVService
}

func NewCVServiceServer(cvService *service.CVService) *CVServiceServer {
	return &CVServiceServer{
		cvService: cvService,
	}
}

func (s *CVServiceServer) CreateCV(ctx context.Context, req *cv.CreateCVRequest) (*cv.CVResponse, error) {
	if uuid.Validate(req.UserId) != nil || req.Name == "" {
		return nil, status.Error(codes.InvalidArgument, "userID and name are required and valid UUID")
	}

	cvModel, err := s.cvService.CreateCV(uuid.MustParse(req.UserId), req.Name)
	if err != nil {
		utils.GetLogger().Info(fmt.Sprintf("Error creating CV: %v", err))
		return nil, err
	}

	return &cv.CVResponse{
		Id:         cvModel.ID.String(),
		ExternalId: cvModel.ExternalID.String(),
		Name:       cvModel.Title,
		CreatedAt:  cvModel.CreatedAt.String(),
	}, nil
}

func (s *CVServiceServer) GetAllCVsByUserID(ctx context.Context, req *cv.GetAllCVsByUserIDRequest) (*cv.GetAllCVsResponse, error) {
	if uuid.Validate(req.UserId) != nil {
		return nil, status.Error(codes.InvalidArgument, "userID is required and must be valid UUID")
	}

	cvs, err := s.cvService.GetAllCVsByUserID(uuid.MustParse(req.UserId))
	if err != nil {
		utils.GetLogger().Info(fmt.Sprintf("Error creating CV: %v", err))

		return nil, err
	}

	var cvList []*cv.CV
	for _, c := range cvs {
		cvList = append(cvList, &cv.CV{
			Id:         c.ID.String(),
			ExternalId: c.ExternalID.String(),
			Name:       c.Title,
			CreatedAt:  c.CreatedAt.Format(time.RFC3339),
		})
	}

	return &cv.GetAllCVsResponse{
		CvList: cvList,
	}, nil
}

func (s *CVServiceServer) GetCVByID(ctx context.Context, req *cv.GetCVByIDRequest) (*cv.CVResponse, error) {
	if uuid.Validate(req.CvId) != nil {
		return nil, status.Error(codes.InvalidArgument, "cvID is required and must be valid UUID")
	}

	cvModel, err := s.cvService.GetCVByID(uuid.MustParse(req.CvId))
	if err != nil {
		utils.GetLogger().Info(fmt.Sprintf("Error getting CV: %v", err))
		return nil, err
	}

	return &cv.CVResponse{
		Id:         cvModel.ID.String(),
		ExternalId: cvModel.ExternalID.String(),
		Name:       cvModel.Title,
		CreatedAt:  cvModel.CreatedAt.String(),
	}, nil
}

func (s *CVServiceServer) DeleteCVByID(ctx context.Context, req *cv.DeleteCVByIDRequest) (*cv.DeleteCVByIDResponse, error) {
	if uuid.Validate(req.CvId) != nil {
		return nil, status.Error(codes.InvalidArgument, "cvID is required and must be valid UUID")
	}

	err := s.cvService.DeleteCVByID(uuid.MustParse(req.CvId))
	if err != nil {
		utils.GetLogger().Info(fmt.Sprintf("Error deleting CV: %v", err))
		return nil, err
	}

	return &cv.DeleteCVByIDResponse{
		Success: true,
	}, nil
}

func (s *CVServiceServer) UpdateCV(ctx context.Context, req *cv.UpdateCVRequest) (*cv.CVResponse, error) {
	if req.Name == "" || uuid.Validate(req.CvId) != nil {
		return nil, status.Error(codes.InvalidArgument, "cvID and name is required and must be valid UUID")
	}

	cvId := uuid.MustParse(req.CvId)

	cvModel, err := s.cvService.UpdateCV(cvId, req.Name)
	if err != nil {
		utils.GetLogger().Info(fmt.Sprintf("Failed to fetch updated CV: %v", err))
		return nil, err
	}

	return &cv.CVResponse{
		Id:         cvModel.ID.String(),
		ExternalId: cvModel.ExternalID.String(),
		Name:       cvModel.Title,
		CreatedAt:  cvModel.CreatedAt.String(),
	}, nil
}

func (s *CVServiceServer) GetOriginalID(ctx context.Context, req *cv.GetOriginalIDRequest) (*cv.GetOriginalIDResponse, error) {
	if uuid.Validate(req.UserId) != nil || uuid.Validate(req.CvId) != nil {
		return nil, status.Error(codes.InvalidArgument, "userID and cvID is required and must be valid UUID")
	}

	originalId, err := s.cvService.GetOriginalID(uuid.MustParse(req.UserId), uuid.MustParse(req.CvId))
	if err != nil {
		utils.GetLogger().Info(fmt.Sprintf("Failed to fetch original ID: %v", err))
		return nil, err
	}

	return &cv.GetOriginalIDResponse{
		Id: originalId.String(),
	}, nil
}
