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
			err := command.callback(config, words[1:])
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
	callback    func(*Config, []string) error
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
		"explore": {
			name:        "explore [map]",
			description: "Gives a list of all pokemons in that given area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch [pokemon]",
			description: "Attempt to catch the specified pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect [pokemon]",
			description: "Inspect said pokemon that you've caught",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Print a list of all the pokemons in your Pokedex",
			callback:    commandPokedex,
		},
	}
}

type Config struct {
	Previous      *string
	Next          *string
	pokeapiClient pokeapi.Client
	pokemonList   map[string]pokeapi.Pokemon
}
