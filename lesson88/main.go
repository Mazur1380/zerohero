package main

import "fmt"

func main() {

	x := []int{1, 2, 3, 4, 5}
	c := make([]int, 5)
	k := 0
	for i := len(x) - 1; i >= 0; i-- {
		c[k] = x[i]
		k = k + 1
	}
	fmt.Println(c)
}
