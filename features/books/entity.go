package books

import "online-library/features/authors"

type BookEntity struct {
	Id            uint
	Title         string
	PublishedYear int
	ISBN          string
	Author        []authors.AuthorEntity
}

type BookServiceInterface interface {
	Create(new BookEntity) (BookEntity, error)
}

type BookDataInterface interface {
	Store(new BookEntity) (uint, error)
	SelectById(id uint) (BookEntity, error)
}
