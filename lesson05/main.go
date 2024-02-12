package main

import "fmt"

func main() {
	var a, b, c int
	a = 5
	b = 9
	c = mul(a, b)
	fmt.Println(c)
	var d int
	d = mul(4, 5)
	fmt.Println(d)
	fmt.Println(mul(6, 9))
}

func mul(x, y int) int {
	var res int
	res = x * y
	return res

}
