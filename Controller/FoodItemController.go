package FoodItemController

import (
	FoodItem "FoodManagement/Entities"
	"fmt"
)

type Fooditemcontroller struct {
	Item FoodItem.Fooditem
}

func Newfooditem(name1 string, price float64, calories int) *Fooditemcontroller {
	return &Fooditemcontroller{
		Item: FoodItem.Fooditem{Name: name1,
			Price:    price,
			Calories: calories},
	}
}

func (fi *Fooditemcontroller) ChangeName(newName string) {
	fi.Item.Name = newName
}
func (fi *Fooditemcontroller) UpdatePrice(newPrice int) {
	fi.Item.Price = float64(newPrice)
}

func (fi *Fooditemcontroller) Display() {
	fmt.Printf("name: %s\nprice: %.2f\ncalories: %d", fi.Item.Name, fi.Item.Price, fi.Item.Calories)
}
