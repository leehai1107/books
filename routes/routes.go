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
		bookGroup.GET("/:id", bookHandler.GetBookById)
		bookGroup.GET("", bookHandler.GetBooks)
		bookGroup.POST("/", bookHandler.CreateBook)
		bookGroup.PUT("/:id", bookHandler.UpdateBook)
		bookGroup.DELETE("/:id", bookHandler.DeleteBook)
		bookGroup.PATCH("/price", bookHandler.UpdateBookPrice)
	}

	return r
}
