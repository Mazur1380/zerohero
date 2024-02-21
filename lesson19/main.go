package main

import "fmt"

func main() {
	c := [5]int{1, 2, 3, 4, 5}
	fmt.Println(c)
	c[0] = 2
	fmt.Println(c)
	fmt.Println(c[2])
	c = [5]int{3, 4, 5, 6, 7}
	fmt.Println(c)
	c = [5]int{}
	fmt.Println(c)
	l := len(c)
	fmt.Println(l)
	fmt.Println([4]string{"hello\n", "buy\n", "Привет\n", "Пока\n"})
	x := [5]string{
		"hello",
		"world",
		"buy",
		"Promes",
		"Spartak",
	}
	fmt.Println(x)
}
