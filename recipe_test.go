package bukaresep

import (
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

			recipe := Recipe{1, name, description, ingredients, instructions}

			if !recipe.isValid() {
				t.Fail()
			}

		})

		t.Run("ID not filled", func(t *testing.T) {
			t.Parallel()

			recipe := Recipe{0, name, description, ingredients, instructions}

			if !recipe.isValid() {
				t.Fail()
			}
		})
	})

	t.Run("invalid cases", func(t *testing.T) {
		t.Parallel()

		t.Run("Name is blank", func(t *testing.T) {
			t.Parallel()

			recipe := Recipe{1, "", description, ingredients, instructions}

			if recipe.isValid() {
				t.Fail()
			}
		})

		t.Run("Description is blank", func(t *testing.T) {
			t.Parallel()

			recipe := Recipe{1, name, "", ingredients, instructions}

			if recipe.isValid() {
				t.Fail()
			}
		})

		t.Run("Ingredients is blank", func(t *testing.T) {
			t.Parallel()

			recipe := Recipe{1, name, description, "", instructions}

			if recipe.isValid() {
				t.Fail()
			}
		})

		t.Run("Instructions is blank", func(t *testing.T) {
			t.Parallel()

			recipe := Recipe{1, name, description, ingredients, ""}

			if recipe.isValid() {
				t.Fail()
			}
		})

		t.Run("All blank", func(t *testing.T) {
			t.Parallel()

			recipe := Recipe{0, "", "", "", ""}

			if recipe.isValid() {
				t.Fail()
			}
		})
	})

}
