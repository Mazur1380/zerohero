package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	var Myint int = 104
	var MyIntString string = strconv.Itoa(Myint)

	var Mystring string = "104"
	var MyStringInt, err = strconv.Atoi(Mystring)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Первый тип переменной %T, значения %s\n", MyIntString, MyIntString)
	fmt.Printf("Первый тип переменной %T, значения %d\n", MyStringInt, MyStringInt)
}
