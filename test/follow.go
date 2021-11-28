package main

func main() {
	var beverage Beverage
	coffee := beverage.Decorate(beverage.Espresso{}, beverage.plusMilk())
}

type Beverage interface {
	cost() int
}
type BeverageFunc func() int

func (b BeverageFunc) cost() int {
	return b()
}

//------------------------------------------------
type Espresso struct {
	beverage Beverage
}

func (e Espresso) cost() int {
	return 2000
}

//------------------------------------------------
type CondimentDecorator func(Beverage) Beverage

//------------------------------------------------
func Decorate(b Beverage, ds ...CondimentDecorator) Beverage {
	decorated := b
	for _, decorate := range ds {
		decorated = decorate(decorated)
	}
	return decorated
}

func plusMilk() CondimentDecorator {
	return func(b Beverage) Beverage {
		return BeverageFunc(func() int {
			price := b.cost() + 400
			return b.cost()
		})
	}
}
