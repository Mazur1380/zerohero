package main

import "fmt"

func main() {
	var kal, ps5 = false, false
	switch {
	case kal && ps5:
		fmt.Println("ok")
	case kal:
		fmt.Println("ok2")
	case ps5:
		fmt.Println("ok3")
	default:
		fmt.Println("no ok")
	}
}
