package delivery

import "online-library/features/authors"

type AuthorResponse struct {
	Id      uint   `json:"id"`
	Name    string `json:"name"`
	Country string `json:"country"`
}

func ToResponse(authorEntity authors.AuthorEntity) AuthorResponse {
	return AuthorResponse{
		Id:      authorEntity.Id,
		Name:    authorEntity.Name,
		Country: authorEntity.Country,
	}
}

func ToListResponse(authorEntity []authors.AuthorEntity) []AuthorResponse {
	var response []AuthorResponse
	for _, v := range authorEntity {
		response = append(response, ToResponse(v))
	}
	return response
}