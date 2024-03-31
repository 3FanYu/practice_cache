package dao

import (
	"github.com/3fanyu/glossika/internal/models"
	"gorm.io/gorm"
)

type ItemDAO struct {
	db *gorm.DB
}

func NewItemDAO(db *gorm.DB) *ItemDAO {
	return &ItemDAO{db: db}
}

type ListOptions struct {
	Page int `form:"page"`
	Size int `form:"size"`
}

// GetItems gets all items from the database
func (dao *ItemDAO) GetItems(filter ListOptions) ([]models.Item, error) {
	var items []models.Item
	result := dao.db.Find(&items)
	return items, result.Error
}
