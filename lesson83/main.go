package main

import "fmt"

func main() {

	// x := []int{
	// 	48, 96, 86, 68,
	// 	57, 82, 63, 70,
	// 	37, 34, 83, 27,
	// 	19, 97, 9, 17,
	// }
	// min := x[0]

	// for i := 1; i < len(x); i++ {
	// 	if x[i] < min {
	// 		min = x[i]
	// 	}
	// }
	// fmt.Println(min)
	e := sum([]int{1, 2, 3, 4, 5})
	fmt.Println(e)

}
func sum(i []int) int {
	x := 0
	for y := range i {
		x = x + i[y]
	}
	return x
}
