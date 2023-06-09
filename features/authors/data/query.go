package data

import (
	"online-library/features/authors"

	"gorm.io/gorm"
)

type authorQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) authors.AuthorDataInterface {
	return &authorQuery{
		db: db,
	}
}

func (q *authorQuery) Store(new authors.AuthorEntity) (uint, error) {
	author := FromEntity(new)
	if err := q.db.Create(&author); err.Error != nil {
		return 0, err.Error
	}
	return author.ID, nil
}

func (q *authorQuery) SelectById(id uint) (authors.AuthorEntity, error) {
	var author Author
	if err := q.db.First(&author, id); err.Error != nil {
		return authors.AuthorEntity{}, err.Error
	}

	return ToEntity(author), nil
}

func (q *authorQuery) Edit(update authors.AuthorEntity, id uint) error {
	author := FromEntity(update)
	if err := q.db.Where("id", id).Updates(&author); err.Error != nil {
		return err.Error
	}
	return nil
}

func (q *authorQuery) Delete(id uint) error {
	var author Author
	if err := q.db.Delete(&author, id); err.Error != nil {
		return err.Error
	}
	return nil
}

func (q *authorQuery) GetAuthorsByBook(bookID uint) ([]authors.AuthorEntity, error) {
	var author []Author
	err := q.db.Table("books").
		Select("authors.id, authors.name, authors.country").
		Joins("JOIN book_authors ON book_authors.author_id = authors.id").
		Where("book_authors.book_id = ? AND authors.deleted_at IS NULL", bookID).
		Find(&author).
		Error

	if err != nil {
		return nil, err
	}

	return ToListEntity(author), nil
}


