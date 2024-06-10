package main

import (
	"fmt"
	"log"

	"github.com/andreipimenov/zerohero/lesson141/store"
)

func main() {

	s := store.New()
	s.CreateUser(1, "Vlad", 1000)
	u, err := s.GetUserById(1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(u)
	s.CreateProduct(1, "Apple", 150)
	p, err := s.GetProductById(1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(p)
	err = s.CreateOrder(1, 1, 1)
	if err != nil {
		log.Fatal(err)
	}
	o, err := s.GetOrderById(1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(o)
	fmt.Println(u)

	s.CreateUser(2, "Djas", 300)
	s.CreateUser(3, "Dima", 400)

	s.CreateProduct(2, "Orange", 200)
	s.CreateProduct(3, "Watermelon", 300)

	s.CreateOrder(2, 2, 2)
	s.CreateOrder(3, 1, 3)
	hs, err := s.GetOrderHistory(1)
	if err != nil {
		log.Fatal(err)
	}
	for _, h := range hs {
		fmt.Println(h)
	}

}
