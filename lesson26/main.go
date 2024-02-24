package main

import "fmt"

func main() {

	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//for i := 0; i < len(x); i++ {
	//fmt.Println(x[i])

	//}
	//for i := len(x) - 1; i >= 0; i-- {
	//fmt.Println(x[i])
	//}
	//for i := len(x) / 2; i < len(x); i++ {
	//	fmt.Println(x[i])
	//}
	for i := len(x) - 1; i >= len(x)/2; i-- {
		fmt.Println(x[i])
	}
}
