package main

import (
	"fmt"
)

func commandPokedex(config *config, args []string) error {
	if len(config.pokedex) == 0 {
		fmt.Println("Your Pokedex is empty. Catch some Pokemon first!")
		return nil
	}

	fmt.Println("Your Pokedex:")
	for name := range config.pokedex {
		fmt.Printf(" - %s\n", name)
	}
	return nil
}
