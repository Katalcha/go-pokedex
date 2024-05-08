package main

import (
	"fmt"
)

func commandMap(cfg *config, args ...string) error {
	locationResponse, err := cfg.pokeApiClient.ListLocations(cfg.nextLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationResponse.Next
	cfg.previousLocationsURL = locationResponse.Previous

	for _, location := range locationResponse.Results {
		fmt.Println(location.Name)
	}

	return nil
}
