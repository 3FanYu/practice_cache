package dao

import (
	"github.com/3fanyu/glossika/internal/models"
	"gorm.io/gorm"
)

type EmailDAO struct {
	db *gorm.DB
}

func NewEmailDAO(db *gorm.DB) *EmailDAO {
	return &EmailDAO{db: db}
}

// CreateEmail inserts a new Email into the database
func (dao *EmailDAO) CreateEmail(email *models.Email) error {
	result := dao.db.Create(&email) // Pass pointer of data to Create
	return result.Error
}

// GetUserByEmail finds a user by email
func (dao *EmailDAO) GetEmailByTargetAddress(db *gorm.DB, targetAddress string) (*[]models.Email, error) {
	var emails []models.Email
	result := db.Where("target_address = ?", targetAddress).Find(&emails)
	if result.Error != nil {
		return nil, result.Error
	}
	return &emails, nil
}
