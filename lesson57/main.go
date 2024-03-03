package main

import (
	"fmt"
	"math"
)

func main() {
	var x []skhool
	x = []skhool{
		skhool{
			name:  "Jonh",
			score: 5,
		},
		skhool{
			name:  "Jane",
			score: 4,
		},
		skhool{
			name:  "Vlad",
			score: 3,
		},
	}
	max := math.MinInt
	index := -1

	for i, e := range x {
		if e.score > max {
			max = e.score
			index = i
		}
	}
	fmt.Println("Самая большая оценка", max, "у", x[index].name)
}

type skhool struct {
	name  string
	score int
}
