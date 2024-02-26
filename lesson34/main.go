package main

import "fmt"

func main() {

	x := []int{0, 1, 2, -3, -4, 5, -6}

	s := 0
	c := 0

	for _, y := range x {
		if y < 0 {
			s = s + 1
		}
		if y >= 0 {
			c = c + 1
		}

	}
	fmt.Println(s, c)
}
