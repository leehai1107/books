package service

import (
	"example/Demo/models"
	"example/Demo/repository"
)

type IBookService interface {
	GetBookById(id int) (models.Book, error)
	GetBooks() ([]models.Book, error)
	CreateBook(creation *models.BookCreation) (int, error)
	UpdateBook(update models.BookUpdate) (int, error)
	DeleteBook(id int) (string, error)
}

type BookService struct {
	IBookRepo repository.IBookRepo
}

func (s BookService) GetBookById(id int) (models.Book, error) {
	//TODO implement me
	panic("implement me")
}

func (s BookService) GetBookByID(id int) (models.Book, error) {
	data, err := s.IBookRepo.GetBookById(id)
	return data, err
}

func (s BookService) GetBooks() ([]models.Book, error) {
	data, err := s.IBookRepo.GetBooks()
	return data, err
}

func (s BookService) CreateBook(creation *models.BookCreation) (int, error) {
	data, err := s.IBookRepo.CreateBook(creation)
	return data, err
}

func (s BookService) UpdateBook(update models.BookUpdate) (int, error) {
	data, err := s.IBookRepo.UpdateBook(update)
	return data, err
}

func (s BookService) DeleteBook(id int) (string, error) {
	data, err := s.IBookRepo.DeleteBook(id)
	return data, err
}
