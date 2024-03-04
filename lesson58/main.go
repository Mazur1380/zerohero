package main

import "fmt"

func main() {

	x := make([]int, 5, 6)
	copy(x, []int{1, 2, 3, 4, 5})
	priem(x)
	x = x[0:6]
	fmt.Println("Main", x)
	x = allo(x)
	fmt.Println("Main2", x)
}

func priem(x []int) {
	x = append(x, 6)
	fmt.Println("priem", x)
}
func allo(x []int) []int {
	x = append(x, 7)
	return x
}
