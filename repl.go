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

		command, exist := getCommand()[words[0]]
		if exist {
			err := command.callback()
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
	callback    func() error
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
	}
}
