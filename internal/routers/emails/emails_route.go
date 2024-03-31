package emails

import (
	uc "github.com/3fanyu/glossika/internal/usecase/emails"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, uc uc.EmailUsecase) {
	itemGroup := router.Group("/v1/emails")
	{
		itemGroup.GET("/", GetItems(uc))
	}
}

func GetItems(uc uc.EmailUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		targetAddress := c.Query("email")
		uc.GetEmails(c, targetAddress)
	}
}
