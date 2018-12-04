package bukaresep

// Repository is an interface as data access to a persistent storage
type Repository interface {
	Get(id int) (*Recipe, error)
	GetAll() ([]*Recipe, error)
	Add(recipe *Recipe) (*Recipe, error)
	Update(recipe *Recipe) error
	Delete(recipe *Recipe) error
}
