package main

import "fmt"

func main() {

	y := fact(10)

	fmt.Println(y)
	fmt.Println(fact(0))
}
func fact(a int) int {
	y := 1
	for i := 1; i <= a; i++ {
		y = y * i

	}
	return y

}
