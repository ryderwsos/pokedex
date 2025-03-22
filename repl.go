package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	var userInput string
	input := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("Pokedex > ")

		if input.Scan() {
			userInput = input.Text()
		}

		words := cleanInput(userInput)

		fmt.Printf("Your command was: %s\n", strings.ToLower(words[0]))
	}
}

func cleanInput(text string) []string {
	lowerText := strings.ToLower(text)
	textArr := strings.Fields(lowerText)
	return textArr
}
