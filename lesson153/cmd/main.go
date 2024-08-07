package main

import (
	"fmt"

	"github.com/andreipimenov/zerohero/lesson153/internal/calc"
	"github.com/andreipimenov/zerohero/lesson153/internal/service"
)

func main() {
	c := &calc.Calc{}
	s := service.New(c)

	res := s.SumMany(1, 2, 3)
	fmt.Println(res)

}
