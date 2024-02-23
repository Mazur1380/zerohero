package main

import "fmt"

func main() {

	m := make(map[string]int)
	m["Jonh"] = 5
	m["Jane"] = 5
	m["Djas"] = 2

	fmt.Println(m)

	x := map[string]int{
		"Djas": 2,
		"Jane": 5,
		"Jonh": 0,
	}
	fmt.Println(x)
	fmt.Println(len(x))
	delete(x, "Djas")
	fmt.Println(len(x))
	fmt.Println(x)
	fmt.Println(x["Jane"])
	fmt.Println(x["xxx"])
	y, ok := x["Jonh"]
	fmt.Println(y, ok)
	t, ok := x["Dima"]
	fmt.Println(t, ok)

}
