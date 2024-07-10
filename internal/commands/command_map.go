package commands

import (
	"fmt"
)

func commandMap(cfg *Config, args ...string) error {
	locationResponse, err := cfg.PokeApiClient.ListLocations(cfg.NextLocationsURL)
	if err != nil {
		return err
	}

	cfg.NextLocationsURL = locationResponse.Next
	cfg.PreviousLocationsURL = locationResponse.Previous

	for _, location := range locationResponse.Results {
		fmt.Println(location.Name)
	}

	return nil
}
