package main

import (
	"github.com/3fanyu/glossika/internal/dao"
	"github.com/3fanyu/glossika/internal/routers/emails"
	"github.com/3fanyu/glossika/internal/routers/items"
	"github.com/3fanyu/glossika/internal/routers/users"
	emailsUC "github.com/3fanyu/glossika/internal/usecase/emails"
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

	r := gin.Default()

	// Register module routes
	users.RegisterRoutes(r, usersUC.NewUsecase(userDAO, emailDAO))
	items.RegisterRoutes(r, itemsUC.NewUsecase(cache, itemDAO))
	emails.RegisterRoutes(r, emailsUC.NewUsecase(emailDAO))

	r.Run(":3000")
}
