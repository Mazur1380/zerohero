package main

import "fmt"

func main() {

	fruits := []string{"Apple", "Orange", "Watermelon", "Strawberry", "Mango", "Watermelon", "Watermelon", "Mango", "Mango"}

	cart := make([]string, 0)

	for i := 0; i < len(fruits); i++ {
		f := fruits[i]
		if !exist(cart, f) {
			cart = append(cart, f)
		}
	}
	fmt.Println(cart)

	// item := "Mango"

	// if exist(fruits, item) {
	// 	fmt.Println("Нашли фрукт")
	// } else {
	// 	fmt.Println("Не нашли фрукт")
	// }
}

func exist(x []string, v string) bool {
	for i := 0; i < len(x); i++ {
		if x[i] == v {
			return true
		}
	}
	return false
}
