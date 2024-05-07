package main

import (
	"errors"
	"fmt"
)

func commandMapb(cfg *config) error {
	if cfg.previousLocationsURL == nil {
		return errors.New("you are on the first page.")
	}

	locationRespone, err := cfg.pokeApiClient.ListLocations(cfg.previousLocationsURL)

	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationRespone.Next
	cfg.previousLocationsURL = locationRespone.Previous

	for _, location := range locationRespone.Results {
		fmt.Println(location.Name)
	}

	return nil
}
