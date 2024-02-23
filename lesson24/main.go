package main

import "fmt"

func main() {

	x := map[string]int{
		"Vlados": 1,
		"Dima":   2,
		"Djas":   3,
		"Andrey": 4,
		"Jr":     5,
	}
	sum2 := 0
	sum := 0

	for _, t := range x {
		if t > 3 {

			if t != 5 {
				sum = sum + t
			}

		} else {
			sum2 = sum2 + t
		}
	}
	fmt.Println(sum, sum2)

}
