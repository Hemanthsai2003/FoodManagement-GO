package main

import (
	FoodItemController "FoodManagement/Controller"
	"FoodManagement/Controller/Ordercontroller"
	"FoodManagement/Entities/Order"
)

func main() {
	dosa := FoodItemController.Newfooditem("dosa", 50.55, 500)
	idly := FoodItemController.Newfooditem("idly", 30.25, 200)

	dosa.UpdatePrice(60)
	idly.ChangeName("podi idly")

	orderTemp := &Order.OrderItem{Item: &dosa.Item, Quantity: 1}
	orderTemp1 := &Order.OrderItem{Item: &idly.Item, Quantity: 10}
	order1 := Ordercontroller.Neworder([]*Order.OrderItem{orderTemp, orderTemp1})
	order1.PrintTotalPrice()
}
