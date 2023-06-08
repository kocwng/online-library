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

// func (s *bookService) Create(new books.BookEntity) (books.BookEntity, error) {
// 	s.validate = validator.New()
// 	errValidate := s.validate.StructExcept(new, "Author")
// 	if errValidate != nil {
// 		return books.BookEntity{}, errValidate
// 	}
// 	bookId, err := s.Data.Store(new)
// 	if err != nil {
// 		return books.BookEntity{}, err
// 	}
// 	return s.Data.SelectById(bookId)
// }


func (s *bookService) Create(new books.BookEntity) (books.BookEntity, error) {
	s.validate = validator.New()
	errValidate := s.validate.StructExcept(new, "Author")
	if errValidate != nil {
		return books.BookEntity{}, errValidate
	}

	bookID, err := s.Data.Store(new)
	if err != nil {
		return books.BookEntity{}, err
	}

	return s.Data.SelectById(bookID)
}

func (s *bookService) Update(feedbackEntity books.BookEntity, id uint) (books.BookEntity, error) {
	//check exist data
	if checkDataExist, err := s.Data.SelectById(id); err != nil {
		return checkDataExist, err
	}

	//update
	err := s.Data.Edit(feedbackEntity, id)
	if err != nil {
		return books.BookEntity{}, err
	}
	return s.Data.SelectById(id)
}

func (s *bookService) AddAuthorAssociation(bookID uint, authorID uint) error {
	return s.Data.AddAuthorAssociation(bookID, authorID)
}

func (s *bookService) GetById(id uint) (books.BookEntity, error) {
	return s.Data.SelectById(id)
}