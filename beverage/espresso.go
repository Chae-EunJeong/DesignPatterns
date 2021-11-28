package beverage

type Espresso struct {
	beverage Beverage
}

func (e Espresso) cost() int {
	return 2000
}
