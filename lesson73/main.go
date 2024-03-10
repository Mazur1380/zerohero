package main

import "fmt"

func main() {
	x := cash(9050)
	fmt.Println(x.Balance())
	fmt.Println(x.BalanceDc(90.5))
}

type cash int

func (c cash) RubCop() (int, int) {
	rub := c / 100
	cop := c % 100
	return int(rub), int(cop)

}

func (c cash) Balance() string {
	rub, cop := c.RubCop()
	s := fmt.Sprint(rub, " Рублей ", cop, " Копеек ")
	return s
}
func (c cash) Dc(ex float64) (int, int) {
	f := float64(c) / ex
	dolar := int(f) / 100
	cent := int(f) % 100
	return dolar, cent
}
func (c cash) BalanceDc(ex float64) string {
	dol, cen := c.Dc(ex)
	v := fmt.Sprint(dol, " Доларов ", cen, " Центов ")
	return v

}
