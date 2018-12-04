package bukaresep

// Recipe is a struct representation of a recipe
type Recipe struct {
	ID           int
	Name         string
	Description  string
	Ingredients  string
	Instructions string
}

// IsValid return true if recipe is valid; false otherwise
// Recipe will valid, if all field all filled
func (recipe *Recipe) isValid() bool {
	return recipe.Name != "" && recipe.Description != "" && recipe.Ingredients != "" && recipe.Instructions != ""
}
