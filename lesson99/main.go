package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {

	numbers := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
	leters := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
	symbols := []string{"!", "@", "#", "$", "%", "^", "&", "*", "(", "_"}

	randomValues := getRandomValues(numbers, leters, symbols)
	fmt.Println(randomValues)

}

func getRandomValues(numbers, leters, symbols []string) string {
	rand.Seed(time.Now().UnixNano())
	randomValues := make([]string, 6)

	for i := 0; i < 6; i++ {
		randIndex := rand.Intn(3)
		switch randIndex {
		case 0:
			randomValues[i] = numbers[rand.Intn(len(numbers))]
		case 1:
			randomValues[i] = leters[rand.Intn(len(leters))]
		case 2:
			randomValues[i] = symbols[rand.Intn(len(symbols))]

		}

	}
	return strings.Join(randomValues, "")

}
