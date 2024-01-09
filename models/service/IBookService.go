package service

import "example/Demo/models/entity"

type IBookServce interface {
	getBooks() []entity.Book
	getBook() entity.Book
	updateBook() entity.Book
}
