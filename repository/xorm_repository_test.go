package repository_test

import (
	"github.com/go-xorm/xorm"
	"github.com/imamfzn/bukaresep-go/entity"
	"github.com/imamfzn/bukaresep-go/repository"

	// it required for xorm
	_ "github.com/mattn/go-sqlite3"
	"testing"
)

var (
	sampleRecipe  = entity.Recipe{0, "chicken katsu", "japanese food", "chicken,salt,egg", "just merge all"}
	sampleRecipes = []entity.Recipe{
		entity.Recipe{1, "chicken katsu", "japanese food", "chicken,salt,egg", "just merge all"},
		entity.Recipe{2, "egg roll", "delicious egg", "egg", "roll the egg"},
		entity.Recipe{3, "meat ball", "indonesian food", "meat,salt", "make the meat like a ball"},
	}
)

func IsEqualRecipe(r1 *entity.Recipe, r2 *entity.Recipe) bool {
	return r1.Name == r2.Name &&
		r1.Description == r2.Description &&
		r1.Ingredients == r2.Ingredients &&
		r1.Instructions == r2.Instructions
}

func CreateRepository(t *testing.T) repository.Repository {
	db, err := CreateDatabase()

	if err != nil {
		t.Errorf("Errow while creating a database from xorm")
	}

	repo, err := repository.NewXormRepository(db)

	if err != nil {
		t.Errorf("Error while create a repository")
	}

	return repo
}

func CreateRepositoryWithSeed(t *testing.T, ns int) repository.Repository {
	repo := CreateRepository(t)

	if ns <= 0 {
		t.Errorf("num of seed must be positive")
	}

	if ns > len(sampleRecipes) {
		t.Errorf("num of seed is too large must be less or equal than %v", len(sampleRecipes))
	}

	for i := 0; i < ns; i++ {
		_, err := repo.Add(&sampleRecipes[i])

		if err != nil {
			t.Errorf("Errow while inserting sample recipe to db")
		}
	}

	return repo
}

func CreateDatabase() (*xorm.Engine, error) {
	db, err := xorm.NewEngine("sqlite3", "file::memory:")

	if err != nil {
		return nil, err
	}

	err = db.Sync(new(entity.Recipe))

	if err != nil {
		return nil, err
	}

	return db, nil
}

func TestNewXormRepository(t *testing.T) {
	db, err := CreateDatabase()

	if err != nil {
		t.Errorf("Errow while creating a database from xorm")
	}

	_, err = repository.NewXormRepository(db)

	if err != nil {
		t.Errorf("Error while create a repository")
	}
}

func TestAdd(t *testing.T) {
	repo := CreateRepository(t)

	actual, err := repo.Add(&sampleRecipe)

	if err != nil {
		t.Errorf("Errow while adding a recipe to database")
	}

	if !IsEqualRecipe(&sampleRecipe, actual) {
		t.Fail()
	}

}

func TestGet(t *testing.T) {
	repo := CreateRepository(t)

	expected, err := repo.Add(&sampleRecipe)

	if err != nil {
		t.Errorf("Errow while adding a sample recipe to database")
	}

	t.Run("exists id", func(t *testing.T) {
		actual, err := repo.Get(expected.ID)

		if err != nil {
			t.Errorf("Errow while getting a recipe by ID")
		}

		if !IsEqualRecipe(expected, actual) {
			t.Fail()
		}
	})

	t.Run("not exists id", func(t *testing.T) {
		recipe, err := repo.Get(-1)

		if err != nil {
			t.Errorf("Errow while getting a recipe by ID")
		}

		if recipe.ID != 0 {
			t.Fail()
		}
	})
}

func TestGetAll(t *testing.T) {
	numOfSample := len(sampleRecipes)
	repo := CreateRepositoryWithSeed(t, numOfSample)

	recipes, err := repo.GetAll()

	if err != nil {
		t.Errorf("Error while getting all recipe")
	}

	if len(recipes) != numOfSample {
		t.Fail()
	}

	for i := 0; i < numOfSample; i++ {
		expected := &sampleRecipes[i]
		actual := recipes[i]

		if !IsEqualRecipe(expected, actual) {
			t.Fail()
		}
	}

}

func TestUpdate(t *testing.T) {
	repo := CreateRepository(t)

	recipe, err := repo.Add(&sampleRecipe)

	if err != nil {
		t.Errorf("Error while inserting sample recipe")
	}

	recipe.Name = "chicken katsu v2.0"
	recipe.Description = "now chicken katsu is very delicious"
	recipe.Ingredients = "just need chicken"
	recipe.Instructions = "mix mix all ingredients"

	err = repo.Update(recipe)

	if err != nil {
		t.Errorf("Error while update a recipe")
	}

	updatedRecipe, err := repo.Get(recipe.ID)

	if err != nil {
		t.Errorf("Error while getting a recipe by ID")
	}

	if !IsEqualRecipe(recipe, updatedRecipe) {
		t.Fail()
	}
}

func TestDelete(t *testing.T) {
	t.Run("emptying recipe data", func(t *testing.T) {
		repo := CreateRepository(t)

		recipe, err := repo.Add(&sampleRecipe)

		if err != nil {
			t.Errorf("Error while inserting sample recipe")
		}

		err = repo.Delete(recipe)

		if err != nil {
			t.Errorf("Error while deleting a recipe")
		}

		recipes, err := repo.GetAll()

		if err != nil {
			t.Errorf("Error while getting all recipe")
		}

		if len(recipes) != 0 {
			t.Fail()
		}
	})

	t.Run("delete a recipe", func(t *testing.T) {
		repo := CreateRepositoryWithSeed(t, 2)

		expectedRecipes, err := repo.GetAll()

		if err != nil {
			t.Errorf("Error while getting all recipe for get sample data")
		}

		removedRecipe := expectedRecipes[1]

		err = repo.Delete(removedRecipe)

		if err != nil {
			t.Errorf("Error while deleting a recipe")
		}

		actualRecipes, err := repo.GetAll()

		if err != nil {
			t.Errorf("Error while getting all recipe")
		}

		if len(actualRecipes) != 1 {
			t.Errorf("expected size of recipe from db is %v, but actual is %v", 1, len(actualRecipes))
		}

		if !IsEqualRecipe(expectedRecipes[0], actualRecipes[0]) {
			t.Errorf("The recipe has been deleted maybe wrong")
		}
	})
}
