package main

import "fmt"

func main() {

	FizzBuzz(100, 1000)
}

func FizzBuzz(a, b int) {
	for i := a; i <= b; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			fmt.Println(i, "FizzBuzz")
		case i%3 == 0:
			fmt.Println(i, "Fizz")
		case i%5 == 0:
			fmt.Println(i, "Buzz")
		}

	}

}
