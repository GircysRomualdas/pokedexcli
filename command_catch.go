package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func commandCatch(config *config, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("catch requires an pokemon name")
	}
	pokemonName := args[0]

	pokemon, err := config.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)

	difficulty := pokemon.BaseExperience
	catchAttemptRoll := rand.Intn(200)
	if catchAttemptRoll > difficulty {
		config.pokedex[pokemonName] = pokemon
		fmt.Printf("%s was caught!\n", pokemon.Name)
	} else {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	}
	return nil
}
