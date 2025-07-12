package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

type LocationArea struct {
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"results"`
}

type LocationAreaDetail struct {
	Name              string `json:"name"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

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

func (c *Client) ListLocationDetails(location string) (LocationAreaDetail, error) {
	fullURL := baseURL + "/location-area/" + location

	if data, ok := c.cache.Get(fullURL); ok {
		locationAreaDetails := LocationAreaDetail{}
		if err := json.Unmarshal(data, &locationAreaDetails); err != nil {
			return LocationAreaDetail{}, err
		}
		return locationAreaDetails, nil
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationAreaDetail{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreaDetail{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreaDetail{}, err
	}

	locationAreaDetails := LocationAreaDetail{}
	if err := json.Unmarshal(data, &locationAreaDetails); err != nil {
		return LocationAreaDetail{}, err
	}
	c.cache.Add(fullURL, data)
	return locationAreaDetails, nil
}
