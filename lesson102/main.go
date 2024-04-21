package main

import (
	"fmt"
	"time"
)

func main() {
	x := createInput()

	start := time.Now()

	s := sumV2(x, 100)
	d := time.Since(start)
	fmt.Println(s)
	fmt.Println(d)

}

func createInput() []int {
	x := make([]int, 100_000_000)
	for i := 0; i < len(x); i++ {
		x[i] = 1
	}
	return x

}

func summ(x []int) int {
	s := 0
	for i := 0; i < len(x); i++ {
		s = s + x[i]
	}
	return s
}
func sumV2(x []int, numGoroutines int) int {
	ch := make(chan int)

	for i := 0; i < numGoroutines; i++ {
		go sumChunk(x, i*1_000_000, (i+1)*1_000_000, ch)
	}

	s := 0
	for i := 0; i < numGoroutines; i++ {
		s = s + <-ch
	}
	return s
}

func sumChunk(x []int, start int, end int, c chan int) {
	s := 0
	for i := start; i < end; i++ {
		s += x[i]
	}
	c <- s
}
