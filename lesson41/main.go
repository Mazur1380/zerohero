package main

import "fmt"

func main() {

	x := []int{1, 2, 3, 4, 5, 8, 9, 7, 67}
	y := []int{5, 6, 7, 8, 9}

	//for _, i := range x {
	//	y = append(y, i)

	//}
	//fmt.Println(y)
	//var z []int
	//for i := 0; i < len(x); i++ {
	//	z = append(z, x[i], y[i])

	//}
	//fmt.Println(z)
	z := make([]int, len(x)+len(y))

	copy(z, x)
	copy(z[len(x):], y)
	fmt.Println(z)
}
