package main

import "fmt"

func main() {
	var a int
	a = 5
	fmt.Println(a)
	var b [3]string
	b = [3]string{"", "", ""}
	fmt.Println(b)
	var c struct {
		name string
		age  int
	}
	c = struct {
		name string
		age  int
	}{
		name: "vlad",
		age:  28,
	}
	fmt.Println(c)
	var d User
	d = User{
		name: "vlad",
		age:  28,
	}
	fmt.Println(d)
	var e func(int, int) int

}

type User struct {
	name string
	age  int
}
