package main

import "fmt"

func main() {
	fmt.Printf("%T %v\n", x, x)
	x := 1.5
	fmt.Printf("%T %v\n", x, x)
	var y = true
	fmt.Printf("%T %v\n", y, y)
	var t = "Hello"
	fmt.Printf("%T %v\n", t, t)
	var r [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Printf("%T %v\n", r, r)
	var e []string = []string{"Burn", "Fire"}
	fmt.Printf("%T %v\n", e, e)
	f := func() {}
	fmt.Printf("%T %v\n", f, f)
	b := struct {
		name string
	}{
		name: "Vlad",
	}
	fmt.Printf("%T %+v\n", b, b)

}

var x int = 10
