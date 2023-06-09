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

func (q *bookQuery) Edit(update books.BookEntity, id uint) error {
	book := FromEntity(update)
	if err := q.db.Where("id", id).Updates(&book); err.Error != nil {
		return err.Error
	}
	return nil
}

func (q bookQuery) Delete(id uint) error {
	var book Book
	if err := q.db.Delete(&book, id); err.Error != nil {
		return err.Error
	}
	return nil
}

// func (q *bookQuery) GetAll(authorId uint) ([]books.BookEntity, error) {
// 	var book []Book
// 	if err := q.db.Where("user_id = ?", authorId).Select("id,title").Order("title").Find(&book); err.Error != nil {
// 		return nil, err.Error
// 	}
// 	return ToListEntity(book), nil
// }

func (q *bookQuery) AddAuthorAssociation(bookID uint, authorID uint) error {
	association := map[string]interface{}{
		"book_id":   bookID,
		"author_id": authorID,
	}

	return q.db.Table("book_authors").Create(&association).Error
}


func (q *bookQuery) GetBooksByAuthor(authorID uint) ([]books.BookEntity, error) {
	var books []Book
	err := q.db.Table("books").
		Select("books.id, books.title, books.published_year, books.isbn").
		Joins("JOIN book_authors ON book_authors.book_id = books.id").
		Where("book_authors.author_id = ? AND books.deleted_at IS NULL", authorID).
		Find(&books).
		Error

	if err != nil {
		return nil, err
	}

	return ToListEntity(books), nil
}


