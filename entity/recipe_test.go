package entity_test

import (
	"github.com/imamfzn/bukaresep-go/entity"
	"testing"
)

func TestValidity(t *testing.T) {
	name := "chicken katsu"
	description := "oriental food"
	ingredients := "chicken katsu ingredients"
	instructions := "chicken katsu instructions"

	t.Run("valid cases", func(t *testing.T) {
		t.Parallel()

		t.Run("all attributes filled", func(t *testing.T) {
			t.Parallel()

			recipe := &entity.Recipe{
				ID:           1,
				Name:         name,
				Description:  description,
				Ingredients:  ingredients,
				Instructions: instructions,
			}

			if !recipe.IsValid() {
				t.Fail()
			}

		})

		t.Run("ID not filled", func(t *testing.T) {
			t.Parallel()

			recipe := entity.Recipe{
				ID:           0,
				Name:         name,
				Description:  description,
				Ingredients:  ingredients,
				Instructions: instructions,
			}

			if !recipe.IsValid() {
				t.Fail()
			}
		})
	})

	t.Run("invalid cases", func(t *testing.T) {
		t.Parallel()

		t.Run("Name is blank", func(t *testing.T) {
			t.Parallel()

			recipe := entity.Recipe{
				ID:           1,
				Name:         "",
				Description:  description,
				Ingredients:  ingredients,
				Instructions: instructions,
			}

			if recipe.IsValid() {
				t.Fail()
			}
		})

		t.Run("Description is blank", func(t *testing.T) {
			t.Parallel()

			recipe := entity.Recipe{
				ID:           1,
				Name:         name,
				Description:  "",
				Ingredients:  ingredients,
				Instructions: instructions,
			}

			if recipe.IsValid() {
				t.Fail()
			}
		})

		t.Run("Ingredients is blank", func(t *testing.T) {
			t.Parallel()

			recipe := entity.Recipe{
				ID:           1,
				Name:         name,
				Description:  description,
				Ingredients:  "",
				Instructions: instructions,
			}

			if recipe.IsValid() {
				t.Fail()
			}
		})

		t.Run("Instructions is blank", func(t *testing.T) {
			t.Parallel()

			recipe := entity.Recipe{
				ID:           1,
				Name:         name,
				Description:  description,
				Ingredients:  ingredients,
				Instructions: "",
			}

			if recipe.IsValid() {
				t.Fail()
			}
		})

		t.Run("All blank", func(t *testing.T) {
			t.Parallel()

			recipe := entity.Recipe{
				Name:         "",
				Description:  "",
				Ingredients:  "",
				Instructions: "",
			}

			if recipe.IsValid() {
				t.Fail()
			}
		})
	})

}

func TestToJSON(t *testing.T) {
	recipe := entity.Recipe{
		ID:           1,
		Name:         "food",
		Description:  "desc",
		Ingredients:  "ing",
		Instructions: "instr",
	}

	expected := `{"id":1,"name":"food","description":"desc","ingredients":"ing","instructions":"instr"}`

	actual, err := recipe.ToJSON()

	if err != nil {
		t.Fail()
	}

	if string(actual) != expected {
		t.Fail()
	}
}
