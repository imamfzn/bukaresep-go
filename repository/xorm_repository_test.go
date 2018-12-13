package repository_test

import (
	"github.com/go-xorm/xorm"
	"github.com/imamfzn/bukaresep-go/entity"
	"github.com/imamfzn/bukaresep-go/repository"

	// it required for xorm
	_ "github.com/mattn/go-sqlite3"
	"testing"
)

func CreateRepository() (repository.Repository, error) {
	db, err := xorm.NewEngine("sqlite3", "file::memory:")

	if err != nil {
		return nil, err
	}

	err = db.Sync(new(entity.Recipe))

	if err != nil {
		return nil, err
	}

	repo, err := repository.NewXormRepository(db)

	if err != nil {
		return nil, err
	}

	return repo, nil
}

func TestAdd(t *testing.T) {
	repo, err := CreateRepository()

	if err != nil {
		t.Fail()
	}

	_, err = repo.Add(&entity.Recipe{Name: "Food-name", Description: "food-desc", Ingredients: "food-ings", Instructions: "food-instr"})

	if err != nil {
		t.Fail()
	}

	recipes, err := repo.GetAll()

	if err != nil {
		t.Fail()
	}

	if len(recipes) != 1 {
		t.Fail()
	}
}

func TestGet(t *testing.T) {
	repo, err := CreateRepository()

	if err != nil {
		t.Fail()
	}

	sampleRecipe, err := repo.Add(&entity.Recipe{Name: "Food-name", Description: "food-desc", Ingredients: "food-ings", Instructions: "food-instr"})

	if err != nil {
		t.Fail()
	}

	t.Run("exists id", func(t *testing.T) {
		_, err := repo.Get(sampleRecipe.ID)

		if err != nil {
			t.Fail()
		}
	})

	t.Run("not exists id", func(t *testing.T) {
		recipe, err := repo.Get(-1)

		if err != nil {
			t.Fail()
		}

		if recipe.ID != 0 {
			t.Fail()
		}
	})
}

func TestGetAll(t *testing.T) {
	repo, err := CreateRepository()

	if err != nil {
		t.Fail()
	}

	numOfSample := 3

	for i := 0; i < numOfSample; i++ {
		_, err = repo.Add(&entity.Recipe{Name: "Food-name", Description: "food-desc", Ingredients: "food-ings", Instructions: "food-instr"})

		if err != nil {
			t.Fail()
		}
	}

	recipes, err := repo.GetAll()

	if err != nil {
		t.Fail()
	}

	if len(recipes) != numOfSample {
		t.Fail()
	}
}

func TestUpdate(t *testing.T) {
	repo, err := CreateRepository()

	if err != nil {
		t.Fail()
	}

	sampleRecipe, err := repo.Add(&entity.Recipe{Name: "Food-name", Description: "food-desc", Ingredients: "food-ings", Instructions: "food-instr"})

	if err != nil {
		t.Fail()
	}

	updatedName := "chicken katsu v2.0"
	sampleRecipe.Name = updatedName

	err = repo.Update(sampleRecipe)

	if err != nil {
		t.Fail()
	}

	if sampleRecipe.Name != updatedName {
		t.Fail()
	}
}

func TestDelete(t *testing.T) {
	repo, err := CreateRepository()

	if err != nil {
		t.Fail()
	}

	sampleRecipe, err := repo.Add(&entity.Recipe{Name: "Food-name", Description: "food-desc", Ingredients: "food-ings", Instructions: "food-instr"})

	if err != nil {
		t.Fail()
	}

	err = repo.Delete(sampleRecipe)

	if err != nil {
		t.Fail()
	}
}
