package handlers

import (
	"context"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	certificates "information-service/internal/certificates/grpc"
	"information-service/internal/certificates/models"
	"information-service/internal/certificates/services"
	"information-service/pkg/utils"
)

type CertificateServiceServer struct {
	certificates.UnimplementedCertificatesServiceServer
	certificateService *services.CertificateService
	logger             *zap.Logger
}

func NewCertificateServiceServer(certificateService *services.CertificateService, logger *zap.Logger) *CertificateServiceServer {
	return &CertificateServiceServer{
		certificateService: certificateService,
		logger:             logger,
	}
}

func (s *CertificateServiceServer) GetCertificates(ctx context.Context, req *certificates.GetCertificatesRequest) (*certificates.AllCertificatesResponse, error) {
	if uuid.Validate(req.CvId) != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid cv_id")
	}

	certificatesList, err := s.certificateService.GetCertificates(uuid.MustParse(req.CvId))
	if err != nil {
		s.logger.Error("error getting certificates", zap.Error(err))
		return nil, err
	}

	var certResp []*certificates.CertificateResponse
	for _, cert := range certificatesList {
		certResp = append(certResp, s.getCertificateResponse(cert))
	}

	return &certificates.AllCertificatesResponse{Certificates: certResp}, nil
}

func (s *CertificateServiceServer) GetCertificateByID(ctx context.Context, req *certificates.GetCertificateByIDRequest) (*certificates.CertificateResponse, error) {
	if uuid.Validate(req.CvId) != nil || uuid.Validate(req.Id) != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid cv_id and id")
	}

	certificate, err := s.certificateService.GetCertificateByID(uuid.MustParse(req.Id), uuid.MustParse(req.CvId))
	if err != nil {
		s.logger.Error("error getting certificate", zap.Error(err))
		return nil, err
	}

	return s.getCertificateResponse(certificate), nil
}

func (s *CertificateServiceServer) CreateCertificate(ctx context.Context, req *certificates.CreateCertificateRequest) (*certificates.CertificateResponse, error) {
	if uuid.Validate(req.CvId) != nil || req.Title == "" || req.Vendor == "" || req.StartDate == "" {
		return nil, status.Error(codes.InvalidArgument, "invalid cv_id, title, vendor or start_date")
	}

	startDate, endDate, err := utils.ParseStartAndEndDate(req.StartDate, req.EndDate)
	if err != nil {
		return nil, err
	}

	certificate, err := s.certificateService.CreateCertificate(uuid.MustParse(req.CvId), req.Title, req.Vendor, startDate, endDate, req.Description)
	if err != nil {
		s.logger.Error("error creating certificate", zap.Error(err))
		return nil, err
	}

	return s.getCertificateResponse(certificate), nil
}

func (s *CertificateServiceServer) UpdateCertificateByID(ctx context.Context, req *certificates.UpdateCertificateByIDRequest) (*certificates.CertificateResponse, error) {
	if uuid.Validate(req.Id) != nil || uuid.Validate(req.CvId) != nil || req.Title == "" || req.Vendor == "" || req.StartDate == "" {
		return nil, status.Error(codes.InvalidArgument, "invalid id, cv_id, title, vendor or start_date")
	}

	startDate, endDate, err := utils.ParseStartAndEndDate(req.StartDate, req.EndDate)
	if err != nil {
		return nil, err
	}

	certificate, err := s.certificateService.UpdateCertificate(uuid.MustParse(req.Id), uuid.MustParse(req.CvId), req.Title, req.Vendor, startDate, endDate, req.Description)
	if err != nil {
		s.logger.Error("error updating certificate", zap.Error(err))
		return nil, err
	}

	return s.getCertificateResponse(certificate), nil
}

func (s *CertificateServiceServer) DeleteCertificateByID(ctx context.Context, req *certificates.DeleteCertificateByIDRequest) (*certificates.DeleteCertificateByIDResponse, error) {
	if uuid.Validate(req.Id) != nil || uuid.Validate(req.CvId) != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid id and cv_id")
	}

	err := s.certificateService.DeleteCertificate(uuid.MustParse(req.Id), uuid.MustParse(req.CvId))
	if err != nil {
		s.logger.Error("error deleting certificate", zap.Error(err))
		return nil, err
	}

	return &certificates.DeleteCertificateByIDResponse{
		Success: true,
	}, nil
}

func (s *CertificateServiceServer) getCertificateResponse(certificate *models.Certificate) *certificates.CertificateResponse {
	return &certificates.CertificateResponse{
		Id:          certificate.ID.String(),
		CvId:        certificate.CvID.String(),
		Title:       certificate.Title,
		Vendor:      certificate.Vendor,
		StartDate:   certificate.StartDate.String(),
		EndDate:     utils.PtrToTimeString(certificate.EndDate),
		Description: certificate.Description,
		CreatedAt:   certificate.CreatedAt.String(),
		UpdatedAt:   utils.PtrToTimeString(&certificate.UpdatedAt),
	}
}
