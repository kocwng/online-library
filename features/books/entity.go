package books

import "online-library/features/authors"

type BookEntity struct {
	Id            uint
	Title         string
	PublishedYear int
	ISBN          string
	Author        []authors.AuthorEntity
	AuthorId      []uint
}

type BookServiceInterface interface {
	Create(new BookEntity) (BookEntity, error)
	AddAuthorAssociation(bookID uint, authorID uint) error
	GetById(id uint) (BookEntity, error)
	Update(update BookEntity, id uint) (BookEntity, error)
	Delete(id uint) error
	GetBooksByAuthor(authorID uint) ([]BookEntity, error)
	// GetAll() ([]BookEntity, error)


}

type BookDataInterface interface {
	Store(new BookEntity) (uint, error)
	SelectById(id uint) (BookEntity, error)
	AddAuthorAssociation(bookID uint, authorID uint) error
	Edit(update BookEntity, id uint) error
	Delete(id uint) error
	GetBooksByAuthor(authorID uint) ([]BookEntity, error)

	// GetAll() ([]BookEntity, error)


}
