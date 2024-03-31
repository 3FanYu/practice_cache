package main

import (
	"log"

	"github.com/3fanyu/glossika/internal/routers/items"
	"github.com/3fanyu/glossika/internal/routers/users"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db := connectDatabase()
	r := gin.Default()

	// Register module routes
	users.RegisterRoutes(r, db)
	items.RegisterRoutes(r, db)

	r.Run(":3000")
}

func connectDatabase() *gorm.DB {
	dsn := "root:root@tcp(db:3306)/glossika?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	return db
}
