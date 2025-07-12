package main

import (
	"fmt"
)

func commandExplore(config *config, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("explore requires an argument")
	}

	areaName := args[0]
	fmt.Printf("Exploring %s...\n", areaName)
	locationDetails, err := config.pokeapiClient.ListLocationDetails(areaName)
	if err != nil {
		return err
	}
	fmt.Println("Found Pokemon:")
	for _, location := range locationDetails.PokemonEncounters {
		fmt.Printf(" - %s\n", location.Pokemon.Name)
	}
	return nil
}
