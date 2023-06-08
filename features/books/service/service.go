package service

import (
	"online-library/features/books"

	"github.com/go-playground/validator"
)

type bookService struct {
	Data     books.BookDataInterface
	validate *validator.Validate
}

func New(data books.BookDataInterface) books.BookServiceInterface {
	return &bookService{
		Data:     data,
		validate: validator.New(),
	}
}

func (s *bookService) Create(new books.BookEntity) (books.BookEntity, error) {
	s.validate = validator.New()
	errValidate := s.validate.StructExcept(new, "Author")
	if errValidate != nil {
		return books.BookEntity{}, errValidate
	}
	bookId, err := s.Data.Store(new)
	if err != nil {
		return books.BookEntity{}, err
	}
	return s.Data.SelectById(bookId)
}