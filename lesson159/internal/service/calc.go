package service

type Calc struct{}

func (c *Calc) Sum(a int, b int) int {
	return a + b
}

func (c *Calc) Sub(a int, b int) int {
	return a - b
}

func (c *Calc) Mul(a int, b int) int {
	return a * b
}

func (c *Calc) Div(a int, b int) int {
	return a / b
}
