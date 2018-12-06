package bukaresep

// Repository is an interface as data access to a persistent storage
type Repository interface {
	// Get will return a recipe by particular id from data store
	Get(id int) (*Recipe, error)

	// Get all will retrieve all recipe from data store
	GetAll() ([]*Recipe, error)

	// Add will create new recipe to data store
	Add(recipe *Recipe) (*Recipe, error)

	// Update will modify recipe value to data store
	Update(recipe *Recipe) error

	// Delete will remove recipe from data store
	Delete(recipe *Recipe) error
}
