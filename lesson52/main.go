package main

import "fmt"

func main() {

	x := []string{"Vlad", "Djas", "Ulfa", "Dima"}
	i := 0
	y := func() string {
		if i > len(x)-1 {
			return ""
		}
		p := x[i]
		i++
		return p

	}
	//fmt.Println(y())
	//fmt.Println(y())
	//fmt.Println(y())
	//fmt.Println(y())
	//fmt.Println(y())
	//fmt.Println(y())
	for {
		d := y()
		if d == "" {
			break
		}
		fmt.Println(d)
	}
}
