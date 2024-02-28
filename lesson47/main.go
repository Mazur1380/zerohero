package main

import "fmt"

func main() {
	//tor(0)
	fmt.Println(fact(100))
}

// func tor(a int) {
//
//	if a > 3 {
//		return
//	}
//
// fmt.Println(a)
// tor(a + 1)
// }
func fact(a int) int {
	if a == 0 {
		return 1
	}
	x := a * fact(a-1)
	return x
}
