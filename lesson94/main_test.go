package main

import (
	"math/rand"
	"testing"
)

func createInput() []int {
	x := make([]int, 100000000)
	for i := 0; i < len(x); i++ {
		x[i] = rand.Intn(1000000)
	}
	return x
}

var input = createInput()
var search = 899999

func BenchmarkBinSearch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		binSearch(input, search)
	}
}

func BenchmarkLinSearch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		linSearch(input, search)
	}
}
