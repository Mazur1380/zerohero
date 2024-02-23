package main

import "fmt"

func main() {
	var x = []int{1, 2, 3, 4, 5}
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	x = append(x, 3, 5, 7)
	fmt.Println(x, len(x), cap(x))
	y := x[2:5]
	fmt.Println(y, len(y), cap(y))
	y[0] = 100
	fmt.Println(x, "\n", y)

	fmt.Println("_____________")
	z := make([]int, 3)
	copy(z, x[2:5])
	fmt.Println(z)
	z[0] = 999
	fmt.Println(x, z)
}
