package models

import "gorm.io/gorm"

type Email struct {
	gorm.Model
	TargetAddress string         `gorm:"type:varchar(255);uniqueIndex"`
	VerifyLink    string         `gorm:"type:varchar(255)"`
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}
