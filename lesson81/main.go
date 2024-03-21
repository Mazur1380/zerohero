package main

import "fmt"

func main() {
	g := 17
	s := sayHello("Vlados", g)
	fmt.Println(s)
	f, e := propusk(g)
	fmt.Println(f, e)

}

func sayHello(name string, age int) string {
	return fmt.Sprintf("Привет, %s! Тебе %d лет", name, age)
}

func propusk(age int) (string, bool) {
	if age >= 18 {
		return "Входи", true
	} else {
		return "Не проходишь", false
	}
}
