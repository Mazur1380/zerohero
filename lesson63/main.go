package main

import (
	"fmt"
	"strings"
)

func main() {

	f := "hello,world,russia"
	x := strings.Split(f, ",")
	fmt.Println(x)

	d := []string{"Vlad", "World", "Hello"}
	g := strings.Join(d, "-")
	fmt.Println(g)

	c := "Hello world"
	j := strings.HasPrefix(c, "Hell")
	fmt.Println(j)

	j = strings.HasSuffix(c, "world")
	fmt.Println(j)

	c = "Hello pidori world"
	j = strings.Contains(c, "pidori")
	fmt.Println(j)
	c = strings.ReplaceAll(c, "pidori", "***")
	fmt.Println(c)

}
