package database_test

import (
	"github.com/imamfzn/bukaresep-go"
	"github.com/imamfzn/bukaresep-go/database"

	"os"
	"testing"
)

func CreateRepository() (bukaresep.Repository, error) {
	os.Setenv("BUKARESEP_DB_FILENAME", "file::memory:")

	repo, err := database.NewRepository()

	if err != nil {
		return nil, err
	}

	return repo, nil
}

func TestNewRepository(t *testing.T) {
	t.Run("right db filename", func(t *testing.T) {
		_, err := CreateRepository()

		if err != nil {
			t.Fail()
		}
	})

	t.Run("wrong db filename", func(t *testing.T) {
		os.Setenv("BUKARESEP_DB_FILENAME", "/dev/null")

		_, err := database.NewRepository()

		if err == nil {
			t.Fail()
		}
	})
}

func TestAdd(t *testing.T) {
	repo, err := CreateRepository()

	if err != nil {
		t.Fail()
	}

	_, err = repo.Add(&bukaresep.Recipe{Name: "Food-name", Description: "food-desc", Ingredients: "food-ings", Instructions: "food-instr"})

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

	sampleRecipe, err := repo.Add(&bukaresep.Recipe{Name: "Food-name", Description: "food-desc", Ingredients: "food-ings", Instructions: "food-instr"})

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
		_, err = repo.Add(&bukaresep.Recipe{Name: "Food-name", Description: "food-desc", Ingredients: "food-ings", Instructions: "food-instr"})

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

	sampleRecipe, err := repo.Add(&bukaresep.Recipe{Name: "Food-name", Description: "food-desc", Ingredients: "food-ings", Instructions: "food-instr"})

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

	sampleRecipe, err := repo.Add(&bukaresep.Recipe{Name: "Food-name", Description: "food-desc", Ingredients: "food-ings", Instructions: "food-instr"})

	if err != nil {
		t.Fail()
	}

	err = repo.Delete(sampleRecipe)

	if err != nil {
		t.Fail()
	}
}
