package books

type UpdateBookRequest struct {
  Title string `json:"title"`
  Author string `json:"author"`
  Description string `json:"description"`
}

func (d dbStore) UpdateBook(c *gin.Context) {
  id := c.Param("id")
  var input UpdateBookRequest

  if err := c.BindJSON(&input); err != nil {
    c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
      "error": err.Error(),
    })
    return
  }

  var result = h.DB.First(&book, id)
  if result.Error != nil {
    c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
      "error": result.Error,
    })
  }

  book := models.Book{
    Title: input.Title,
    Author: input.Author,
    Description: input.Descripton,
  }

  h.DB.Save(&book)
  c.JSON(http.StatusOK, &book)
}
