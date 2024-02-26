package main

import "fmt"

func main() {

	a := []int{2, 4, 6}
	b := []int{1, 2, 3}
	d := []int{2, 3, 4}
	c := []int{}

	//c = append(c, 1, 2, 3)
	for i := 0; i < len(a); i++ {
		y := a[i] + b[i] + d[i]
		c = append(c, y)
	}
	fmt.Println(c)
}
