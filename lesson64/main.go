package main

import (
	"fmt"
)

func main() {

	x := "Привет мир ✅"
	for _, c := range x {
		fmt.Println(string(c))
	}
	fmt.Println(len(x))
	a := []rune(x)
	fmt.Println(len(a))

}
