package main

func main() {

}

func summ(x []int) int {
	g := 0
	for _, i := range x {
		g = g + i
	}
	return g
}
