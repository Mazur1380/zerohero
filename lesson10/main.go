package main

import "fmt"

func main() {
	kal, ps5 := true, false
	var super bool
	super = kal && ps5

	if super {
		fmt.Println("мы берем тебя на работу")
	} else if kal {
		fmt.Println("Научим плейстейщн")
	} else if ps5 {
		fmt.Println("Мы научим кальяну")
	} else {
		fmt.Println("мы не берем тебя на работу")
	}
}
