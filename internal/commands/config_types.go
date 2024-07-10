package commands

import "github.com/Katalcha/go-pokedex/internal/pokeApi"

type Config struct {
	PokeApiClient        pokeApi.Client
	NextLocationsURL     *string
	PreviousLocationsURL *string
	CaughtPokemon        map[string]pokeApi.Pokemon
}
