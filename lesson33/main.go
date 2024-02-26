package main

import "fmt"

func main() {

	x := []int{0, 1, 2, -3, 4, -5}
	s := 0
	c := 0
	for i := 0; i < len(x); i++ {
		if x[i] < 0 {
			s = s + 1
		}
		if x[i] >= 0 {
			c = c + 1
		}
	}
	fmt.Println(s, c)
}
