package main

import "fmt"

func main() {
	hello()
	x := hello
	x()
	var y func()
	y = x
	y()
	c := func() {
		fmt.Println("War craft")
	}
	c()
	c = hello
	c()

}

func hello() {
	fmt.Println("Hello world")
}
