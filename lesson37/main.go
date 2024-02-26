package main

import (
	"fmt"
	"math"
)

func main() {
	x := []int{5, 6, 7, 8, 9}

	min := math.MaxInt

	for _, y := range x {
		if min > y {
			min = y

		}
	}
	fmt.Println(min)
}
