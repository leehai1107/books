package models

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

// unsued
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
	ID       int    `json:"id" gorm:"column:id"`
	Title    string `json:"title" gorm:"column:title;"`
	AuthorId int    `json:"author_id" gorm:"column:author_id;"`
	Quantity int    `json:"quantity" gorm:"column:quantity;"`
	Status   bool   `json:"status" gorm:"column:status;"`
}

func (BookUpdate) TableName() string {
	return Book{}.TableName()
}
