package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Book struct {
	ID       int    `json:"id" gorm:"column:id"`
	Title    string `json:"title" gorm:"column:title;"`
	AuthorId int    `json:"author_id" gorm:"column:author_id;"`
	Quantity int    `json:"quantity" gorm:"column:quantity;"`
	Status   bool   `json:"status" gorm:"column:status;"`
}

func (Book) TableName() string {
	return "books"
}

type Author struct {
	AuthorID int    `json:"author_id" gorm:"column:author_id"`
	Name     string `json:"name" gorm:"column:name"`
}

type AuthorDTO struct {
	AuthorID int    `json:"-" gorm:"column:author_id"`
	Name     string `json:"name" gorm:"column:name"`
}

type BookCreation struct {
	ID       int    `json:"-" gorm:"column:id"`
	Title    string `json:"title" gorm:"column:title;"`
	AuthorId int    `json:"author_id" gorm:"column:author_id;"`
	Quantity int    `json:"quantity" gorm:"column:quantity;"`
	Status   bool   `json:"status" gorm:"column:status;"`
}

func (BookCreation) TableName() string {
	return Book{}.TableName()
}

type BookUpdate struct {
	Title    string `json:"title" gorm:"column:title;"`
	AuthorId int    `json:"author_id" gorm:"column:author_id;"`
	Quantity int    `json:"quantity" gorm:"column:quantity;"`
	Status   bool   `json:"status" gorm:"column:status;"`
}

func (BookUpdate) TableName() string {
	return Book{}.TableName()
}

//func getBooks(c *gin.Context) {
//	c.IndentedJSON(http.StatusOK, books)
//}

func main() {

	dsn := os.Getenv("DB_CONN_STR")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(db)

	router := gin.Default()
	v1 := router.Group("/v1")
	{
		items := v1.Group("/books")
		{
			items.POST("", createBook(db))
			items.GET("/:id", getBook(db))
			items.GET("", getBooks(db))
			items.PATCH("/:id", updateBook(db))
			items.DELETE("/:id")
		}
	}

	//router.GET("/books", getBooks)
	router.Run("localhost:8080")
}

func createBook(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data BookCreation
		if err := c.ShouldBind(&data); err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := db.Create(&data).Error; err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.IndentedJSON(http.StatusOK, gin.H{"bookId": data.ID})
	}
}

func getBook(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var data Book

		//v1/books/:id
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		//data.ID = id
		if err := db.Where("id = ?", id).First(&data).Error; err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.IndentedJSON(http.StatusOK, gin.H{"data": data})

	}
}

func updateBook(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {

		// Start a GORM transaction.
		tx := db.Begin()

		// Ensure the transaction is rolled back if there is an error or it is not committed.
		defer func() {
			if r := recover(); r != nil {
				tx.Rollback()
			}
		}()

		var data BookUpdate

		//v1/books/:id
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			tx.Rollback()
			return
		}

		if err := c.ShouldBind(&data); err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			tx.Rollback()
			return
		}

		// Wrap the update operation in a transaction
		err = tx.Transaction(func(tx *gorm.DB) error {
			if err := tx.Where("id = ?", id).Updates(&data).Error; err != nil {
				return err
			}
			return nil
		})

		// Commit the transaction on successful update.
		if err := tx.Commit().Error; err != nil {
			// Handle the case where the transaction commit fails.
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "failed to commit transaction"})
			tx.Rollback()
			return
		}

		c.IndentedJSON(http.StatusOK, gin.H{"message": "update successful"})

	}
}

func getBooks(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var data []Book

		if err := db.Find(&data).Error; err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.IndentedJSON(http.StatusOK, gin.H{"data": data})

	}
}
