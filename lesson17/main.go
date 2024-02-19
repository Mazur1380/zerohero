package main

import "fmt"

func main() {

	x := 1
	s := "Четное"
	if x%2 == 1 {
		s = "Не четное"

	}
	f := "Число %v - %v\n"

	fmt.Printf(f, x, s)

}
