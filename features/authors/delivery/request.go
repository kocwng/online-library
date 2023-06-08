package delivery

import "online-library/features/authors"

type AuthorRequest struct {
	Name    string
	Country string
}

func ToEntity(authorRequest AuthorRequest) authors.AuthorEntity {
	return authors.AuthorEntity{
		Name:    authorRequest.Name,
		Country: authorRequest.Country,
	}
}