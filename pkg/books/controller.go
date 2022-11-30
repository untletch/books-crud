package books

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type crud interface {
	GetBook(c *gin.Context)
	GetBooks(c *gin.Context)
	AddBook(c *gin.Context)
	UpdateBook(c *gin.Context)
	DeleteBook(c *gin.Context)
}

type dbStore struct {
	DB *gorm.DB
}

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	d := &dbStore{DB: db}

	routes := r.Group("v1/books")
	routes.GET("/", d.GetBooks)
	routes.GET("/:id", d.GetBook)
	routes.POST("/", d.AddBook)
	routes.PUT("/:id", d.UpdateBook)
	routes.DELETE("/:id", d.DeleteBook)
}
