package main

import (
	"fmt"
	"slices"
)

func main() {

	fruits := []string{"Apple", "Orange", "Watermelon", "Strawberry", "Mango", "Watermelon", "Watermelon", "Mango", "Mango"}

	cart := make([]string, 0)

	s := make(map[string]struct{})

	for i := 0; i < len(fruits); i++ {
		s[fruits[i]] = struct{}{}
	}

	for key := range s {
		cart = append(cart, key)
	}
	slices.Sort(cart)

	fmt.Println(cart)
}
