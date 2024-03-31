package users

import (
	"net/http"

	"github.com/3fanyu/glossika/internal/dao"
	uc "github.com/3fanyu/glossika/internal/usecase/users"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, uc uc.UserUsecase) {
	userGroup := router.Group("/v1/user")
	{
		userGroup.POST("/register", CreateUser(uc))
		userGroup.POST("/sign_in", AuthenticateUser(uc))
	}
}

func CreateUser(uc uc.UserUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input dao.RegisterInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		uc.CreateUser(c, input)
	}
}

func AuthenticateUser(uc uc.UserUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input dao.AuthInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		uc.Auth(c, input)
	}
}
