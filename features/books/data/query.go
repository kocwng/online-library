package data

import (
	"online-library/features/books"

	"gorm.io/gorm"
)

type bookQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) books.BookDataInterface {
	return &bookQuery{
		db: db,
	}
}

func (q *bookQuery) Store(new books.BookEntity) (uint, error) {
	book := FromEntity(new)
	if err := q.db.Create(&book); err.Error != nil {
		return 0, err.Error
	}
	return book.ID, nil
}

func (q *bookQuery) SelectById(id uint) (books.BookEntity, error) {
	var book Book
	if err := q.db.First(&book, id); err.Error != nil {
		return books.BookEntity{}, err.Error
	}

	return ToEntity(book), nil
}