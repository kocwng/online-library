package delivery

import (
	"online-library/features/authors/delivery"
	"online-library/features/books"
)

type BookRequest struct {
	Title         string                   `json:"title" form:"title"`
	PublishedYear int                      `json:"published_year" form:"published_year"`
	ISBN          string                   `json:"isbn" form:"isbn"`
	Author        []delivery.AuthorRequest `json:"author" form:"author"`
	AuthorId      []uint                    `json:"author_id" form:"author_id"`
}

type BookAuthorRequest struct {
	BookID uint `json:"book_id"`
	AuthorID uint `json:"author_id"`

}

func ToEntity(bookRequest BookRequest) books.BookEntity {
	result := books.BookEntity{
		Title:         bookRequest.Title,
		PublishedYear: bookRequest.PublishedYear,
		ISBN:          bookRequest.ISBN,
		AuthorId: bookRequest.AuthorId,
	}

	// for _, v := range bookRequest.AuthorId {
	// 	result.AuthorId = append(result.AuthorId, delivery.ToEntity(v))
	// }

	for _, v := range bookRequest.Author {
		result.Author = append(result.Author, delivery.ToEntity(v))
	}
	return result
}
