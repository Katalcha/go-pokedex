package pokeApi

import (
	"net/http"
	"time"
)

// CLIENT
type Client struct {
	httpClient http.Client
}

// NEW CLIENT
func NewClient(timeout time.Duration) Client {
	client := Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
	return client
}
