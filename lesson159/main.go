package main

import (
	"errors"
	"log"

	"github.com/andreipimenov/zerohero/lesson159/internal/service"
)

func main() {

	c := &service.Calc{}
	nums := [2]int{600, 500}
	err := checkBalance(c, nums)
	if err != nil {
		log.Fatal(err)
	}

}

func checkBalance(c calc, nums [2]int) error {
	result := c.Sub(nums[0], nums[1])
	if result < 0 {
		return errors.New("not enough balance")
	}
	return nil
}

type calc interface {
	Sub(a int, b int) int
}
