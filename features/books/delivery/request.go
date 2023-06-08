package delivery

import (
	"online-library/features/authors/delivery"
	"online-library/features/books"
)

type BookRequest struct {
	Title         string
	PublishedYear int
	ISBN          string
	Author        []delivery.AuthorRequest
}

func ToEntity(bookRequest BookRequest) books.BookEntity {
	result := books.BookEntity{
		Title: bookRequest.Title,
		PublishedYear: bookRequest.PublishedYear,
		ISBN: bookRequest.ISBN,
	}

	for _, v := range bookRequest.Author{
		result.Author = append(result.Author, delivery.ToEntity(v))
	}
	return result
}