package models

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	Name      string         `gorm:"type:varchar(255)"`
	Quantity  string         `gorm:"type:varchar(255)"`
	Category  string         `gorm:"type:varchar(255)"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
