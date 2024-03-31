package items

import (
	"net/http"

	"github.com/3fanyu/glossika/internal/dao"
	uc "github.com/3fanyu/glossika/internal/usecase/items"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, uc uc.ItemUsecase, middleware func() gin.HandlerFunc) {
	itemGroup := router.Group("/v1/items")
	itemGroup.Use(middleware())
	{
		itemGroup.GET("/recommendations", GetItems(uc))
	}
}

func GetItems(uc uc.ItemUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		var opt dao.ListOptions
		if err := c.ShouldBindQuery(&opt); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		uc.GetItems(c, opt)
	}
}
