package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email             string `gorm:"type:varchar(255);uniqueIndex"`
	EncryptedPassword string `gorm:"type:varchar(255)"`
	Salt              string `gorm:"type:varchar(255);not null"`
	VerifyToken       string `gorm:"type:varchar(255)"`
	VerifiedAt        *time.Time
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         gorm.DeletedAt `gorm:"index"`
}
