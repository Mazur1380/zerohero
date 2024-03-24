package main

import "fmt"

func main() {
	t := []int{5, 3, 9, 4, 8, 2}
	n := Exist(t, 1)
	if n {
		fmt.Println("Done")
	} else {
		fmt.Println("Not found")
	}

}

func Exist(x []int, y int) bool {
	for i := range x {
		if x[i] == y {
			return true
		}
	}
	return false
}
