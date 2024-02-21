package main

import "fmt"

func main() {
	var a []int
	a = append(a, 4, 5, 6, 7, 8)
	fmt.Println(a)
	a = append(a, 1, 2, 3)
	fmt.Println(a)
	b := []int{
		1, 2, 3, 4,
	}
	fmt.Println(a, b)
	a = b
	fmt.Println(a, b)
	x := len(a)
	fmt.Println(x)
	a = append(a, 4, 59, 6)
	x = len(a)
	fmt.Println(a, x)
}
