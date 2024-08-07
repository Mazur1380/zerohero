package calc

type Calc struct{}

func (c *Calc) Sum(a, b int) int {
	return a + b
}

func (m *Calc) Mul(a, b int) int {
	return a * b
}
