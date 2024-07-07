package Order

import FoodItem "FoodManagement/Entities"

type Order struct {
	FoodRequest []*OrderItem
	TotalPrice  float64
}

type OrderItem struct {
	Item     *FoodItem.Fooditem
	Quantity int
}
