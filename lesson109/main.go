package main

import "fmt"

func main() {

	v := "Hello 🌍"
	for i, c := range v {
		fmt.Printf("%d of '%c'\n", i, c)
	}
}
