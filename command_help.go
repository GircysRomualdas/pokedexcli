package main

import (
	"fmt"
)

func commandHelp(config *config, args []string) error {
	commands := getCommands()
	commandList := ""
	for name, command := range commands {
		commandList += fmt.Sprintf("%s: %s\n", name, command.description)
	}
	fmt.Print(`Welcome to the Pokedex!
Usage:
` + commandList)
	return nil
}
