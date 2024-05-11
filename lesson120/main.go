package main

import (
	"sync"
	"time"
)

func main() {
	m := &sync.Map{}
	go setTimes(m, "John", 100)
	go setTimes(m, "John", 100)
	go setTimes(m, "John", 100)
	time.Sleep(time.Second)

}

func setTimes(m *sync.Map, key string, times int) {
	for i := 0; i < times; i++ {
		m.Store(key, "1")
	}
}
