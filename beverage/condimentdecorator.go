package beverage

type CondimentDecorator func(Beverage) Beverage

func Decorate(b Beverage, ds ...CondimentDecorator) Beverage {
	decorated := b
	for _, decorate := range ds {
		decorated = decorate(decorated)
	}
	return decorated
}
