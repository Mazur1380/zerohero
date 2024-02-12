package main

import "fmt"

var (
	a int
	b int
)

func main() {
	a, b = 3, 5
	sum()
}

func sum() {
	fmt.Println(a + b)
}
