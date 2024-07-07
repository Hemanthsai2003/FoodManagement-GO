package main

import FoodItemController "FoodManagement/Controller"

func main() {
	dosa := FoodItemController.Newfooditem("dosa", 50.55, 500)
	idly := FoodItemController.Newfooditem("idly", 30.25, 200)

	dosa.UpdatePrice(60)
	idly.ChangeName("podi idly")

	dosa.Display()
	idly.Display()
}
