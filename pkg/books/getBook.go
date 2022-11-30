package books

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/untletch/gin-books/pkg/common/models"
)

func (d dbStore) GetBook(c *gin.Context) {
	id := c.Param("id")
	var book models.Book
	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": result.Error,
		})
	}
	c.JSON(http.StatusOK, &book)
}
