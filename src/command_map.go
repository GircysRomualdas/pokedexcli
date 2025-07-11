package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const locationAreaURL = "https://pokeapi.co/api/v2/location-area"

func commandMap(config *config) error {
	url := locationAreaURL
	if config.next != "" {
		url = config.next
	}
	locationArea, err := getLocationArea(url)
	if err != nil {
		return err
	}
	config.next = locationArea.Next
	config.previous = locationArea.Previous

	for _, location := range locationArea.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func commandMapBack(config *config) error {
	if config.previous == "" {
		return fmt.Errorf("you're on the first page")
	}
	url := config.previous
	locationArea, err := getLocationArea(url)
	if err != nil {
		return err
	}
	config.next = locationArea.Next
	config.previous = locationArea.Previous

	for _, location := range locationArea.Results {
		fmt.Println(location.Name)
	}

	return nil
}

type LocationArea struct {
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"results"`
}

func getLocationArea(url string) (LocationArea, error) {
	res, err := http.Get(url)
	if err != nil {
		return LocationArea{}, err
	}
	defer res.Body.Close()

	var locationArea LocationArea
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&locationArea); err != nil {
		return LocationArea{}, err
	}
	return locationArea, nil
}
