package models

type Book struct {
	ID       int     `json:"id" gorm:"column:id"`
	Title    string  `json:"title" gorm:"column:title;"`
	AuthorId int     `json:"author_id" gorm:"column:author_id;"`
	Quantity int     `json:"quantity" gorm:"column:quantity;"`
	Price    float32 `json:"price" gorm:"column:price"`
	Status   bool    `json:"status" gorm:"column:status;"`
}

func (Book) TableName() string {
	return "books"
}

type BookCreation struct {
	ID       int     `json:"-" gorm:"column:id"`
	Title    string  `json:"title" gorm:"column:title;"`
	AuthorId int     `json:"author_id" gorm:"column:author_id;"`
	Quantity int     `json:"quantity" gorm:"column:quantity;"`
	Price    float32 `json:"price" gorm:"column:price"`
	Status   bool    `json:"status" gorm:"column:status;"`
}

func (BookCreation) TableName() string {
	return Book{}.TableName()
}

type BookUpdate struct {
	Title    string  `json:"title" gorm:"column:title;"`
	AuthorId int     `json:"author_id" gorm:"column:author_id;"`
	Quantity int     `json:"quantity" gorm:"column:quantity;"`
	Price    float32 `json:"price" gorm:"column:price"`
	Status   bool    `json:"status" gorm:"column:status;"`
}

func (BookUpdate) TableName() string {
	return Book{}.TableName()
}

type BookUpdatePrice struct {
	BookIds []int   `json:"book_ids" binding:"required"`
	Price   float32 `json:"price" binding:"required"`
}

func (BookUpdatePrice) TableName() string {
	return Book{}.TableName()
}
