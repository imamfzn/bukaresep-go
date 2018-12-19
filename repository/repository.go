package repository

import "github.com/imamfzn/bukaresep-go/entity"

// Repository is an interface as data access to a persistent storage
type Repository interface {
	// Get will return a recipe by particular id from data store
	Get(id int) (*entity.Recipe, error)

	// Get all will retrieve all recipe from data store
	GetAll() ([]*entity.Recipe, error)

	// Add will create new recipe to data store
	Add(recipe *entity.Recipe) (*entity.Recipe, error)

	// Update will modify recipe value to data store
	Update(recipe *entity.Recipe) error

	// Delete will remove recipe from data store
	Delete(recipe *entity.Recipe) error
}
