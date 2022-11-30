package books

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/untletch/gin-books/pkg/common/models"
)

type AddBookRequest struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
}

func (d dbStore) AddBook(c *gin.Context) {
	var input AddBookRequest

	if err := c.BindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	book := models.Book{
		Title:       input.Title,
		Author:      input.Author,
		Description: input.Description,
	}

	result := d.DB.Create(&book)
	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": result.Error,
		})
		return
	}

	c.JSON(http.StatusCreated, &book)
}
