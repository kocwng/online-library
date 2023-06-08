package service

import (
	"online-library/features/books"

	"github.com/go-playground/validator"
)

type Service struct {
	Data     books.BookDataInterface
	validate *validator.Validate
}

func New(data books.BookDataInterface) books.BookServiceInterface {
	return &Service{
		Data:     data,
		validate: validator.New(),
	}
}

