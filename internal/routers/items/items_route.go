package items

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(router *gin.Engine, db *gorm.DB) {
	itemGroup := router.Group("/items")
	{
		itemGroup.GET("/", GetItems)
	}
}

func GetItems(c *gin.Context) {
	// Handler logic here
}
