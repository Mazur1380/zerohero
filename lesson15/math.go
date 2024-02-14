package main

func avg(j [5]float64) float64 {
	var a, b, c, d, f int = 0, 1, 2, 3, 4

	var sum float64
	sum = sum + j[a]
	sum = sum + j[b]
	sum = sum + j[c]
	sum = sum + j[d]
	sum = sum + j[f]
	sum = sum / 5
	return sum

}
