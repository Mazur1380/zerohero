package main

import (
	"fmt"
	"math"
)

func main() {

	us := []User{
		{name: "Vlad", balance: 200},
		{name: "Djas", balance: 400},
		{name: "Dima", balance: 600},
	}
	fmt.Println(us)

	max := math.MinInt

	for _, u := range us {
		if u.balance > max {
			max = u.balance
		}
	}
	fmt.Println(max)
}

type User struct {
	name    string
	balance int
}
