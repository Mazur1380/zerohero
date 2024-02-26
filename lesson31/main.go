package main

import "fmt"

func main() {

	x := []string{"petr", "fedor", "vlad", "ivan"}
	f := false
	for i := 0; i < len(x); i++ {
		if x[i] == "ivan" {
			f = true
		}

	}
	if f {
		fmt.Println("Ivan est")

	} else {
		fmt.Println("Ivana net")
	}
}
