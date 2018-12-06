package database

import (
	"github.com/go-xorm/xorm"
	"github.com/imamfzn/bukaresep-go"
	_ "github.com/mattn/go-sqlite3"
	"github.com/subosito/gotenv"

	"os"
)

const DEFAULT_DB_DRIVER = "sqlite3"

type repository struct {
	db *xorm.Engine
}

// New repository will return an implementation of Repository.
// It will use sqlite3 as database driver implementation.
func NewRepository() (bukaresep.Repository, error) {
	gotenv.Load()

	db, err := xorm.NewEngine(DEFAULT_DB_DRIVER, os.Getenv("BUKARESEP_DB_FILENAME"))

	if err != nil {
		return nil, err
	}

	err = db.Sync(new(bukaresep.Recipe))

	if err != nil {
		return nil, err
	}

	return &repository{db: db}, nil
}

// Get will return a row by recipe id that has been transformed to Recipe by xorm
// if recipe not, it will return recipe object will ID = 0
// It also will return an error if error occured from database
func (repo *repository) Get(id int) (*bukaresep.Recipe, error) {
	recipe := &bukaresep.Recipe{}

	_, err := repo.db.ID(id).Get(recipe)

	return recipe, err
}

// Get all will return all row of recipe table and will transformed as list of Recipe by xorm.
// It also return an error if error occured from database.
func (repo *repository) GetAll() ([]*bukaresep.Recipe, error) {
	recipes := []*bukaresep.Recipe{}

	err := repo.db.Find(&recipes)

	return recipes, err
}

// Add will insert a new recipe to database.
// It also return an error if error occured from database.
func (repo *repository) Add(recipe *bukaresep.Recipe) (*bukaresep.Recipe, error) {
	_, err := repo.db.Insert(recipe)

	return recipe, err
}

// Update will update the row recipe from database by recipe object from parameter.
// It will return an error if error occured from database.
func (repo *repository) Update(recipe *bukaresep.Recipe) error {
	_, err := repo.db.ID(recipe.ID).Update(recipe)

	return err
}

// Delte will remove the row by recipe ID from recipe obeject parameter.
// It will return an error if error occured from database.
func (repo *repository) Delete(recipe *bukaresep.Recipe) error {
	_, err := repo.db.ID(recipe.ID).Delete(recipe)

	return err
}
