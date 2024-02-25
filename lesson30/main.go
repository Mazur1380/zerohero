package main

import (
	"fmt"

	"github.com/andreipimenov/zerohero/lesson30/hello"
)

func main() {

	fmt.Println(hello.A)
	hello.A = 20
	fmt.Println(hello.A)

}
