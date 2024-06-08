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
		{name: "Alex", balance: 600},
	}
	fmt.Println(us)

	max := math.MinInt
	maxUsers := []User{}

	for _, u := range us {
		if u.balance > max {
			max = u.balance
			maxUsers = []User{}
			maxUsers = append(maxUsers, u)
		} else if u.balance == max {
			maxUsers = append(maxUsers, u)
		}
	}
	fmt.Println(maxUsers)
}

type User struct {
	name    string
	balance int
}
