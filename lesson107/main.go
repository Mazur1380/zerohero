package main

import (
	"fmt"
	"strings"
)

func main() {
	x := []string{"Dima", "Djas", "Vlad"}
	v := filter(x, func(e string) bool {
		return !strings.HasPrefix(e, "Di")
	})
	fmt.Println(v)

}

func filter[T comparable](x []T, f func(e T) bool) []T {
	v := make([]T, 0)

	for _, e := range x {
		if f(e) {
			v = append(v, e)
		}

	}
	return v
}
