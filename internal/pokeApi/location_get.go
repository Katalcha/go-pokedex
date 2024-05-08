package pokeApi

import (
	"encoding/json"
	"io"
	"net/http"
)

// GETLOCATION
func (client *Client) GetLocation(locationName string) (Location, error) {
	url := baseURL + "/location-area/" + locationName
	fail := Location{}

	if value, ok := client.cache.Get(url); ok {
		locationResponse := Location{}
		err := json.Unmarshal(value, &locationResponse)
		if err != nil {
			return fail, err
		}
		return locationResponse, nil
	}

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fail, err
	}

	response, err := client.httpClient.Do(request)
	if err != nil {
		return fail, err
	}
	defer response.Body.Close()

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return fail, err
	}

	locationResponse := Location{}
	err = json.Unmarshal(data, &locationResponse)
	if err != nil {
		return fail, err
	}

	client.cache.Add(url, data)

	return locationResponse, nil

}
