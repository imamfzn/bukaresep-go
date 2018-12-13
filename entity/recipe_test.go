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

		t.Run("all attributes filed", func(t *testing.T) {
			t.Parallel()

			recipe := &entity.Recipe{1, name, description, ingredients, instructions}

			if !recipe.IsValid() {
				t.Fail()
			}

		})

		t.Run("ID not filled", func(t *testing.T) {
			t.Parallel()

			recipe := entity.Recipe{0, name, description, ingredients, instructions}

			if !recipe.IsValid() {
				t.Fail()
			}
		})
	})

	t.Run("invalid cases", func(t *testing.T) {
		t.Parallel()

		t.Run("Name is blank", func(t *testing.T) {
			t.Parallel()

			recipe := entity.Recipe{1, "", description, ingredients, instructions}

			if recipe.IsValid() {
				t.Fail()
			}
		})

		t.Run("Description is blank", func(t *testing.T) {
			t.Parallel()

			recipe := entity.Recipe{1, name, "", ingredients, instructions}

			if recipe.IsValid() {
				t.Fail()
			}
		})

		t.Run("Ingredients is blank", func(t *testing.T) {
			t.Parallel()

			recipe := entity.Recipe{1, name, description, "", instructions}

			if recipe.IsValid() {
				t.Fail()
			}
		})

		t.Run("Instructions is blank", func(t *testing.T) {
			t.Parallel()

			recipe := entity.Recipe{1, name, description, ingredients, ""}

			if recipe.IsValid() {
				t.Fail()
			}
		})

		t.Run("All blank", func(t *testing.T) {
			t.Parallel()

			recipe := entity.Recipe{0, "", "", "", ""}

			if recipe.IsValid() {
				t.Fail()
			}
		})
	})

}

func TestToJSON(t *testing.T) {
	recipe := entity.Recipe{1, "food", "desc", "ing", "instr"}
	expected := `{"id":1,"name":"food","description":"desc","ingredients":"ing","instructions":"instr"}`

	actual, err := recipe.ToJSON()

	if err != nil {
		t.Fail()
	}

	if string(actual) != expected {
		t.Fail()
	}
}
