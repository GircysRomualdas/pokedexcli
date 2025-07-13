package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/GircysRomualdas/pokedexcli/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(config *config, args []string) error
}

type config struct {
	next          string
	previous      string
	pokeapiClient pokeapi.Client
	pokedex       map[string]pokeapi.Pokemon
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays a map of the Pokedex",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "map_back",
			description: "Get the previous map of the Pokedex",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Explore area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Catch a Pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect a Pokemon",
			callback:    commandInspect,
		},
	}
}

func startRepl(config *config) {
	scanner := bufio.NewScanner(os.Stdin)

	commands := getCommands()

	for {
		fmt.Fprint(os.Stderr, "Pokedex > ")
		scanner.Scan()
		sliceInput := cleanInput(scanner.Text())
		if len(sliceInput) == 0 {
			continue
		}
		command, ok := commands[sliceInput[0]]
		args := sliceInput[1:]
		if !ok {
			fmt.Fprint(os.Stderr, "Unknown command\n")
			continue
		}
		if err := command.callback(config, args); err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
		}
	}
}

func cleanInput(text string) []string {
	textLowered := strings.ToLower(text)
	return strings.Fields(textLowered)
}
