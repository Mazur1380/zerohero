package main

import (
	"fmt"
)

type storeTypes interface {
	int | string
}

type Storer[T storeTypes] interface {
	Add(e T)
	Len() int
}

type Store[T storeTypes] struct {
	data []T
}

func (s *Store[T]) Add(e T) {
	s.data = append(s.data, e)
}

func (s *Store[T]) Len() int {
	return len(s.data)
}

func main() {
	s := &Store[string]{}
	s.Add("Dima")
	s.Add("Djas")
	s.Add("Vlad")

	fmt.Println(s)
	fmt.Println(s.Len())

	var x Storer[string]
	x = s
	fmt.Println(x.Len())
}
