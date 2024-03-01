package main

import "fmt"

func main() {
	y := hello()
	x := y()
	fmt.Println(x)
	fmt.Println(y())
	fmt.Println(y())
	fmt.Println(y())

}
func hello() func() int {
	n := 0
	f := func() int {
		n++
		return n
	}
	return f
}
