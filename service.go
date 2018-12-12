package bukaresep

import (
	"errors"

	"github.com/imamfzn/bukaresep-go/entity"
	"github.com/imamfzn/bukaresep-go/repository"
)

// Service is an interface as bukaresep business process (requirement) usecases.
type Service interface {
	// GetRecipe will retrieve a recipe by particular id.
	GetRecipe(id int) (*entity.Recipe, error)

	// GetAllRecipe will retrieve all recipe.
	GetAllRecipe() ([]*entity.Recipe, error)

	// AddRecipe will create new recipe.
	AddRecipe(name string, description string, ingredients string, instructions string) (*entity.Recipe, error)

	// UpdateRecipe will modify recipe value.
	UpdateRecipe(recipe *entity.Recipe) error

	// DeleteRecipe will remove a recipe.
	DeleteRecipe(recipe *entity.Recipe) error
}

type service struct {
	repo repository.Repository
}

// NewService will return an implementation of a Service. It should be supplied
// with an implementation o a Repository.
func NewService(repo repository.Repository) Service {
	return &service{
		repo: repo,
	}
}

// Get recipe will return a recipe by particular id from repository, and will return
// an error otherwise.
func (s *service) GetRecipe(id int) (*entity.Recipe, error) {
	return s.repo.Get(id)
}

// Get recipe will return a list of recipe from repository, and will return
// an error otherwise.
func (s *service) GetAllRecipe() ([]*entity.Recipe, error) {
	return s.repo.GetAll()
}

// Add recipe will create a new recipe with required name, description, ingeredients, and instructions
// It will return a Recipe created in repository, or an error if a the Recipe
// is invalid (it means if one of all field not filled).
func (s *service) AddRecipe(name string, description string, ingredients string, instructions string) (*entity.Recipe, error) {
	recipe := entity.Recipe{
		Name:         name,
		Description:  description,
		Ingredients:  ingredients,
		Instructions: instructions,
	}

	if !recipe.IsValid() {
		return nil, errors.New("Recipe is invalid")
	}

	return s.repo.Add(&recipe)
}

// Update recipe will update to Recipe in the repository with the values suplied in the
// recipe object parameter. It also will return an error if the recipe become invalid.
func (s *service) UpdateRecipe(recipe *entity.Recipe) error {
	if !recipe.IsValid() {
		return errors.New("Recipe is invalid")
	}

	return s.repo.Update(recipe)
}

// Delete recipe will remove the recipe object from repository.
func (s *service) DeleteRecipe(recipe *entity.Recipe) error {
	return s.repo.Delete(recipe)
}
