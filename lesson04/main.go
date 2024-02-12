package main

import "fmt"

func main() {
	var a, b int
	a = 5
	b = 7
	sum(a, b)
}

func sum(x int, y int) {
	var res int
	res = x + y
	fmt.Println(res)
}
