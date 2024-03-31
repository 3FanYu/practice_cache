package emails

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewUsecase(dao EmailDAO) EmailUsecase {
	return &impl{dao: dao}
}

type impl struct {
	dao EmailDAO
}

func (im *impl) GetEmails(c *gin.Context, targetAddress string) {
	emails, err := im.dao.GetEmailByTargetAddress(targetAddress)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"emails": emails})
}
