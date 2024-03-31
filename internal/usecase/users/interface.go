package users

import (
	"github.com/3fanyu/glossika/internal/dao"
	"github.com/3fanyu/glossika/internal/models"
	"github.com/gin-gonic/gin"
)

type UserDAO interface {
	GetUserByEmail(email string) (*models.User, error)
	CreateUser(user *models.User) error
}

type UserUsecase interface {
	Auth(c *gin.Context, input dao.AuthInput)
	CreateUser(c *gin.Context, input dao.RegisterInput)
}

type EmailDAO interface {
	CreateEmail(email *models.Email) error
}
