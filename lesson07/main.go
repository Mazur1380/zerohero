package main

import "fmt"

func main() {
	// fmt.Println(lol(10))
	var a int
	a = 10
	a = 10 - 2
	a = a * 2
	// a = a + 1
	// a++
	a += 1
	fmt.Println(a)

}

func mul(x int) int {
	return x * 2
}

func div(y int) int {
	return y / 2
}

func lol(f int) int {
	var pic int
	pic = mul(f)

	var sum int
	sum = div(pic)

	return sum
}
