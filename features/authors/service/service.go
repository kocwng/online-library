package service

import (
	"online-library/features/authors"

	"github.com/go-playground/validator/v10"
)

type authorService struct {
	Data     authors.AuthorDataInterface
	validate *validator.Validate
}

func New(data authors.AuthorDataInterface) authors.AuthorServiceInterface {
	return &authorService{
		Data:     data,
		validate: validator.New(),
	}
}

func (s *authorService) Create(new authors.AuthorEntity) (authors.AuthorEntity, error) {
	s.validate = validator.New()
	errValidate := s.validate.StructExcept(new, "Author")
	if errValidate != nil {
		return authors.AuthorEntity{}, errValidate
	}

	authorID, err := s.Data.Store(new)
	if err != nil {
		return authors.AuthorEntity{}, err
	}

	return s.Data.SelectById(authorID)
}

func (s *authorService) Update(feedbackEntity authors.AuthorEntity, id uint) (authors.AuthorEntity, error) {
	//check exist data
	if checkDataExist, err := s.Data.SelectById(id); err != nil {
		return checkDataExist, err
	}

	//update
	err := s.Data.Edit(feedbackEntity, id)
	if err != nil {
		return authors.AuthorEntity{}, err
	}
	return s.Data.SelectById(id)
}

// func (s *authorService) AddAuthorAssociation(authorID uint, authorID uint) error {
// 	return s.Data.AddAuthorAssociation(authorID, authorID)
// }

func (s *authorService) GetById(id uint) (authors.AuthorEntity, error) {
	return s.Data.SelectById(id)
}

func (s *authorService) Delete(id uint) error {
	_, err := s.Data.SelectById(id)
	if err != nil {
		return nil
	}
	return s.Data.Delete(id)
}

func (s *authorService) GetAuthorsByBook(bookID uint) ([]authors.AuthorEntity, error) {
	books, err := s.Data.GetAuthorsByBook(bookID)
	if err != nil {
		return nil, err
	}

	return books, nil
}
