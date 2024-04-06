package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func getInt(envVar string) (int, error) {
	v := os.Getenv(envVar)
	if v == "" {
		return 0, nil
	}
	x, err := strconv.Atoi(v)
	if err != nil {
		return 0, err
	}
	return x, nil
}

func main() {
	a, err := getInt("A")
	if err != nil {
		log.Fatal(err)
	}

	b, err := getInt("B")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(a + b)
}
