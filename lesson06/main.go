package main

import "fmt"

func main() {
	var a, b, c int
	a = 5
	b = 6
	c = 7
	var dock int
	dock = sum(a, b, c)
	fmt.Println(dock)
}

func sum(a, b, c int) int {
	var sum int
	sum = a + b + c
	return sum
}
