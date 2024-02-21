package main

import "fmt"

func main() {

	c := [5]float64{
		8, 6, 4, 7, 9,
	}
	sum := 0.0

	for i := 0; i < 5; i++ {
		//fmt.Println(c[i])
		sum = sum + c[i]

	}
	fmt.Println(sum)

	sum = 0

	for i := range c {
		sum = sum + c[i]
	}

	fmt.Println(sum)

	sum = 0

	for _, element := range c {
		sum = sum + element

	}
	fmt.Println(sum)
}
