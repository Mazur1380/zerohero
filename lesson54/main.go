package main

import "fmt"

func main() {

	type user struct {
		age  int
		name string
	}
	var u user
	u.age = 18
	u.name = "vlad"
	fmt.Println(u)
}
