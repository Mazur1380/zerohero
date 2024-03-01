package main

import "fmt"

func main() {

	y := hello()
	fmt.Println(y())
	fmt.Println(y())
	fmt.Println(y())

}
func hello() func() int {
	x := 2
	return func() int {
		x = x * 2
		return x

	}
}
