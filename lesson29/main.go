package main

import "fmt"

func main() {

	x := []int{1, 2, 3, 4, 5}
	for i := 0; i <= len(x)-1; i++ {
		x[i] = x[i] * 2
	}
	fmt.Println(x)
}
