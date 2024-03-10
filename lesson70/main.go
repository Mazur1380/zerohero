package main

import "fmt"

func main() {
	c := cash(100)
	fmt.Printf("%T %v\n", c, c)
	f := []byte(d)
	fmt.Printf("%T %v\n", f, f)
}

type cash int

var d string = "Hello"
