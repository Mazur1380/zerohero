package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	x := []int{1, 2, 3}
	mu := &sync.Mutex{}
	go incTimes(x, 0, 1000, mu)
	go incTimes(x, 0, 1000, mu)
	go incTimes(x, 0, 1000, mu)
	go PrintTimes(x, 0, mu)
	time.Sleep(5 * time.Second)
}

func incTimes(x []int, index int, times int, mu *sync.Mutex) {
	for i := 0; i < times; i++ {
		mu.Lock()
		x[index]++
		mu.Unlock()
	}
}

func PrintTimes(x []int, index int, mu *sync.Mutex) {
	for {
		mu.Lock()
		fmt.Println(x[index])
		mu.Unlock()
	}
}
