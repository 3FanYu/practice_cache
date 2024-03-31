package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email             string         `gorm:"type:varchar(255);uniqueIndex"`
	EncryptedPassword string         `gorm:"type:varchar(255)"`
	VerifyToken       *string        `gorm:"type:varchar(255);index"`
	DeletedAt         gorm.DeletedAt `gorm:"index"`
}
