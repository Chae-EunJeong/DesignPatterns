package client

import (
	"fmt"
	"net/http"
)

type ProxyTimeoutClient struct{}

// Do will do a pretend http request. Do be doo.
func (ptc ProxyTimeoutClient) Do(r *http.Request) (*http.Response, error) {
	fmt.Println("proxy request to: " + r.Host)
	return &http.Response{}, nil
}
