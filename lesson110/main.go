package main

import (
	"fmt"

	"github.com/Mazur1380/vlad1380/v3/math"
)

func main() {

	x := math.Sum(5, 3, 4)
	fmt.Println(x)
	fmt.Println(math.Mul(10, 5))
	fmt.Println(math.Sum(5, 10, 15))
	v := math.Sum(1.5, 3.4, 5.2, 4.7, 8.5, 3, 6)
	fmt.Println(v)
	m := math.Div[float64](100, 3)
	fmt.Println(m)

}
