package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ryderwsos/pokedex/internal/pokeapi"
)

func startRepl(config *Config) {
	var userInput string
	input := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("Pokedex > ")

		if input.Scan() {
			userInput = input.Text()
		}

		words := cleanInput(userInput)

		command, exist := getCommand()[words[0]]
		if exist {
			err := command.callback(config)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("unkown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	lowerText := strings.ToLower(text)
	textArr := strings.Fields(lowerText)
	return textArr
}

type cliCommand struct {
	name        string
	description string
	callback    func(*Config) error
}

func getCommand() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Display all console commands",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Gives a list of the next 20 areas on the Pokedex",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Gives a list of the previous 20 areas on the Pokedex",
			callback:    commandMapb,
		},
	}
}

type Config struct {
	Previous      *string
	Next          *string
	pokeapiClient pokeapi.Client
}
