package bukaresep

import (
	"errors"
)

// Service is an interface as bukaresep business process (requirement) usecases
type Service interface {
	// GetRecipe will retrieve a recipe by particular id
	GetRecipe(id int) (*Recipe, error)

	// GetAllRecipe will retrieve all recipe
	GetAllRecipe() ([]*Recipe, error)

	// AddRecipe will create new recipe
	AddRecipe(name string, description string, ingredients string, instructions string) (*Recipe, error)

	// UpdateRecipe will modify recipe value
	UpdateRecipe(recipe *Recipe) error

	// DeleteRecipe will remove a recipe
	DeleteRecipe(recipe *Recipe) error
}

type service struct {
	repo Repository
}

// NewService will create a service instance
func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) GetRecipe(id int) (*Recipe, error) {
	return s.repo.Get(id)
}

func (s *service) GetAllRecipe() ([]*Recipe, error) {
	return s.repo.GetAll()
}

func (s *service) AddRecipe(name string, description string, ingredients string, instructions string) (*Recipe, error) {
	recipe := Recipe{
		Name:         name,
		Description:  description,
		Ingredients:  ingredients,
		Instructions: instructions,
	}

	if !recipe.isValid() {
		return nil, errors.New("Recipe is invalid")
	}

	return s.repo.Add(&recipe)
}

func (s *service) UpdateRecipe(recipe *Recipe) error {
	if !recipe.isValid() {
		return errors.New("Recipe is invalid")
	}

	return s.repo.Update(recipe)
}

func (s *service) DeleteRecipe(recipe *Recipe) error {
	return s.repo.Delete(recipe)
}
