package main

import (
	"fmt"
	"os"
)

func commandExit(config *config, args []string) error {
	fmt.Print("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}
