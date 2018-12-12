package repository

import (
	"github.com/go-xorm/xorm"
	"github.com/imamfzn/bukaresep-go/entity"
)

type xormRepository struct {
	db *xorm.Engine
}

// NewRepository will return an implementation of Repository.
// It will use xorm driver implementation.
func NewRepository(db *xorm.Engine) (Repository, error) {
	return &xormRepository{db: db}, nil
}

// Get will return a row by recipe id that has been transformed to Recipe by xorm
// if recipe not, it will return recipe object will ID = 0
// It also will return an error if error occured from database
func (repo *xormRepository) Get(id int) (*entity.Recipe, error) {
	recipe := &entity.Recipe{}

	_, err := repo.db.ID(id).Get(recipe)

	return recipe, err
}

// Get all will return all row of recipe table and will transformed as list of Recipe by xorm.
// It also return an error if error occured from database.
func (repo *xormRepository) GetAll() ([]*entity.Recipe, error) {
	recipes := []*entity.Recipe{}

	err := repo.db.Find(&recipes)

	return recipes, err
}

// Add will insert a new recipe to database.
// It also return an error if error occured from database.
func (repo *xormRepository) Add(recipe *entity.Recipe) (*entity.Recipe, error) {
	_, err := repo.db.Insert(recipe)

	return recipe, err
}

// Update will update the row recipe from database by recipe object from parameter.
// It will return an error if error occured from database.
func (repo *xormRepository) Update(recipe *entity.Recipe) error {
	_, err := repo.db.ID(recipe.ID).Update(recipe)

	return err
}

// Delte will remove the row by recipe ID from recipe obeject parameter.
// It will return an error if error occured from database.
func (repo *xormRepository) Delete(recipe *entity.Recipe) error {
	_, err := repo.db.ID(recipe.ID).Delete(recipe)

	return err
}
