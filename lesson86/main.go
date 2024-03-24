package main

import "fmt"

func main() {
	a := []int{2, 3, 4, 6, 7}
	a = append(a, 0)
	copy(a[4:], a[3:5])
	a[3] = 5
	fmt.Println(a)
	// copy(b, a[:3])
	// b[3] = 5
	// copy(b[4:], a[3:])
	// fmt.Println(b)

	// c := make([]int, len(a)+len(b))

	// copy(c, a)
	// copy(c[len(a):], b)
	// fmt.Println(c)
}
