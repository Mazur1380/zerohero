package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

var x int64

func main() {
	m := map[string]int{}
	go incTimes(m, 100)
	go incTimes(m, 100)
	go incTimes(m, 100)
	go printValue(m)
	go printValue(m)
	go printValue(m)
	time.Sleep(3 * time.Second)

}

func incTimes(m map[string]int, times int) {
	for i := 0; i < times; i++ {
		for {
			swapped := atomic.CompareAndSwapInt64(&x, 0, 1)
			if swapped {
				break
			}
		}
		m["jonh"]++
		atomic.StoreInt64(&x, 0)
	}
}

func printValue(m map[string]int) {
	for {
		for {
			swapped := atomic.CompareAndSwapInt64(&x, 0, 1)
			if swapped {
				break
			}
		}
		fmt.Println(m["jonh"])
		atomic.StoreInt64(&x, 0)
	}
}
