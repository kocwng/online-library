package data

import (
	"online-library/features/authors"

	"gorm.io/gorm"
)

type Author struct {
	gorm.Model
	Name    string
	Country string
}

func FromEntity(authorEntity authors.AuthorEntity) Author {
	return Author{
		Name: authorEntity.Name,
		Country: authorEntity.Country,
		
	}
}

func ToEntity(author Author) authors.AuthorEntity {
	return authors.AuthorEntity{
		Id:            author.ID,
		Name: author.Name,
		Country: author.Country,

	}
}

func ToListEntity(author []Author) []authors.AuthorEntity {
	var authorEntity []authors.AuthorEntity
	for _, v := range author {
		authorEntity = append(authorEntity, ToEntity(v))
	}
	return authorEntity
}