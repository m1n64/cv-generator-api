package services

import (
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"information-service/internal/contacts/models"
	"information-service/internal/contacts/repositories"
)

type ContactService struct {
	contactRepo repositories.ContactRepository
	db          *gorm.DB
}

func NewContactService(contactRepo repositories.ContactRepository, db *gorm.DB) *ContactService {
	return &ContactService{
		contactRepo: contactRepo,
		db:          db,
	}
}

func (s *ContactService) GetContacts(cvID uuid.UUID) ([]*models.Contact, error) {
	contacts, err := s.contactRepo.GetContactsByCvID(cvID)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return contacts, nil
}

func (s *ContactService) GetContact(id uuid.UUID, cvID uuid.UUID) (*models.Contact, error) {
	contact, err := s.contactRepo.GetContact(id, cvID)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return contact, nil
}

func (s *ContactService) CreateContact(cvID uuid.UUID, title string, link *string) (*models.Contact, error) {
	contact := &models.Contact{
		CvID:  cvID,
		Title: title,
		Link:  link,
	}

	err := s.db.Transaction(func(tx *gorm.DB) error {
		var err error
		contact, err = s.contactRepo.CreateContact(contact)

		return err
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return contact, nil
}

func (s *ContactService) UpdateContact(contactID uuid.UUID, cvID uuid.UUID, title string, link *string) (*models.Contact, error) {
	contact := &models.Contact{
		ID:    contactID,
		CvID:  cvID,
		Title: title,
		Link:  link,
	}

	err := s.db.Transaction(func(tx *gorm.DB) error {
		var err error
		contact, err = s.contactRepo.UpdateContact(contactID, contact)

		return err
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return contact, nil
}

func (s *ContactService) DeleteContact(id uuid.UUID, cvID uuid.UUID) error {
	err := s.db.Transaction(func(tx *gorm.DB) error {
		return s.contactRepo.DeleteContact(id, cvID)
	})

	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	return nil
}

func (s *ContactService) DeleteContactsByCvID(cvID uuid.UUID) error {
	err := s.db.Transaction(func(tx *gorm.DB) error {
		return s.contactRepo.DeleteContactsByCvID(cvID)
	})

	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	return nil
}
