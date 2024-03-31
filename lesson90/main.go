package main

import "fmt"

func main() {

	var money = map[string]int{
		"dolars": 100,
		"euros":  200,
		"apple":  3,
	}
	money["orange"] = 10
	fmt.Println(money)

	var money2 = map[string]int{}
	for x, y := range money {
		money2[x] = y
	}
	fmt.Println(money2)
	delete(money, "dolars")
	fmt.Println(money, money2)
}
