package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	wg := &sync.WaitGroup{}
	wg.Add(3)
	one := make(chan struct{}, 1)
	two := make(chan struct{}, 1)
	go p(wg, "one", nil, one)
	go p(wg, "two", one, two)
	go p(wg, "three", two, nil)
	wg.Wait()
	fmt.Println("MAIN DONE")
}

func p(wg *sync.WaitGroup, name string, from, to chan struct{}) {
	if from != nil {
		<-from

	}
	for i := 1; i <= 3; i++ {
		fmt.Println(name, i)
		time.Sleep(time.Second)
	}
	if to != nil {
		to <- struct{}{}
	}
	wg.Done()
}
