package main

import "fmt"

func main() {

	var x hookah
	fmt.Printf("%+v\n", x)
	x.name = "Alpha hookah"
	x.cup = "Turka"
	x.height = 15
	x.premiun = true
	fmt.Printf("%+v\n", x)
	fmt.Println(x.name)

}

type hookah struct {
	name    string
	cup     string
	height  int
	premiun bool
}
