package items

import (
	"net/http"

	business "github.com/3fanyu/glossika/internal/business/items"
	"github.com/3fanyu/glossika/internal/routers/middleware"
	"github.com/3fanyu/glossika/pkg/cachekit"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(router *gin.Engine, db *gorm.DB, cache cachekit.Cache) {
	itemGroup := router.Group("/v1/items")
	itemGroup.Use(middleware.AuthMiddleware())
	{
		itemGroup.GET("/recommendations", GetItems(db, cache))
	}
}

func GetItems(db *gorm.DB, cache cachekit.Cache) gin.HandlerFunc {
	return func(c *gin.Context) {
		var opt business.ListOptions
		if err := c.ShouldBindQuery(&opt); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		business.List(c, db, opt, cache)
	}
}
