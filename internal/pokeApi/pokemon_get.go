package pokeApi

import (
	"encoding/json"
	"io"
	"net/http"
)

// GETPOKEMON
func (client *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + pokemonName

	if value, ok := client.cache.Get(url); ok {
		pokemonResponse := Pokemon{}
		err := json.Unmarshal(value, &pokemonResponse)
		if err != nil {
			return Pokemon{}, err
		}
		return pokemonResponse, nil
	}

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}

	response, err := client.httpClient.Do(request)
	if err != nil {
		return Pokemon{}, err
	}
	defer response.Body.Close()

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return Pokemon{}, err
	}

	pokemonResponse := Pokemon{}
	err = json.Unmarshal(data, &pokemonResponse)
	if err != nil {
		return Pokemon{}, err
	}

	client.cache.Add(url, data)

	return pokemonResponse, nil
}
