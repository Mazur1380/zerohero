package main

import "fmt"

func main() {
	x := 7
	fmt.Println(x, &x)
	y := &x
	fmt.Printf("%T %v\n", y, y)
	fmt.Println(*y)
	*y = 10
	fmt.Println(x)
}
