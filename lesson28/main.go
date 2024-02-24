package main

import "fmt"

func main() {

	x := []string{"a", "b", "c", "d", "e", "f", "g"}
	//x[0] = x[1]
	//n := x[0]
	//x[0] = x[1]
	//x[1] = n
	//x[0], x[1] = x[1], x[0]
	last := len(x) - 1
	//x[0], x[last] = x[last], x[0]
	//x[1], x[last-1] = x[last-1], x[1]
	//x[2], x[last-2] = x[last-2], x[2]
	for i := 0; i < len(x)/2; i++ {
		x[i], x[last-i] = x[last-i], x[i]
	}
	fmt.Println(x)

}
