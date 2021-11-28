package beverage

type Beverage interface {
	cost(int) int
}
type BeverageFunc func(int) int

func (b BeverageFunc) cost(a int) int {
	price := a + b.cost()
	return b(price)
}
