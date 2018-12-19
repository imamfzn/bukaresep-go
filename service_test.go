package bukaresep_test

import (
	"errors"
	"github.com/imamfzn/bukaresep-go"
	"github.com/imamfzn/bukaresep-go/entity"
	"testing"
)

func TestGetRecipe(t *testing.T) {
	service := bukaresep.NewService(CreateMockRepo())

	t.Run("get exists recipe", func(t *testing.T) {
		_, err := service.GetRecipe(1)

		if err != nil {
			t.FailNow()
		}
	})

	t.Run("get not exists recipe", func(t *testing.T) {
		_, err := service.GetRecipe(-1)

		if err == nil {
			t.FailNow()
		}
	})
}

func TestGetAllRecipe(t *testing.T) {
	service := bukaresep.NewService(CreateMockRepo())

	recipes, err := service.GetAllRecipe()

	if err != nil {
		t.FailNow()
	}

	if len(recipes) != 3 {
		t.FailNow()
	}
}

func TestAdd(t *testing.T) {
	service := bukaresep.NewService(CreateMockRepo())

	t.Run("add valid recipe", func(t *testing.T) {
		recipes, err := service.GetAllRecipe()

		if err != nil {
			t.FailNow()
		}

		recipesCount := len(recipes)

		_, err = service.AddRecipe("bento", "oriental food again", "bento ings", "bentos instr")

		if err != nil {
			t.FailNow()
		}

		currentRecipes, err := service.GetAllRecipe()

		if err != nil {
			t.FailNow()
		}

		if len(currentRecipes) != recipesCount+1 {
			t.FailNow()
		}
	})

	t.Run("add invalid recipe", func(t *testing.T) {
		recipes, err := service.GetAllRecipe()

		if err != nil {
			t.FailNow()
		}

		recipesCount := len(recipes)

		_, err = service.AddRecipe("", "oriental food again", "bento ings", "bentos instr")

		if err == nil {
			t.FailNow()
		}

		currentRecipes, err := service.GetAllRecipe()

		if err != nil {
			t.FailNow()
		}

		if len(currentRecipes) != recipesCount {
			t.FailNow()
		}

	})
}

func TestUpdate(t *testing.T) {
	service := bukaresep.NewService(CreateMockRepo())

	t.Run("update valid recipe", func(t *testing.T) {
		recipe, err := service.GetRecipe(1)

		if err != nil {
			t.FailNow()
		}

		recipe.Name = "updated name"

		err = service.UpdateRecipe(recipe)

		if err != nil {
			t.FailNow()
		}
	})

	t.Run("update invalid recipe", func(t *testing.T) {
		recipe, err := service.GetRecipe(1)

		if err != nil {
			t.FailNow()
		}

		recipe.Name = ""

		err = service.UpdateRecipe(recipe)

		if err == nil {
			t.FailNow()
		}
	})
}

func TestDelete(t *testing.T) {
	service := bukaresep.NewService(CreateMockRepo())
	recipes, err := service.GetAllRecipe()

	if err != nil {
		t.FailNow()
	}

	recipesCount := len(recipes)

	recipe, err := service.GetRecipe(1)

	if err != nil {
		t.FailNow()
	}

	err = service.DeleteRecipe(recipe)

	if err != nil {
		t.FailNow()
	}

	currentRecipes, err := service.GetAllRecipe()

	if err != nil {
		t.FailNow()
	}

	if len(currentRecipes) != recipesCount-1 {
		t.FailNow()
	}
}

type mockRepo struct {
	nextID  int
	count   int
	storage map[int]*entity.Recipe
}

func CreateMockRepo() *mockRepo {
	repo := mockRepo{1, 0, map[int]*entity.Recipe{}}

	// Suply sample data
	repo.Add(&entity.Recipe{Name: "Chicken Katsu", Description: "Oriental Food", Ingredients: "Food Ingredients", Instructions: "Recipe instructions"})
	repo.Add(&entity.Recipe{Name: "Nasi Padang", Description: "Indonesian Food", Ingredients: "Food Ingredients", Instructions: "Recipe instructions"})
	repo.Add(&entity.Recipe{Name: "Dorayaki", Description: "Doraemon Cake", Ingredients: "Cake Ingredients", Instructions: "Recipe instructions"})

	return &repo
}

func (repo *mockRepo) Get(id int) (*entity.Recipe, error) {
	recipe, exists := repo.storage[id]

	if !exists {
		return nil, errors.New("not found")
	}

	return recipe, nil
}

func (repo *mockRepo) GetAll() ([]*entity.Recipe, error) {
	recipes := []*entity.Recipe{}

	for _, recipe := range repo.storage {
		recipes = append(recipes, recipe)
	}

	return recipes, nil
}

func (repo *mockRepo) Add(recipe *entity.Recipe) (*entity.Recipe, error) {
	recipe.ID = repo.nextID
	repo.nextID++
	repo.count++

	repo.storage[recipe.ID] = recipe

	return recipe, nil
}

func (repo *mockRepo) Update(recipe *entity.Recipe) error {
	if _, exists := repo.storage[recipe.ID]; !exists {
		return errors.New("recipe not found")
	}

	repo.storage[recipe.ID] = recipe

	return nil
}

func (repo *mockRepo) Delete(recipe *entity.Recipe) error {
	delete(repo.storage, recipe.ID)

	repo.count--

	return nil
}
