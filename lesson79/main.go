package main

import "fmt"

func main() {

	a := Account{
		Name:    "Vlad",
		Balance: 500,
	}
	b := Account{
		Name:    "Vlad",
		Balance: 500,
	}

	c := Account{
		Name:    "Vlad",
		Balance: 500,
	}
	sum := 0
	x := []Account{a, b, c}
	for _, i := range x {
		sum = sum + i.Balance
	}
	fmt.Println(sum)

}

type Account struct {
	Name    string
	Balance int
}
