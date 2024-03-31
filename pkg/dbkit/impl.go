package dbkit

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewGormDB(dsn string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	return db
}
