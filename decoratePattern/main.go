package main

import (
	"log"
	"net/http"
	"os"

	"github.com/JeongChaeEun/Design_Patterns_in_Golang/decoratePattern/client"
)

// We can construct a client with all the functionality we need.
func main() {
	c := client.Decorate(client.ProxyTimeoutClient{},
		client.Logging(log.New(os.Stdout, " Logging >> ", 0)),
		client.Authorization("secretpassword"),
	)
	req, _ := http.NewRequest("GET", "http://example.com", nil)
	c.Do(req)
}
