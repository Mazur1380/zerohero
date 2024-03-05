package main

import "fmt"

func main() {

	var a string
	a = "Hello world"
	fmt.Println(a)

	var b string
	b = "hello world\nPrivet sobaki"
	fmt.Println(b)

	a = "Goal"
	b = "Spartak"
	c := a + b
	fmt.Println(c)
	f := `Привет мир
Как ваши нечего
У евс тоже все ок`
	fmt.Println(f)

}
