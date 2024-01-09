package dto

type AuthorDTO struct {
	AuthorID int    `json:"-" gorm:"column:author_id"`
	Name     string `json:"name" gorm:"column:name"`
}
