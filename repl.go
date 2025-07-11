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
	callback    func(config *config) error
}

type config struct {
	next          string
	previous      string
	pokeapiClient pokeapi.Client
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
		if !ok {
			fmt.Fprint(os.Stderr, "Unknown command\n")
			continue
		}
		if err := command.callback(config); err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
		}
	}
}

func cleanInput(text string) []string {
	textLowered := strings.ToLower(text)
	return strings.Fields(textLowered)
}
