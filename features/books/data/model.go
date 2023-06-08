package data

import (
	"online-library/features/authors/data"
	"online-library/features/books"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title         string
	PublishedYear int
	ISBN          string
	Authors []data.Author `gorm:"many2many:book_authors;"`
}

func FromEntity(bookEntity books.BookEntity) Book {
	return Book{
		Title:         bookEntity.Title,
		PublishedYear: bookEntity.PublishedYear,
		ISBN:          bookEntity.ISBN,
	}
}

func ToEntity(book Book) books.BookEntity {
	return books.BookEntity{
		Id:            book.ID,
		Title:         book.Title,
		PublishedYear: book.PublishedYear,
		ISBN:          book.ISBN,
	}
}

func ListToEntity(book []Book) []books.BookEntity {
	var bookEntity []books.BookEntity
	for _, v := range book {
		bookEntity = append(bookEntity, ToEntity(v))
	}
	return bookEntity
}