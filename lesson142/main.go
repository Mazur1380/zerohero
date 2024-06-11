package main

import "fmt"

func main() {

	var a *int
	var b int = 100

	a = &b

	fmt.Println(*a)

	*a = 200

	fmt.Println(b)
}
