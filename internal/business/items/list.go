package items

import (
	"fmt"
	"net/http"
	"time"

	"github.com/3fanyu/glossika/internal/dao"
	"github.com/3fanyu/glossika/internal/models"
	"github.com/3fanyu/glossika/pkg/cachekit"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

const (
	cachePrefix   = "items"
	cacheDuration = 10 * time.Minute
)

type ListOptions struct {
	Page int `form:"page"`
	Size int `form:"size"`
}

func List(c *gin.Context, db *gorm.DB, opt ListOptions, cache cachekit.Cache) {
	var items []models.Item
	key := fmt.Sprintf("%v-%v", opt.Page, opt.Size)
	if err := cache.GetByFunc(c, cachePrefix, key, &items, func() (interface{}, error) {
		itemDAO := dao.NewItemDAO(db)
		return itemDAO.GetItems(dao.ItemFilter{})	
	}); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch items"})	
	}
	c.JSON(http.StatusOK, items)
}
