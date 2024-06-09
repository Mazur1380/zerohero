package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	order := &Order{
		ID:             1,
		Status:         "Created",
		DeliveryAdress: "SPB",
	}
	apple := Product{
		ID:    1,
		Price: 100,
		Name:  "Apple",
	}
	orange := Product{
		ID:    2,
		Price: 200,
		Name:  "Orange",
	}
	order.AddItem(apple, 5)
	order.AddItem(orange, 1)
	order.SetCost()
	order.AddItem(orange, 1)

	fmt.Printf("%+v\n", order)
}

type Order struct {
	ID             int
	Cost           int
	Status         string
	DeliveryAdress string
	Items          []*OrderItem
}

type Product struct {
	ID    int
	Price int
	Name  string
}

type OrderItem struct {
	Product Product
	Count   int
}

func (o *Order) AddItem(p Product, c int) {
	item := &OrderItem{
		Product: p,
		Count:   c,
	}
	for _, i := range o.Items {
		if i.Product.ID == p.ID {
			i.Count = i.Count + c
			return
		}
	}
	o.Items = append(o.Items, item)
}

func (o *Order) SetCost() {
	s := 0
	for _, i := range o.Items {
		s += i.Count * i.Product.Price
	}
	o.Cost = s
}

func (o *Order) String() string {
	b, _ := json.MarshalIndent(o, "", "    ")
	return string(b)
}
