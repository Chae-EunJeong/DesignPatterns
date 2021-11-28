package client

type Decorator func(Client) Client

// Decorate will decorate a client with a slice of passed decorators
func Decorate(c Client, ds ...Decorator) Client {
	decorated := c
	for _, decorate := range ds {
		decorated = decorate(decorated)
	}
	return decorated
}
