package pokeApi

import (
	"encoding/json"
	"io"
	"net/http"
)

// LISTLOCATIONS
func (client *Client) ListLocations(pageURL *string) (ResponseShallowLocations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ResponseShallowLocations{}, err
	}

	response, err := client.httpClient.Do(request)
	if err != nil {
		return ResponseShallowLocations{}, err
	}
	defer response.Body.Close()

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return ResponseShallowLocations{}, err
	}

	locationsResponse := ResponseShallowLocations{}
	err = json.Unmarshal(data, &locationsResponse)
	if err != nil {
		return ResponseShallowLocations{}, err
	}

	return locationsResponse, nil
}
