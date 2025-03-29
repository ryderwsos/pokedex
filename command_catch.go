package main

import (
	"fmt"
	"math/rand"
)

// high = high chance of catching, first evo pokemon has around ~60-70 base exp
var chance int = 30

func commandCatch(config *Config, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("no pokemon specified")
	}
	pokemonData, err := config.pokeapiClient.GetPokemon(args[0])
	if err != nil {
		return err
	}

	catchChance := rand.Intn(int(pokemonData.BaseExperience))

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonData.Name)
	if catchChance < chance {
		fmt.Printf("%s was caught!\n", pokemonData.Name)
		config.pokemonList[pokemonData.Name] = pokemonData
	} else {
		fmt.Printf("%s escaped!\n", pokemonData.Name)
	}

	return nil

}
