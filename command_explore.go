package main

import (
	"fmt"
)

func commandExplore(config *Config, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("no map specified")
	}
	pokemonName := args[0]
	encountersList, err := config.pokeapiClient.GetPokemonEncounters(pokemonName)
	if err != nil {
		return err
	}

	for _, value := range encountersList.PokemonEncounters {
		fmt.Println(value.Pokemon.Name)
	}
	return err
}
