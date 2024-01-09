package dto

type BookDTO struct {
	ID       int    `json:"-" gorm:"column:id"`
	Title    string `json:"title" gorm:"column:title;"`
	AuthorId int    `json:"author_id" gorm:"column:author_id;"`
	Quantity int    `json:"quantity" gorm:"column:quantity;"`
	Status   bool   `json:"status" gorm:"column:status;"`
}
