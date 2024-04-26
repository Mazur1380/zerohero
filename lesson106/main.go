package main

import "fmt"

func main() {
	x := sum(2, 4)
	fmt.Println(x)
}

func sum[T int | int32 | float64 | float32](a, b T) T {
	return a + b
}
