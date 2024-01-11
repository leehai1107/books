package repository

import (
	"example/Demo/models"
	"gorm.io/gorm"
)

type IBookRepo interface {
	GetBookById(id int) (models.Book, error)
	GetBooks() ([]models.Book, error)
	CreateBook(creation *models.BookCreation) (int, error)
	UpdateBook(update *models.BookUpdate, bookId int) (int, error)
	DeleteBook(id int) (string, error)
}

type BookRepository struct {
	Db *gorm.DB
}

func (r BookRepository) GetBookById(id int) (models.Book, error) {
	var data models.Book
	if err := r.Db.Where("id = ?", id).First(&data).Error; err != nil {
		return models.Book{}, err
	}

	return data, nil
}

func (r BookRepository) GetBooks() ([]models.Book, error) {
	var data []models.Book

	if err := r.Db.Find(&data).Error; err != nil {
		return nil, err
	}

	return data, nil
}

func (r BookRepository) CreateBook(creation *models.BookCreation) (int, error) {

	if err := r.Db.Create(&creation).Error; err != nil {
		return -1, err
	}

	return creation.ID, nil
}

func (r BookRepository) UpdateBook(update *models.BookUpdate, bookId int) (int, error) {
	if err := r.Db.Model(&models.BookUpdate{}).Where("id = ?", bookId).Updates(&update).Error; err != nil {
		return -1, err
	}
	return bookId, nil
}

func (r BookRepository) DeleteBook(id int) (string, error) {
	if err := r.Db.Table(models.Book{}.TableName()).Where("id = ?", id).Delete(nil).Error; err != nil {
		return "Error", err
	}
	return "Book Deleted!", nil
}
