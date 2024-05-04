package main

import (
	"embed"
	"fmt"
)

//go:embed *.txt
var files embed.FS

func main() {
	h, _ := files.ReadFile("hello.txt")
	fmt.Println(string(h))
	s, _ := files.ReadFile("second.txt")
	fmt.Println(string(s))

}
