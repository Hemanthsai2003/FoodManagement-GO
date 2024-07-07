package Order

import FoodItem "FoodManagement/Entities"

type Order struct {
	FoodRequest []struct {
		Item     *FoodItem.Fooditem
		Quantity int
	}
	TotalPrice float64
}
