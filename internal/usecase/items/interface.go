package items

import (
	"github.com/3fanyu/glossika/internal/dao"
	"github.com/3fanyu/glossika/internal/models"
	"github.com/gin-gonic/gin"
)

type ItemDAO interface {
	GetItems(filter dao.ListOptions) ([]models.Item, error)
}

type ItemUsecase interface {
	GetItems(c *gin.Context, opt dao.ListOptions)
}
