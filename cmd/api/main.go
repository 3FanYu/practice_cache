package main

import (
	"github.com/3fanyu/glossika/internal/dao"
	"github.com/3fanyu/glossika/internal/routers/items"
	"github.com/3fanyu/glossika/internal/routers/users"
	itemsUC "github.com/3fanyu/glossika/internal/usecase/items"
	usersUC "github.com/3fanyu/glossika/internal/usecase/users"
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

	//Usecases

	r := gin.Default()

	// Register module routes
	users.RegisterRoutes(r, usersUC.NewUsecase(userDAO, emailDAO))
	items.RegisterRoutes(r, itemsUC.NewUsecase(cache, itemDAO))

	r.Run(":3000")
}
