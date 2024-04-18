package main

import (
	"flag"
	"fmt"
)

func main() {

	var num1 int
	var num2 int
	var operation string

	flag.IntVar(&num1, "num1", 0, "Первое чилсо")
	flag.IntVar(&num2, "num2", 0, "Второе число")
	flag.StringVar(&operation, "operation", "+", "Математическая операция")

	flag.Parse()

	result := 0
	switch operation {
	case "+":
		result = num1 + num2

	}

	fmt.Printf("Результат операции %v %v %v = %v\n", num1, operation, num2, result)

}
