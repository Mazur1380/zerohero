package main

import "fmt"

func main() {
	s := ""
	fmt.Scanln(&s)
	check(s)
}

func check(s string) {
	switch s {
	case "light":
		fmt.Println("light hookah")
	case "medium":
		fmt.Println("medium hookah")
	case "strong":
		fmt.Println("strong hookah")
	default:
		fmt.Println("no edition")
	}
}
