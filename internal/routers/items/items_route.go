package items

import (
	business "github.com/3fanyu/glossika/internal/business/items"
	"github.com/3fanyu/glossika/internal/routers/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(router *gin.Engine, db *gorm.DB) {
	itemGroup := router.Group("/items")
	itemGroup.Use(middleware.AuthMiddleware())
	{
		itemGroup.GET("/", GetItems(db))
	}
}

func GetItems(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		business.List(c, db)
	}
}
