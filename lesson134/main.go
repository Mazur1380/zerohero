package main

import "fmt"

func main() {

	s := []int{100, 200, 300, 400, 500}

	sum := 0

	for _, i := range s {
		sum = sum + i
	}

	avg := sum / len(s)

	sum = 0

	for _, i := range s {
		if i > avg {
			sum = sum + i
		}
	}
	fmt.Println(sum)
}
