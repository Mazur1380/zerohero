package service

import (
	"github.com/andreipimenov/zerohero/lesson153/internal/calc"
)

type Service struct {
	calc *calc.Calc
}

func New(calc *calc.Calc) *Service {
	return &Service{
		calc: calc,
	}
}

func (s *Service) SumMany(numbers ...int) int {
	sum := 0
	for _, number := range numbers {
		sum = s.calc.Sum(sum, number)
	}
	return sum

}
