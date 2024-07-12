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
func Inp() (string, int) {
	var Itemname string
	var Itemquantity int

	fmt.Printf("Enter your Order : ")
	fmt.Scan(&Itemname)

	fmt.Printf("Enter Quantity of Item : ")
	fmt.Scan(&Itemquantity)

	return Itemname, Itemquantity
}

func (fi *Ordercontroller) PrintTotalPrice(Itemname string, Itemquantity int) {
	fmt.Printf("TotalPrice : %.2f\n", fi.Ord.TotalPrice)
	fmt.Printf("You have ordered %d %s", Itemquantity, Itemname)
}
