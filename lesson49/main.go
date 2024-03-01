package main

import "fmt"

func main() {
	f1 := func(c int) int {
		return c * 2
	}
	x := summ(10, f1)
	fmt.Println(x)
	f2 := func(c int) int {
		return c * 3
	}
	x = summ(10, f2)
	fmt.Println(x)
}
func summ(
	a int,
	f func(int) int,
) int {
	return f(a)

}
