package service

import (
	"context"
	"example/Demo/initilization"
	"example/Demo/models"
	"example/Demo/repository"
	"github.com/goccy/go-json"
)

type IBookService interface {
	GetBookById(id int) (models.Book, error)
	GetBooks() ([]models.Book, error)
	CreateBook(creation *models.BookCreation) (int, error)
	UpdateBook(update *models.BookUpdate) (int, error)
	DeleteBook(id int) (string, error)
}

type BookService struct {
	IBookRepo repository.IBookRepo
}

func (s BookService) GetBookById(id int) (models.Book, error) {
	data, err := s.IBookRepo.GetBookById(id)
	return data, err
}

func (s BookService) GetBooks() ([]models.Book, error) {
	// Check if data is cached
	cachedData, err := initilization.RedisClient.Get(context.Background(), "books").Result()
	if err == nil {
		// Use cached data if available
		var books []models.Book
		if err := json.Unmarshal([]byte(cachedData), &books); err == nil {
			return books, nil
		}
	}

	// Fetch data from the database if not in the cache
	data, err := s.IBookRepo.GetBooks()
	if err != nil {
		return nil, err
	}

	// Store data in the cache
	jsonData, err := json.Marshal(data)
	if err == nil {
		initilization.RedisClient.Set(context.Background(), "books", jsonData, 0)
	}

	return data, nil
}

func (s BookService) CreateBook(creation *models.BookCreation) (int, error) {
	data, err := s.IBookRepo.CreateBook(creation)
	return data, err
}

func (s BookService) UpdateBook(update *models.BookUpdate) (int, error) {
	data, err := s.IBookRepo.UpdateBook(update)
	return data, err
}

func (s BookService) DeleteBook(id int) (string, error) {
	data, err := s.IBookRepo.DeleteBook(id)
	return data, err
}
