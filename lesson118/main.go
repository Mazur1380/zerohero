package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	go incTimes(10000, "jas")
	go incTimes(10000, "jas")
	time.Sleep(time.Second)
	mu.Lock()
	fmt.Println(m)
	mu.Unlock()

}

var m = make(map[string]int)
var mu sync.Mutex

func incTimes(times int, key string) {
	for i := 0; i < times; i++ {
		inc(key)
	}
}

func inc(key string) {
	mu.Lock()
	m[key]++
	mu.Unlock()
}
