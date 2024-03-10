package main

import "fmt"

func main() {

	x := "Hello"
	fmt.Println(x[1])
	y := []byte(x)
	fmt.Println(y)
	y[1] = 'a'
	fmt.Println(y)
	x = string(y)
	fmt.Println(x)

	r := []float64{10, 15}
	sum := 0.0
	lenn := len(r)
	for _, i := range r {
		sum = sum + i
	}
	avg := sum / float64(lenn)
	fmt.Println(avg)

}
