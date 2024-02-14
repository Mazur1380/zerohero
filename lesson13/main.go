package main

import "fmt"

func main() {
	var a, b, c float64
	var op string
	a, b, op = fom()

	//a := 0
	//b := 0
	//c := 0
	//var a, b, c float64

	// if op == "+" {
	// 	c = a + b
	// } else if op == "-" {
	// 	c = a - b
	// } else if op == "*" {
	// 	c = a * b
	// } else if op == "/" {
	// 	c = a / b
	// } else {
	// 	fmt.Println("Не поддерживаемая операция")
	// 	return
	// }
	switch op {
	case "+":
		c = a + b
	case "-":
		c = a - b
	case "*":
		c = a * b
	case "/":
		c = a / b
	default:
		fmt.Println("Не поддерживаемая операция")
		return
	}
	fmt.Printf("%v %v %v = %.2f\n", a, op, b, c)

}
func fom() (float64, float64, string) {
	a, b := 0.0, 0.0
	fmt.Println("Введите первое число")
	fmt.Scanln(&a)
	fmt.Println("Укажите математическую операцию")
	op := ""
	fmt.Scanln(&op)
	fmt.Println("Введите второе число")
	fmt.Scanln(&b)
	return a, b, op

}
