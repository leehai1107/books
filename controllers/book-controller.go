package controllers

import (
	"example/Demo/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

// BookRoutes sets up routes for the books controller
func BookRoutes(router *gin.Engine, db *gorm.DB) {
	items := router.Group("/books")
	{
		items.POST("", createBook(db))
		items.GET("/:id", getBook(db))
		items.GET("", getBooks(db))
		items.PATCH("/:id", updateBook(db))
		items.DELETE("/:id", deleteBook(db))
	}
}

func createBook(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data models.BookCreation
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

func getBook(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data models.Book

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

func updateBook(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		// Start a GORM transaction.
		tx := db.Begin()

		// Ensure the transaction is rolled back if there is an error or it is not committed.
		defer func() {
			if r := recover(); r != nil {
				tx.Rollback()
			}
		}()

		var data models.BookUpdate

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

func getBooks(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data []models.Book

		if err := db.Find(&data).Error; err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.IndentedJSON(http.StatusOK, gin.H{"data": data})

	}
}

func deleteBook(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		//books/:id
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		//data.ID = id
		if err := db.Table(models.Book{}.TableName()).Where("id = ?", id).Delete(nil).Error; err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.IndentedJSON(http.StatusOK, gin.H{"message": "delete successfully!"})

	}
}
