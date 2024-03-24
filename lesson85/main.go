package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	x := []int{9, 3, 5, 4, 7, 3, 7, 5, 3, 8, 5}

	Sort(x)
	fmt.Println(x)
	y := []int{9, 3, 5, 4, 7, 3, 7, 5, 3, 8, 5}
	SortSTD(y)
	fmt.Println(y)

}
func Swop(x []int) {
	max := math.MinInt
	index := -1
	for i := 0; i < len(x); i++ {
		if x[i] > max {
			max = x[i]
			index = i
		}

	}
	x[0], x[index] = x[index], x[0]
}

func Sort(x []int) {
	for i := 0; i < len(x)-1; i++ {
		Swop(x[i:])
	}
}
func SortSTD(y []int) {
	sort.Slice(y, func(i int, j int) bool {
		return y[i] > y[j]
	})
}
