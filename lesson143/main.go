package main

import (
	"fmt"

	"github.com/andreipimenov/zerohero/lesson143/list"
)

func main() {

	l := list.New()
	l.Push("Djas")
	l.Push("Vlad")
	l.Push("Dima")
	l.Print()
	l.PushMany("Ulfa", "Alex")
	l.Print()

	name := l.Pop()
	fmt.Println("Element", name)
	l.Print()

	name = l.Pop()
	fmt.Println("Element", name)
	l.Print()
	l.Delete("Ulfa")
	l.Print()
	fmt.Println(l.Len())
	fmt.Println(l.Exist("Alex"))

}
