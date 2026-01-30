package model

import "fmt"

// ✅ Struct Definition
type Model struct {
	Name   string
	Brand  string
	Price  int
	ModelID string
}

// ✅ Function to Create Model Data
func CreateModelData() []Model {

	models := []Model{
		{Name: "iPhone 15", Brand: "Apple", Price: 80000, ModelID: "M101"},
		{Name: "Galaxy S24", Brand: "Samsung", Price: 75000, ModelID: "M102"},
		{Name: "OnePlus 12", Brand: "OnePlus", Price: 65000, ModelID: "M103"},
	}

	return models
}

// ✅ Function to Display Model Data
func Display(models []Model) {

	for i := 0; i < len(models); i++ {

		fmt.Println("Model", i+1)
		fmt.Println("Name    :", models[i].Name)
		fmt.Println("Brand   :", models[i].Brand)
		fmt.Println("Price   :", models[i].Price)
		fmt.Println("ModelID :", models[i].ModelID)
		fmt.Println("------------------------")
	}
}
