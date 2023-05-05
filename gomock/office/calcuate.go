package office

type Calculator interface {
	Add(a, b int) int
	Subtract(a, b int) int
}

func DoubleCalculator(c Calculator, a, b int) int {
	x := c.Add(a, b)
	y := c.Subtract(a, b)
	return x + y
}
