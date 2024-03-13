package main

import "fmt"

func main() {
	h := Human{
		Name: "Djas",
		Age:  25,
	}
	h.Walk()
	a := Animal{
		Name: "Cat",
	}
	a.Walk()
	var w Walker
	w = a
	w.Walk()
}

type Human struct {
	Name string
	Age  int
}

func (h Human) Walk() {
	fmt.Println("Человек идет")
}

type Animal struct {
	Name string
}

func (a Animal) Walk() {
	fmt.Println("Животное идет")
}

type Walker interface {
	Walk()
}
