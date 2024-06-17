package main

import (
	"fmt"
	"strings"
)

func main() {

	x := []string{"Пн", "Вт", "Ср", "Чт", "Пт", "Сб", "Вс"}
	g := make([]string, 5)
	copy(g, x[:5])
	d := strings.Join(g, ", ")
	fmt.Printf("Рабочие дни: %s\n", d)
	x = append(x[5:6], x[6:]...)
	p := strings.Join(x, ", ")
	fmt.Printf("Выходные дни: %s\n", p)
}
