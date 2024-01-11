package service

import (
	"context"
	"example/Demo/initilization"
	"example/Demo/models"
	"example/Demo/repository"
	"github.com/goccy/go-json"
	"github.com/redis/go-redis/v9"
)

type IBookService interface {
	GetBookById(id int) (models.Book, error)
	GetBooks(ctx context.Context) ([]models.Book, error)
	CreateBook(creation *models.BookCreation, ctx context.Context) (int, error)
	UpdateBook(update *models.BookUpdate, bookId int, ctx context.Context) (int, error)
	DeleteBook(id int, ctx context.Context) (string, error)
}

type BookService struct {
	IBookRepo repository.IBookRepo
}

func (s BookService) GetBookById(id int) (models.Book, error) {
	data, err := s.IBookRepo.GetBookById(id)
	return data, err
}

func (s BookService) GetBooks(ctx context.Context) ([]models.Book, error) {
	// Check if data is cached
	cachedData, err := initilization.RedisClient.Get(ctx, "books").Result()
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
		initilization.RedisClient.Set(ctx, "books", jsonData, 0)
	}

	return data, nil
}

func (s BookService) CreateBook(creation *models.BookCreation, ctx context.Context) (int, error) {
	data, err := s.IBookRepo.CreateBook(creation)
	InvalidCacheKey("books", ctx, initilization.RedisClient)
	return data, err
}

func (s BookService) UpdateBook(update *models.BookUpdate, bookId int, ctx context.Context) (int, error) {
	data, err := s.IBookRepo.UpdateBook(update, bookId)
	InvalidCacheKey("books", ctx, initilization.RedisClient)
	return data, err
}

func (s BookService) DeleteBook(id int, ctx context.Context) (string, error) {
	data, err := s.IBookRepo.DeleteBook(id)
	InvalidCacheKey("books", ctx, initilization.RedisClient)
	return data, err
}

func InvalidCacheKey(key string, ctx context.Context, rdb *redis.Client) {
	err := rdb.Del(ctx, key).Err()
	if err != nil {
		panic(err)
	}
}
