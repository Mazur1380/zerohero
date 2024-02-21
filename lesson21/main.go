package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5, 6}
	sum := 0

	for _, y := range x {
		sum = sum + y

	}
	fmt.Println(sum)

}
