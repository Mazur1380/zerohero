package main

import (
	"fmt"
)

func main() {

	f := "Hello"
	for _, c := range f {
		fmt.Println(string(c))
	}
	x := []byte(f)
	fmt.Println(x)

	b := f[0]
	fmt.Println(b, string(b))
	fmt.Println(len(f))

}
