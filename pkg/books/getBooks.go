package books

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/untletch/gin-books/pkg/common/models"
)

func (d dbStore) GetBooks(c *gin.Context) {
	var books []models.Book
	var result = d.DB.Find(&books)
	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": result.Error,
		})
		return
	}
	log.Println("found all books")
	c.JSON(http.StatusOK, &books)
}
