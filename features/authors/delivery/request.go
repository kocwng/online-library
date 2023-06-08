package delivery

import "online-library/features/authors"

type AuthorRequest struct {
	Name    string `json:"name" form:"name"`
	Country string `json:"country" form:"country"`
}

func ToEntity(authorRequest AuthorRequest) authors.AuthorEntity {
	return authors.AuthorEntity{
		Name:    authorRequest.Name,
		Country: authorRequest.Country,
	}
}