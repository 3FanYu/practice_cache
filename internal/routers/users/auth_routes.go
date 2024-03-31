package users

import (
	"net/http"

	business "github.com/3fanyu/glossika/internal/business/users"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(router *gin.Engine, db *gorm.DB) {
	userGroup := router.Group("/v1/user")
	{
		userGroup.POST("/register", CreateUser(db))
		userGroup.POST("/sign_in", AuthenticateUser(db))
	}
}

func CreateUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input business.RegisterInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		business.CreateUser(c, db, input)
	}
}

func AuthenticateUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input business.AuthInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		business.Auth(c, db, input)
	}
}
