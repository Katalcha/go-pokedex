package commands

import (
	"errors"
	"fmt"
)

func commandMapb(cfg *Config, args ...string) error {
	if cfg.PreviousLocationsURL == nil {
		return errors.New("you are on the first page")
	}

	locationRespone, err := cfg.PokeApiClient.ListLocations(cfg.PreviousLocationsURL)
	if err != nil {
		return err
	}

	cfg.NextLocationsURL = locationRespone.Next
	cfg.PreviousLocationsURL = locationRespone.Previous

	for _, location := range locationRespone.Results {
		fmt.Println(location.Name)
	}

	return nil
}
