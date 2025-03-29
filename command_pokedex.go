package main

import "fmt"

func commandPokedex(config *Config, args []string) error {
	fmt.Println("Your pokedex")
	for name, _ := range config.pokemonList {
		fmt.Printf(" - %s\n", name)
	}

	return nil
}
