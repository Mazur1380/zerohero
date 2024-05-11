package main

import (
	"fmt"
	"time"
)

func main() {
	x := 0
	ch := make(chan struct{}, 1)
	ch <- struct{}{}
	go incTimes(&x, 100, ch)
	go incTimes(&x, 100, ch)
	go incTimes(&x, 100, ch)
	time.Sleep(3 * time.Second)

	printValue(&x, ch)

}

func incTimes(dest *int, times int, ch chan struct{}) {
	for i := 0; i < times; i++ {
		<-ch
		*dest = *dest + 1
		ch <- struct{}{}

	}
}

func printValue(dest *int, ch chan struct{}) {
	<-ch
	fmt.Println(*dest)
	ch <- struct{}{}
}
