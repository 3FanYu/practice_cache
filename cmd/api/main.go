package main

import (
	"log"
	"os"

	"github.com/3fanyu/glossika/internal/dao"
	"github.com/3fanyu/glossika/internal/routers/emails"
	"github.com/3fanyu/glossika/internal/routers/items"
	"github.com/3fanyu/glossika/internal/routers/middleware"
	"github.com/3fanyu/glossika/internal/routers/users"
	emailsUsecase "github.com/3fanyu/glossika/internal/usecase/emails"
	itemsUsecase "github.com/3fanyu/glossika/internal/usecase/items"
	usersUsecase "github.com/3fanyu/glossika/internal/usecase/users"
	"github.com/3fanyu/glossika/pkg/cachekit"
	"github.com/3fanyu/glossika/pkg/dbkit"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db := dbkit.NewGormDB(os.Getenv("MYSQL_DSN") + "?charset=utf8mb4&parseTime=True&loc=Local")
	cache := cachekit.NewCache(os.Getenv("REDIS_ADDR"))

	//DAOs
	userDAO := dao.NewUserDAO(db)
	itemDAO := dao.NewItemDAO(db)
	emailDAO := dao.NewEmailDAO(db)

	//usecases
	usersUC := usersUsecase.NewUsecase(userDAO, emailDAO)
	itemsUC := itemsUsecase.NewUsecase(cache, itemDAO)
	emailsUC := emailsUsecase.NewUsecase(emailDAO)

	r := gin.Default()

	// Register module routes
	users.RegisterRoutes(r, usersUC)
	items.RegisterRoutes(r, itemsUC, middleware.AuthMiddleware)
	emails.RegisterRoutes(r, emailsUC)

	r.Run(os.Getenv("APP_ADDR"))
}
