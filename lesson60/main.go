package main

import "fmt"

func main() {
	x := make([]int, 6, 12)
	fmt.Println(len(x), cap(x), x)
	x = append(x, 9)
	fmt.Println(len(x), cap(x), x)
	b := x
	b = append(b, 12)
	b[0] = 25
	fmt.Println(len(b), cap(b), b)
	fmt.Println(len(x), cap(x), x)
	x = append(x, 34)
	fmt.Println("------")
	fmt.Println(len(b), cap(b), b)
	fmt.Println(len(x), cap(x), x)
}
