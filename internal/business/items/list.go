package items

import (
	"net/http"

	"github.com/3fanyu/glossika/internal/dao"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func List(c *gin.Context, db *gorm.DB) {
	itemDAO := dao.NewItemDAO(db)
	items, err := itemDAO.GetItems(dao.ItemFilter{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch items"})
		return
	}
	c.JSON(http.StatusOK, items)
}
