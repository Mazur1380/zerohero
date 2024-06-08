package main

import "fmt"

func main() {
	var av AmericanVelocity = 120.4

	ev := convert(av)

	fmt.Println(av, ev)
	ev = convert(120.4)
	fmt.Println(ev)
}

type AmericanVelocity float64
type EuropeanVelocity float64

func convert(av AmericanVelocity) EuropeanVelocity {
	ev := EuropeanVelocity(av * 5793.624)
	return ev
}
