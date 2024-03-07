package main

import "fmt"

func main() {

	x := "Hello world"
	y := []byte{}

	for i := len(x) - 1; i >= 0; i-- {
		y = append(y, x[i])
	}
	g := ""
	g = string(y)
	fmt.Println(g)
}
