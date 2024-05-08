package main

import (
	"time"

	"github.com/Katalcha/go-pokedex/internal/pokeApi"
)

func main() {

	pokeClient := pokeApi.NewClient(5*time.Second, 5*time.Minute)

	cfg := &config{
		pokeApiClient: pokeClient,
	}

	loadPokedex(cfg)
}
