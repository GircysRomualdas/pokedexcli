package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL string) (LocationArea, error) {
	fullURL := baseURL + "/location-area"
	if pageURL != "" {
		fullURL = pageURL
	}
	if data, ok := c.cache.Get(fullURL); ok {
		locationArea := LocationArea{}
		if err := json.Unmarshal(data, &locationArea); err != nil {
			return LocationArea{}, err
		}
		return locationArea, nil
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationArea{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationArea{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationArea{}, err
	}

	locationArea := LocationArea{}
	if err := json.Unmarshal(data, &locationArea); err != nil {
		return LocationArea{}, err
	}
	c.cache.Add(fullURL, data)
	return locationArea, nil
}
