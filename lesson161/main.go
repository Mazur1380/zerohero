package main

import (
	"errors"
	"fmt"
)

var ErrorNotFound = errors.New("not found")

func main() {

	err := errors.New("user not found")
	err2 := fmt.Errorf("failed to get user: %w", err)
	fmt.Println(err2)

	err3 := errors.New("error3")
	err4 := errors.New("error4")

	err5 := errors.Join(err3, err4)

	fmt.Println(err5)

	err6 := fmt.Errorf("not found: %w", ErrorNotFound)
	fmt.Println(err6)
}
