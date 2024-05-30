package main

import (
	"fmt"
	"math"
)

func main() {

	x := []string{"Orange", "Apple", "Watermelon", "Melon", "Strawberry", "Orange", "Orange", "Melon", "Melon"}

	y := map[string]int{}

	for _, v := range x {
		y[v] = y[v] + 1
	}
	fmt.Println(y)

	price := map[string]int{
		"Orange":     150,
		"Apple":      200,
		"Watermelon": 250,
		"Melon":      300,
		"Strawberry": 350,
	}

	cart := map[string]int{}

	for product := range y {
		cart[product] = y[product] * price[product]
	}

	fmt.Println(cart)

	maxPrice := math.MinInt
	maxProduct := ""

	for prod, total := range cart {
		if maxPrice < total {
			maxPrice = total
			maxProduct = prod
		}

	}

	fmt.Println(maxPrice, maxProduct)

	for prod, total := range cart {
		if total >= 300 {
			cart[prod] = cart[prod] - 50
		}
	}
	fmt.Println(cart)

	minPrice := math.MaxInt
	minProduct := ""

	for prod, total := range cart {
		if total <= minPrice {
			minPrice = total
			minProduct = prod
		}
	}
	fmt.Println(minPrice, minProduct)

	delete(cart, minProduct)
	fmt.Println(cart)

	summ := 0

	for _, total := range cart {
		summ = summ + total

	}
	fmt.Println(summ)
}
