package main

import "fmt"

func main() {

	a := []string{"a", "b", "c", "d", "e", "f"}

	for i := 0; i < len(a)/2; i++ {
		n := len(a) - 1 - i
		fmt.Println(a[n])
	}
}
