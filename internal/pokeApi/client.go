package pokeApi

import (
	"net/http"
	"time"

	"github.com/Katalcha/go-pokedex/internal/pokeCache"
)

// CLIENT
type Client struct {
	cache      pokeCache.Cache
	httpClient http.Client
}

// NEW CLIENT
func NewClient(timeout, cacheInterval time.Duration) Client {
	client := Client{
		cache: pokeCache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
	return client
}
