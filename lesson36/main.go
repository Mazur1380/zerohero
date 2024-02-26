package main

import "fmt"

func main() {

	x := []int{1, 2, 3, 4, 5, 9, 4, 3, 2, 1}

	max := -1

	for _, item := range x {
		if item > max {
			max = item
		}
	}
	fmt.Println(max)
}
