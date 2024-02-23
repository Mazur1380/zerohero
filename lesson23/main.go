package main

import "fmt"

func main() {

	x := map[string]int{
		"dima": 0,
		"vlad": 1,
		"djas": 2,
	}
	sum := 0

	for _, w := range x {
		sum = sum + w
	}
	fmt.Println(sum)

	for w := range x {
		fmt.Println(w)
	}

}
