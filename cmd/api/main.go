package main

import (
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
)

func main() {
	db := dbkit.NewGormDB("root:root@tcp(db:3306)/glossika?charset=utf8mb4&parseTime=True&loc=Local")
	cache := cachekit.NewCache("redis:6379")

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

	r.Run(":3000")
}
