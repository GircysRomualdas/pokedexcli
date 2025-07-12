package main

import (
	"time"

	"github.com/GircysRomualdas/pokedexcli/internal/pokeapi"
	"github.com/GircysRomualdas/pokedexcli/internal/pokecache"
)

func main() {
	cache := pokecache.NewCache(5 * time.Second)
	client := pokeapi.NewClient(5*time.Second, cache)
	config := &config{
		pokeapiClient: client,
		pokedex:       make(map[string]pokeapi.Pokemon),
	}
	startRepl(config)
}
