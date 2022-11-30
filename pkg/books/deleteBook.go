package books

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/untletch/gin-books/pkg/common/models"
)

func (d dbStore) DeleteBook(c *gin.Context) {
	var id = c.Param("id")
	var book models.Book

	var result = d.DB.First(&book, id)
	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": result.Error,
		})
		return
	}

	d.DB.Delete(&book)
	c.Status(http.StatusOK)
}
