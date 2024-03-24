package main

import (
	"fmt"
	"math"
)

func main() {
	x := []int{9, 3, 5, 4, 7, 3, 7, 5, 3, 8, 5}

	sort(x)
	fmt.Println(x)
}
func Swop(x []int) {
	min := math.MaxInt
	index := -1
	for i := 0; i < len(x); i++ {
		if x[i] < min {
			min = x[i]
			index = i
		}

	}
	x[0], x[index] = x[index], x[0]
}

func sort(x []int) {
	for i := 0; i < len(x)-1; i++ {
		Swop(x[i:])
	}

}
