package main

import "fmt"

func main() {
	x := tarelka{
		nabor: []string{"Конфеты", "Печенье", "Вафли"},
		price: 100,
	}
	y := tarelka{
		nabor: []string{"Конфеты", "Печенье", "Вафли", "Зефир"},
		price: 500,
	}
	fmt.Println(x, y)
	fmt.Println("В набор входит", x.nabor, "По цене", x.price)
	fmt.Println("В набор входит", y.nabor, "Входит", len(y.nabor))
	y.nabor = append(y.nabor, "Карамель")
	y.price = 1000
	fmt.Println(y)

}

type tarelka struct {
	nabor []string
	price int
}
