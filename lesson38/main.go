package main

import "fmt"

func main() {

	x := []string{"a", "b", "c", "d", "a"}

	c := 0

	for _, i := range x {
		if i == "a" {
			c = c + 1
		}

	}
	fmt.Println(c)
}
