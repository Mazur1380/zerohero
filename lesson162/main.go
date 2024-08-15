package main

import (
	"errors"
	"fmt"
)

func main() {
	err := NewMyError("user not found")
	err2 := fmt.Errorf("service error: %w", err)
	fmt.Println(err2)

	var myErr *MyError
	ok := errors.As(err2, &myErr)
	if ok {
		fmt.Printf("found our error: %v", myErr)
	} else {
		fmt.Println("my error not found")
	}

}

type MyError struct {
	message string
}

func (m *MyError) Error() string {
	return m.message
}

func NewMyError(m string) *MyError {
	return &MyError{
		message: m,
	}
}
