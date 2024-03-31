package dao

import (
	"github.com/3fanyu/glossika/internal/models"
	"gorm.io/gorm"
)

type UserDAO struct {
	db *gorm.DB
}

func NewUserDAO(db *gorm.DB) *UserDAO {
	return &UserDAO{db: db}
}

// CreateUser inserts a new User into the database
func (dao *UserDAO) CreateUser(user *models.User) error {
	result := dao.db.Create(&user) // Pass pointer of data to Create
	return result.Error
}

// GetUserByEmail finds a user by email
func (dao *UserDAO)GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	result := dao.db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
