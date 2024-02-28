package main

import "fmt"

func main() {
	hello()
	hello()
	number(5)
	sum(10, 20)
	var w int
	w = 10
	sum(w, w)
	x := mull(w, w)
	fmt.Println(x)
	fmt.Println(mull(w, 8))
	d := 10
	a := 4
	d, a = swap(d, a)
	fmt.Println(d, a)
	x, y := swap(5, 8)
	fmt.Println(x, y)
	fmt.Println(swap(3, 6))
	j := []int{4, 2, 7, 5}
	add(j, 20)
	pr(2, 3, 4, 4, 5, 5, 6, 6, 44, 3)
	s := summ(1, 4, 5, 3, 5, 4, 3, 4, 5)
	fmt.Println(s)
	v := exist(j, 0)
	fmt.Println(v)

}
func hello() {
	fmt.Println("HELLO WORLD")
}
func number(a int) {
	fmt.Println(a)
}
func sum(a, b int) {
	fmt.Println(a + b)
}
func mull(c, b int) int {
	d := c * b
	return d
}
func swap(c, b int) (int, int) {
	return b, c
}
func add(d []int, a int) {
	d = append(d, a)
	fmt.Println(d)
}
func pr(a ...int) {
	fmt.Println(a)
}
func summ(a ...int) int {
	t := 0
	for _, i := range a {
		t = t + i
	}
	return t
}
func exist(a []int, b int) bool {
	for _, i := range a {
		if b == i {
			return true
		}
	}
	return false
}
