package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
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
	textArr := strings.Fields(text)
	return textArr
}
