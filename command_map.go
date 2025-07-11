package main

import (
	"fmt"
)

func commandMapf(config *config) error {
	locationArea, err := config.pokeapiClient.ListLocations(config.next)
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

func commandMapb(config *config) error {
	if config.previous == "" {
		return fmt.Errorf("you're on the first page")
	}
	locationArea, err := config.pokeapiClient.ListLocations(config.previous)
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
