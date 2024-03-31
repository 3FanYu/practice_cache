package items

import (
	"fmt"
	"net/http"
	"time"

	"github.com/3fanyu/glossika/internal/dao"
	"github.com/3fanyu/glossika/internal/models"
	"github.com/3fanyu/glossika/pkg/cachekit"
	"github.com/gin-gonic/gin"
)

const (
	cachePrefix   = "items"
	cacheDuration = 10 * time.Minute
)

func NewUsecase(cache cachekit.Cache, dao ItemDAO) ItemUsecase {
	return &impl{cache: cache, dao: dao}
}

type impl struct {
	cache cachekit.Cache
	dao   ItemDAO
}

func (im *impl) GetItems(c *gin.Context, opt dao.ListOptions) {
	var items []models.Item
	key := fmt.Sprintf("%v-%v", opt.Page, opt.Size)
	if err := im.cache.GetByFunc(c, cachePrefix, key, &items, func() (interface{}, error) {
		return im.dao.GetItems(opt)
	}); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch items"})
	}
	c.JSON(http.StatusOK, items)
}
