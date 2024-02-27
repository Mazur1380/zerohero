package main

import "fmt"

func main() {

	x := []int{1, 2, 3, 4, 5}

	for i := 0; i <= 4; i++ {
		y := x[len(x)-1]
		for i := len(x) - 2; i >= 0; i-- {
			x[i+1] = x[i]
		}
		x[0] = y
	}

	fmt.Println(x)
}
