package main

import "fmt"

func commandInspect(config *Config, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("no pokemon specified, please follow inspect [pokemon name]")
	}
	pokemon, ok := config.pokemonList[args[0]]
	if !ok {
		return fmt.Errorf("%s not in your pokemon list", args[0])
	}

	fmt.Printf("Height: %d\nWeight: %d", pokemon.Height, pokemon.Weight)

	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("	-%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Println("Types:")
	for _, ptype := range pokemon.Types {
		fmt.Printf("	-%s\n", ptype.Type.Name)
	}
	return nil
}
