package main

import "fmt"

func main() {

	x := []int{1, 2, 3, 4, 5, 5, 5, 3, 65, 35, 777, 5}
	y := 0
	for i := 0; i < len(x); i++ {
		if x[i] == 5 {
			y = y + 1
		}
	}
	fmt.Println(y)
}
