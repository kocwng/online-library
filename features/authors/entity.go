package authors

type AuthorEntity struct {
	Id uint
	Name string 
	Country string
	AuthorId uint
}

type AuthorServiceInterface interface {
	Create(new AuthorEntity) (AuthorEntity, error)
	// AddAuthorAssociation(authorID uint, authorID uint) error
	GetById(id uint) (AuthorEntity, error)
	Update(update AuthorEntity, id uint) (AuthorEntity, error)
	Delete(id uint) error
	// GetAll() ([]AuthorEntity, error)


}

type AuthorDataInterface interface {
	Store(new AuthorEntity) (uint, error)
	SelectById(id uint) (AuthorEntity, error)
	// AddAuthorAssociation(authorID uint, authorID uint) error
	Edit(update AuthorEntity, id uint) error
	Delete(id uint) error
	// GetAll() ([]AuthorEntity, error)


}
