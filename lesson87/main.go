package main

import "fmt"

func main() {
	t := []int{5, 3, 9, 4, 4, 4}
	n := Exist(t, 4)
	if n > 0 {
		fmt.Println("Done", n)
	} else {
		fmt.Println("Not found")
	}

}

func Exist(x []int, y int) int {
	counter := 0
	for i := range x {
		if x[i] == y {
			counter = counter + 1
		}
	}
	return counter

}
