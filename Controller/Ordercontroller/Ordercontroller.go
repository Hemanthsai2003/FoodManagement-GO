package Ordercontroller

import (
	"FoodManagement/Entities/Order"
	"fmt"
)

type Ordercontroller struct {
	Ord Order.Order
}

func Neworder(orders []*Order.OrderItem) *Ordercontroller {
	var totalPrice float64 = 0
	for i := 0; i < len(orders); i++ {
		totalPrice = totalPrice + (orders[i].Item.Price * float64(orders[i].Quantity))
	}
	return &Ordercontroller{
		Ord: Order.Order{FoodRequest: orders, TotalPrice: totalPrice},
	}
}

func (fi *Ordercontroller) PrintTotalPrice() {
	fmt.Printf("TotalPrice : %.2f", fi.Ord.TotalPrice)
}
