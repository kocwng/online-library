package delivery

import "online-library/features/books"

type BookResponse struct {
	Id            uint   `json:"id"`
	Title         string `json:"title"`
	PublishedYear int    `json:"published_year"`
	ISBN          string `json:"isbn"`
}

func ToResponse(bookEntity books.BookEntity) BookResponse {
	return BookResponse{
		Id:            bookEntity.Id,
		Title:         bookEntity.Title,
		PublishedYear: bookEntity.PublishedYear,
		ISBN:          bookEntity.ISBN,
	}
}

func ToListResponse(bookEntity []books.BookEntity) []BookResponse {
	var response []BookResponse
	for _, v := range bookEntity {
		response = append(response, ToResponse(v))
	}
	return response
}
