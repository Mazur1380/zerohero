package main

import "fmt"

func main() {
	a := 1
	b := 1

	x := []int{a, b}
	for i := 0; i < 100; i++ {
		a, b = b, a+b
		x = append(x, b)
	}
	fmt.Println(x)

}
