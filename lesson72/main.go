package main

import "fmt"

func main() {

	u := user{
		name: "Vlad",
		age:  20,
	}
	fmt.Println(u)
	fmt.Println(u.GetName())
	u2 := user{
		name: "Djas",
		age:  25,
	}
	fmt.Println(u2.GetName())
}

type user struct {
	name string
	age  int
}

func (u user) GetName() string {
	return u.name
}
