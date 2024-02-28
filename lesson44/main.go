package main

import "fmt"

func main() {

	x := []int{12000, 15000, 25000, -40000, 45000}

	sum := 0
	for i := 0; i < len(x)-1; i++ {
		y := x[i+1] - x[i]
		if y < 0 {
			y = -y
		}
		if y > sum {
			sum = y
		}

	}
	fmt.Println(sum)
}
