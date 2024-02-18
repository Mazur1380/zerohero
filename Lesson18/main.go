package main

import "fmt"

var x int

func main() {

	fmt.Println(x)
	x = 5
	fmt.Println(x)
	var x int
	fmt.Println(x)

	{
		var x int = 15
		fmt.Println(x)

	}
	fmt.Println(x)
}
