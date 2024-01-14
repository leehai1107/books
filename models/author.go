package models

type Author struct {
	AuthorID int    `json:"author_id" gorm:"column:author_id"`
	Name     string `json:"name" gorm:"column:name"`
}

type AuthorDTO struct {
	AuthorID int    `json:"-" gorm:"column:author_id"`
	Name     string `json:"name" gorm:"column:name"`
}
