package routes

import (
	"example/Demo/handlers"
	"example/Demo/repository"
	"example/Demo/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateRoutes(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// Routes for book
	bookHandler := handlers.BookHandler{
		IBookService: service.BookService{
			IBookRepo: &repository.BookRepository{
				Db: db,
			},
		},
	}

	bookGroup := r.Group("/books")
	{
		// Add your book routes here using Gin
		// For example:
		bookGroup.GET("/:id", bookHandler.GetBookById)
		bookGroup.GET("", bookHandler.GetBooks)
		bookGroup.POST("/", bookHandler.CreateBook)
		bookGroup.PUT("/:id", bookHandler.UpdateBook)
		bookGroup.DELETE("/:id", bookHandler.DeleteBook)
	}

	return r
}
