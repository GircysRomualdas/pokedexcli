package main

import (
	"strings"
)

func cleanInput(text string) []string {
	textLowered := strings.ToLower(text)
	return strings.Fields(textLowered)
}
