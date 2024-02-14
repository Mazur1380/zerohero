package main

import "fmt"

func main() {
	var god [5]float64
	god[0] = 1
	god[1] = 2
	god[2] = 3
	god[3] = 4
	god[4] = 5
	fmt.Println(god)
	var a float64
	a = avg(god)
	fmt.Println(a)
}
