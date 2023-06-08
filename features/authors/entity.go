package authors

type AuthorEntity struct {
	Id uint
	Name string 
	Country string
	BookId uint
}

type AuthorServiceInterface interface {
	Create(new AuthorEntity) (AuthorEntity, error)
	GetById(id uint) (AuthorEntity, error)

}

type AuthorDataInterface interface {
	Store(new AuthorEntity) (uint, error)
	SelectById(id uint) (AuthorEntity, error)

}