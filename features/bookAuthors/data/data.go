package data

import (
	bookauthors "online-library/features/bookAuthors"

	"gorm.io/gorm"
)

type BookAuthor struct {
	gorm.Model
	BookId   uint `json:"book_id"`
	AuthorId uint `json:"author_id"`
}

func FromEntity(bookAuthorEntity bookauthors.BookAuthorEntity) BookAuthor {
	return BookAuthor{
		BookId: bookAuthorEntity.BookId,
		AuthorId: bookAuthorEntity.AuthorId,
	}
}

func ToEntity(bookAuthor BookAuthor) bookauthors.BookAuthorEntity{
	return bookauthors.BookAuthorEntity{
		Id: bookAuthor.ID,
		BookId: bookAuthor.BookId,
		AuthorId: bookAuthor.AuthorId,
	}
}