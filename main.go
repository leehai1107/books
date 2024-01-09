package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type book struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
	Status   bool   `json:"status"`
}

type bookDTO struct {
	Title    string `json:"title" gorm:"column:title;"`
	Author   string `json:"author gorm:"column:author;"`
	Quantity int    `json:"quantity gorm:"column:quantity;"`
	Status   bool   `json:"status gorm:"column:status;"`
}

var books = []book{
	{ID: 1, Title: "In Search of Lost Time", Author: "Marcel Proust", Quantity: 2, Status: false},
	{ID: 2, Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Quantity: 5, Status: true},
	{ID: 3, Title: "War and Peace", Author: "Leo Tolstoy", Quantity: 6, Status: true},
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func (bookDTO) TableName() string {
	return "books"
}

func main() {

	dsn := os.Getenv("DB_CONN_STR")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	//CRUD: Create, Read, update, Delete
	//	POST / v 1/ items (create a new item)
	//	C,€r ,/vl/items (12 st Items) / v 1/2 tems *page-I
	//	C,€r /vl/items/:åd (get item detan by ld)
	//	(PUT // PATCH) / v 1/2 tems/:id (update an Item by ld)
	//	DELETE (delete item by ld)

	fmt.Println(db)

	router := gin.Default()
	v1 := router.Group("/v1")
	{
		items := v1.Group("/books")
		{
			items.POST("")
			items.GET("/:id")
			items.GET("", getBooks)
			items.PATCH("/:id")
			items.DELETE("/:id")
		}
	}

	//router.GET("/books", getBooks)
	router.Run("localhost:8080")
}

//func createBook() func(*gin.Context) {
//	return func(c *gin.Context) {
//		var data bookDTO
//		c.IndentedJSON(http.StatusOK, data)
//	}
//}

func createBook(c *gin.Context) {
	var data bookDTO
	if err := c.ShouldBind(&data); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "add book to database"})
}
