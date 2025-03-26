package main

import (
	"fmt"
)

func commandHelp(config *Config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("-----------------------------------------------------")
	fmt.Println("Pokedex commands, type first word:")
	for _, i := range getCommand() {
		fmt.Printf("\"%s\" - %s\n", i.name, i.description)
	}
	fmt.Println("-----------------------------------------------------")
	return nil
}
