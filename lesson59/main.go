package main

import "fmt"

func main() {

	x := []int{9, 6, 45, 26, 74, 26}
	fmt.Println(len(x), cap(x), x)
	b := x
	fmt.Println(len(b), cap(b), b)
	b[0] = 99
	fmt.Println(x)
	b = append(b, 18)
	fmt.Println("-------")
	fmt.Println(len(x), cap(x), x)
	fmt.Println(len(b), cap(b), b)
	b[0] = 100
	fmt.Println(x)
	fmt.Println(b)

}
