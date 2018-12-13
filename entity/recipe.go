package entity

import "encoding/json"

// Recipe is a struct representation of a recipe
type Recipe struct {
	ID           int    `xorm:"pk autoincr 'id'" json:"id"`
	Name         string `xorm:"not null" json:"name"`
	Description  string `xorm:"not null" json:"description"`
	Ingredients  string `xorm:"not null" json:"ingredients"`
	Instructions string `xorm:"not null" json:"instructions"`
}

// IsValid return true if recipe is valid; false otherwise
// Recipe will valid, if all field all filled
func (recipe *Recipe) IsValid() bool {
	return recipe.Name != "" && recipe.Description != "" && recipe.Ingredients != "" && recipe.Instructions != ""
}

// ToJSON will return a transformed recipe struct to a json format
// as []byte data structure from json.Marshal
func (recipe *Recipe) ToJSON() ([]byte, error) {
	return json.Marshal(recipe)
}
