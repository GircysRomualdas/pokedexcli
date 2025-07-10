package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Fprint(os.Stderr, "Pokedex > ")
		scanner.Scan()
		rawInput := scanner.Text()
		sliceInput := cleanInput(rawInput)
		if len(sliceInput) == 0 {
			continue
		}
		fmt.Println("Your command was:", sliceInput[0])
	}
}

func cleanInput(text string) []string {
	textLowered := strings.ToLower(text)
	return strings.Fields(textLowered)
}
