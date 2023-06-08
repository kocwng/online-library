package data

import (

	"online-library/features/authors/data"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title         string
	PublishedYear int
	ISBN          string
	Authors []data.Author `gorm:"many2many:book_authors;"`
}

