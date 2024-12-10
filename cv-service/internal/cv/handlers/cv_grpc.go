package handlers

import (
	"context"
	"cv-service/internal/cv/grpc/cv"
	"cv-service/internal/cv/models"
	"cv-service/internal/cv/repositories"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"time"
)

type CVServiceServer struct {
	cv.UnimplementedCVServiceServer
	cvRepo repositories.CVRepository
}

func NewCVServiceServer(cvRepo repositories.CVRepository) *CVServiceServer {
	return &CVServiceServer{
		cvRepo: cvRepo,
	}
}

func (s *CVServiceServer) CreateCV(ctx context.Context, req *cv.CreateCVRequest) (*cv.CVResponse, error) {
	if req.UserId == "" || req.Name == "" {
		return nil, status.Error(codes.InvalidArgument, "userID and name are required")
	}

	cvModel := &models.CV{
		UserID: uuid.MustParse(req.UserId),
		Title:  req.Name,
	}

	if err := s.cvRepo.CreateCV(cvModel); err != nil {
		log.Printf("Error saving CV: %v", err)
		return nil, status.Error(codes.Internal, err.Error())
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

	cvs, err := s.cvRepo.GetAllCVsByUserID(uuid.MustParse(req.UserId))
	if err != nil {
		log.Printf("Error getting CVs: %v", err)
		return nil, status.Error(codes.Internal, err.Error())
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

	cvModel, err := s.cvRepo.GetCVByID(uuid.MustParse(req.CvId))
	if err != nil {
		log.Printf("Error getting CV: %v", err)
		return nil, status.Error(codes.Internal, err.Error())
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

	err := s.cvRepo.DeleteCVByID(uuid.MustParse(req.CvId))
	if err != nil {
		log.Printf("Error getting CV: %v", err)
		return nil, status.Error(codes.Internal, err.Error())
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

	updatedCV := &models.CV{
		Title: req.Name,
	}

	err := s.cvRepo.UpdateCVByID(cvId, updatedCV)
	if err != nil {
		log.Printf("Error getting CV: %v", err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	updatedCVFromDB, err := s.cvRepo.GetCVByID(cvId)
	if err != nil {
		log.Printf("Error fetching updated CV: %v", err)
		return nil, status.Error(codes.Internal, "failed to fetch updated CV")
	}

	return &cv.CVResponse{
		Id:         updatedCVFromDB.ID.String(),
		ExternalId: updatedCVFromDB.ExternalID.String(),
		Name:       updatedCVFromDB.Title,
		CreatedAt:  updatedCVFromDB.CreatedAt.String(),
	}, nil
}

func (s *CVServiceServer) GetOriginalID(ctx context.Context, req *cv.GetOriginalIDRequest) (*cv.GetOriginalIDResponse, error) {
	if req.UserId == "" || req.CvId == "" {
		return nil, status.Error(codes.InvalidArgument, "userID and cvID is required")
	}

	originalId, err := s.cvRepo.GetOriginalIDByExternalID(uuid.MustParse(req.CvId), uuid.MustParse(req.UserId))
	if err != nil {
		log.Printf("Error getting CV: %v", err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &cv.GetOriginalIDResponse{
		Id: originalId.String(),
	}, nil
}
