package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	wg := &sync.WaitGroup{}
	wg.Add(3)
	go p(wg)
	go p(wg)
	go p(wg)
	wg.Wait()
	fmt.Println("MAIN DONE")

}

func p(wg *sync.WaitGroup) {
	for i := 1; i <= 3; i++ {
		fmt.Println(i)
		time.Sleep(time.Second)
	}
	wg.Done()
}
