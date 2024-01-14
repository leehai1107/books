package repository

import (
	"example/Demo/models"
	"gorm.io/gorm"
	"log"
	"sync"
)

type IBookRepo interface {
	GetBookById(id int) (models.Book, error)
	GetBooks(page models.Paging) ([]models.Book, error)
	CreateBook(creation *models.BookCreation) (int, error)
	UpdateBook(update *models.BookUpdate, bookId int) (int, error)
	UpdateBookPrice(bookId []int, price float32) ([]int, error)
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

func (r BookRepository) GetBooks(page models.Paging) ([]models.Book, error) {
	var data []models.Book

	if err := r.Db.Table(models.Book{}.TableName()).Count(&page.Total).Error; err != nil {
		log.Println(err)
	}

	if err := r.Db.
		Order("id asc").
		Offset((page.Page - 1) * page.Limit).
		Limit(page.Limit).
		Find(&data).Error; err != nil {
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
	err := r.Db.Transaction(func(tx *gorm.DB) error {
		err := r.Db.Model(&models.BookUpdate{}).Where("id = ?", bookId).Updates(&update).Error
		return err
	})
	if err != nil {
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
func (r BookRepository) UpdateBookPrice(bookIds []int, price float32) ([]int, error) {
	wg := &sync.WaitGroup{}
	bookChannel := make(chan int, 10)

	for _, book := range bookIds {
		wg.Add(1)
		go updatePrice(wg, book, bookChannel, price, r)
		bookChannel <- book
	}

	close(bookChannel)
	wg.Wait()
	log.Println("All books price are updated!")

	return bookIds, nil
}

func updatePrice(wg *sync.WaitGroup, bookId int, bookChannel chan int, price float32, r BookRepository) {
	defer wg.Done()

	select {
	case <-bookChannel:
		err := r.Db.Transaction(func(tx *gorm.DB) error {
			err := r.Db.Model(&models.BookUpdate{}).Where("id = ?", bookId).Update("price", price).Error
			return err
		})
		if err != nil {
			log.Println(err)
		}
		log.Printf("BookId %d has updated price!\n", bookId)
	}
}
