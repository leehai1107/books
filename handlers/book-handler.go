package handlers

import (
	"example/Demo/models"
	"example/Demo/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type BookHandler struct {
	IBookService service.IBookService
}

func (s *BookHandler) GetBookById(c *gin.Context) {
	BookId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	book, err := s.IBookService.GetBookById(BookId)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"data": book})
}

func (s *BookHandler) GetBooks(c *gin.Context) {
	data, err := s.IBookService.GetBooks(c)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"data": data})
}

func (s *BookHandler) CreateBook(c *gin.Context) {
	var data models.BookCreation
	if err := c.ShouldBind(&data); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := s.IBookService.CreateBook(&data, c)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"bookId": result})
}

func (s *BookHandler) UpdateBook(c *gin.Context) {
	BookId, err := strconv.Atoi(c.Param("id"))

	var data models.BookUpdate
	if err := c.ShouldBind(&data); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := s.IBookService.UpdateBook(&data, BookId, c)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"bookId": result})
}

func (s *BookHandler) DeleteBook(c *gin.Context) {
	BookId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	book, err := s.IBookService.DeleteBook(BookId, c)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"data": book})
}
