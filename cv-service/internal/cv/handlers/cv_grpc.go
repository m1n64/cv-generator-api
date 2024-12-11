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
	cvService service.CVService
}

func NewCVServiceServer(cvService service.CVService) *CVServiceServer {
	return &CVServiceServer{
		cvService: cvService,
	}
}

func (s *CVServiceServer) CreateCV(ctx context.Context, req *cv.CreateCVRequest) (*cv.CVResponse, error) {
	if req.UserId == "" || req.Name == "" {
		return nil, status.Error(codes.InvalidArgument, "userID and name are required")
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
	if req.UserId == "" {
		return nil, status.Error(codes.InvalidArgument, "userID is required")
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
	if req.CvId == "" {
		return nil, status.Error(codes.InvalidArgument, "cvID is required")
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
	if req.CvId == "" {
		return nil, status.Error(codes.InvalidArgument, "cvID is required")
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
	if req.CvId == "" || req.Name == "" {
		return nil, status.Error(codes.InvalidArgument, "cvID and name is required")
	}

	cvId := uuid.MustParse(req.CvId)

	cvModel, err := s.cvService.UpdateCV(cvId, req.Name)
	if err == nil {
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
	if req.UserId == "" || req.CvId == "" {
		return nil, status.Error(codes.InvalidArgument, "userID and cvID is required")
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
