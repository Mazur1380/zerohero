package main

import "fmt"

func main() {
	var work [7]float64
	work[0] = 1
	work[1] = 4.5
	work[2] = 5.5
	work[3] = 6.5
	work[4] = 7.5

	fmt.Println(work)
	var last int
	last = len(work) - 1
	work[last] = -100
	if work[0] < work[last] {
		work[0], work[last] = work[last], work[0]
	}
	fmt.Println(work)
	fmt.Println(len(work))

}
