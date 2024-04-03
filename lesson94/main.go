package main

import "fmt"

func main() {

	d := []int{1, 3, 45, 67, 89, 98, 100}
	v := binSearch(d, 68)
	fmt.Println(v)

}
func binSearch(x []int, y int) bool {
	if len(x) == 0 {
		return false
	}
	left := 0
	right := len(x) - 1

	for left <= right {
		m := (right-left)/2 + left
		if x[m] == y {
			return true
		}
		if x[m] < y {
			left = m + 1
		} else {
			right = m - 1
		}
	}
	return false
}
