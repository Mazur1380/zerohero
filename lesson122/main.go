package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	x := 0
	mu := &sync.RWMutex{}
	go incTimes(&x, 100, mu)
	go incTimes(&x, 100, mu)
	go incTimes(&x, 100, mu)
	go printValue(&x, mu)
	go printValue(&x, mu)
	go printValue(&x, mu)
	time.Sleep(3 * time.Second)

}

func incTimes(dest *int, times int, mu *sync.RWMutex) {
	for i := 0; i < times; i++ {
		mu.Lock()
		*dest = *dest + 1
		mu.Unlock()

	}
}

func printValue(dest *int, mu *sync.RWMutex) {
	for {
		mu.RLock()
		fmt.Println(*dest)
		mu.RUnlock()
	}
}
