package main

import (
	"fmt"
	"math"
)

func main() {

	x := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}

	cal := math.MaxInt

	for i := 0; i < len(x); i++ {
		if x[i] < cal {
			cal = x[i]
		}
	}
	fmt.Println(cal)
}
