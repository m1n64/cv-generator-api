package repositories

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"information-service/internal/contacts/models"
)

type contactRepository struct {
	db *gorm.DB
}

func NewContactRepository(db *gorm.DB) ContactRepository {
	return &contactRepository{
		db: db,
	}
}

func (r *contactRepository) CreateContact(contact *models.Contact) (*models.Contact, error) {
	if err := r.db.Create(contact).Error; err != nil {
		return nil, err
	}

	return contact, nil
}

func (r *contactRepository) GetContact(id uuid.UUID, cvID uuid.UUID) (*models.Contact, error) {
	var contact models.Contact
	if err := r.db.Where("id = ? AND cv_id = ?", id, cvID).First(&contact).Error; err != nil {
		return nil, err
	}

	return &contact, nil
}

func (r *contactRepository) GetContactsByCvID(cvID uuid.UUID) ([]*models.Contact, error) {
	var contacts []*models.Contact
	if err := r.db.Where("cv_id = ?", cvID).Find(&contacts).Error; err != nil {
		return nil, err
	}

	return contacts, nil
}

func (r *contactRepository) UpdateContact(id uuid.UUID, contact *models.Contact) (*models.Contact, error) {
	var existingContact models.Contact
	if err := r.db.First(&existingContact, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, err
	}

	if err := r.db.Model(&existingContact).Updates(contact).Error; err != nil {
		return nil, err
	}

	return contact, nil
}

func (r *contactRepository) DeleteContact(id uuid.UUID, cvID uuid.UUID) error {
	return r.db.Where("id = ? AND cv_id = ?", id, cvID).Delete(&models.Contact{}).Error
}

func (r *contactRepository) DeleteContactsByCvID(cvID uuid.UUID) error {
	return r.db.Where("cv_id = ?", cvID).Delete(&models.Contact{}).Error
}
