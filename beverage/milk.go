package beverage

func plusMilk() CondimentDecorator {
	return func(b Beverage) Beverage {
		return BeverageFunc(func() int {
			price := b.cost() + 400
			return price
		})
	}
}
