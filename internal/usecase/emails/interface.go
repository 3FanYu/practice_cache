package emails

import (
	"github.com/gin-gonic/gin"
	"github.com/3fanyu/glossika/internal/models"
)

type EmailUsecase interface {
	GetEmails(c *gin.Context, targetAddress string)
}

type EmailDAO interface {
	GetEmailByTargetAddress(targetAddress string) (*[]models.Email, error)
}
