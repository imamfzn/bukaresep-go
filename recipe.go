package bukaresep

// Recipe is a struct representation of a recipe
type Recipe struct {
	ID           int    `xorm: "pk autoincr 'id'"`
	Name         string `xorm: "not null"`
	Description  string `xorm: "not null"`
	Ingredients  string `xorm: "not null"`
	Instructions string `xorm: "not null"`
}

// IsValid return true if recipe is valid; false otherwise
// Recipe will valid, if all field all filled
func (recipe *Recipe) isValid() bool {
	return recipe.Name != "" && recipe.Description != "" && recipe.Ingredients != "" && recipe.Instructions != ""
}
